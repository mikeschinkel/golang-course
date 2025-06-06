# Challenging Go Infrastructure Scenarios

## Scenario 1: Real-time Kubernetes Operator Debugging in Complex Container Environment
*Client - Kubernetes Operator Team*

### Background Context
At Client, I worked on the Kubernetes Operator team where we needed to debug complex operator logic running inside Kubernetes clusters. The challenge was setting up a development environment that allowed real-time debugging of Go code running in a Kubernetes context while maintaining the full complexity of the distributed system.

### The Infrastructure Challenge
The development setup involved multiple nested virtualization layers:
- Docker Desktop running on macOS
- K3d (Kubernetes in Docker) running inside Docker
- K3s (lightweight Kubernetes) running inside K3d
- Client Kubernetes Operator (written in Go) running as pods inside K3s
- Need for real-time debugging with Delve debugger from GoLand/VSCode

**The Problem:** Traditional debugging approaches failed because:
1. The Go application was running inside multiple container layers
2. Network routing between the IDE debugger and the Delve process was complex
3. Pod restarts would break debugging sessions
4. Performance overhead from the nested virtualization was significant
5. Port forwarding through multiple network layers was unreliable

### The Technical Challenge
The main issues were:
- **Network Connectivity**: Establishing stable debugger connections through Docker → K3d → K3s → Pod
- **Process Lifecycle**: Kubernetes would restart pods, breaking debugging sessions
- **Resource Constraints**: Multiple virtualization layers consumed significant CPU/memory
- **Development Velocity**: The typical edit-commit-deploy-build-test-debug cycle took 2+ hours

### The Solution Architecture
I developed a multi-part solution:

1. **Custom Delve Integration**: Modified the operator's Docker image to include Delve and expose debugging ports
2. **Network Bridge Setup**: Created a complex port-forwarding chain:
3. **IDE Configuration**: Set up GoLand/VSCode remote debugging configurations with retry logic

### Infrastructure Impact
- **Development Velocity**: Reduced debug cycle from 2+ hours to ~5 minutes
- **Resource Optimization**: Tuned container resource limits to balance debugging capability with system stability
- **Team Productivity**: Enabled the entire team to debug complex operator logic in real-time
- **CI/CD Integration**: Created separate debug and production build pipelines

### Key Lessons Learned
- Container networking in nested environments requires careful port management
- Kubernetes debugging needs special consideration for pod lifecycle management
- Development tooling can significantly impact team velocity in complex infrastructures

---

## Scenario 2: HTTPS Proxy Tunnel Implementation for API Testing Infrastructure
*Resideo/FrontSide - Microservices & Testing*

### Background Context
While building a microservices testing framework for Resideo, we needed to create comprehensive API tests for services that made calls to external APIs (specifically GitHub API). The challenge was intercepting and mocking HTTPS traffic from Go applications that used hardcoded external hostnames.

### The Infrastructure Challenge
Our testing architecture involved:
- Go microservices using the Google `go-github` library (hardcoded to `api.github.com`)
- Node.js Mountebank server for API mocking
- Need to intercept HTTPS traffic without modifying application code
- GitOps workflow requiring reliable API testing in CI/CD pipelines

**The Problem:** When Go's `net/http` client encounters a proxy for HTTPS requests, it sends an HTTP `CONNECT` request to establish a tunnel. However, Mountebank didn't support `CONNECT` requests and would simply close connections.

### The Technical Challenge
The core issues were:

1. **TLS Tunnel Establishment**: HTTPS proxying requires handling `CONNECT` requests to establish encrypted tunnels
2. **Certificate Management**: Need to handle TLS certificates in the proxy context
3. **Transparent Proxying**: The application code couldn't be modified (third-party library)
4. **Testing Infrastructure**: CI/CD pipelines needed reliable, fast API mocking

### The Root Cause Analysis
When I traced through the network traffic, I discovered:

```go
// Go's net/http automatically sends CONNECT for HTTPS proxying
GET / CONNECT HTTP/1.1
Host: api.github.com:443
Proxy-Connection: keep-alive
```

Mountebank's Node.js server was built on the standard HTTP module, which [explicitly closes connections for unhandled CONNECT requests](https://nodejs.org/docs/latest-v14.x/api/http.html#http_event_connect).

### The Solution Implementation
I implemented a two-phase solution:

**Phase 1: Go-based MitM Proxy**
- Built a custom HTTP proxy in Go that properly handled `CONNECT` requests
- Implemented TLS certificate generation for transparent HTTPS interception
- Added request/response logging and modification capabilities

**Phase 2: Mountebank Enhancement**
- Diagnosed the Node.js limitation in Mountebank's codebase
- Implemented `CONNECT` request handling in Mountebank's HTTP server
- Added support for HTTPS tunnel establishment
- Contributed the fix back to the open-source project

### Infrastructure Impact
- **Testing Reliability**: Eliminated flaky tests caused by external API dependencies
- **CI/CD Performance**: Tests ran 5x faster using local mocks vs external APIs
- **Development Workflow**: Enabled offline development and testing
- **Open Source Contribution**: Fixed a fundamental limitation affecting the broader community

### The Technical Implementation
The key was implementing proper HTTP tunnel establishment:

```go
func handleConnect(w http.ResponseWriter, r *http.Request) error {
    var err error
    var targetConn net.Conn
    var clientConn net.Conn
    
    // Establish connection to target server
    if targetConn, err = net.Dial("tcp", r.Host); err != nil {
        goto end
    }
    
    // Hijack the client connection
    if clientConn, err = hijackConnection(w); err != nil {
        goto end
    }
    
    // Send 200 Connection Established
    if err = sendConnectResponse(clientConn); err != nil {
        goto end
    }
    
    // Proxy data bidirectionally
    err = proxyData(clientConn, targetConn)
    
end:
    closeConnections(clientConn, targetConn)
    return err
}
```

### Key Lessons Learned
- HTTPS proxying requires understanding both HTTP and TLS protocols
- Third-party library constraints can force creative infrastructure solutions
- Contributing fixes to open-source dependencies benefits the entire ecosystem
- Proper error handling and connection management is critical in proxy implementations

---

---

## Scenario 3: Custom DNS Provider for Kubernetes External-DNS with API Constraints
*Client - External DNS Provider Development*

### Background Context
Client needed a Kubernetes External-DNS provider for their proprietary internal DNS system. External-DNS is a Kubernetes controller that automatically manages DNS records for services and ingresses, but it required a custom provider to integrate with Client's unique DNS API.

### The Infrastructure Challenge
The project involved multiple complex layers:
- Kubernetes External-DNS controller expecting standard DNS provider interface
- Client's proprietary DNS API with significant limitations and quirks
- OpenAPI specification that was massive and needed transformation
- M1 MacBook development environment requiring ARM-compatible tooling
- Complex networking setup with Tart VMs, K3s, and NetPlan

**The Core Problem:** Client's DNS API had fundamental design limitations that didn't align with how External-DNS expected DNS providers to behave, requiring architectural creativity to bridge the gap.

### The Technical Challenge
The main issues were:

1. **API Design Mismatch**: Client's DNS API was designed for manual operations, not programmatic bulk operations
2. **Information Asymmetry**: The API could create records but had limited ability to query existing state
3. **Rate Limiting**: Aggressive rate limiting that wasn't documented in their OpenAPI spec
4. **Development Environment**: Setting up reliable K3s testing on M1 Mac with Tart VMs and complex NetPlan networking

### The Root Cause Analysis
After analyzing the OpenAPI spec and testing the API, I discovered several critical limitations:

```go
// The API could create records but couldn't reliably list them
// This broke External-DNS's reconciliation model which expects:
// 1. List current state
// 2. Compare with desired state  
// 3. Apply changes

type DNSProvider interface {
    // External-DNS expects these operations
    Records(ctx context.Context, zone string) ([]Record, error)  // ❌ Limited support
    ApplyChanges(ctx context.Context, changes *Changes) error    // ✅ Worked
    AdjustEndpoints(endpoints []*Endpoint) []*Endpoint          // ❌ Needed custom logic
}
```

The API would return incomplete data for `Records()`, making state reconciliation nearly impossible with standard approaches.

### The Solution Architecture
I developed a multi-layered solution:

**1. OpenAPI Transformation Pipeline**
The original OpenAPI spec was 15,000+ lines and included many unneeded endpoints:
**2. Development Environment Solution**
Set up reliable K3s testing on M1 Mac:
- Used Tart for ARM-native Linux VMs
- Configured NetPlan networking (steep learning curve)
- Created reproducible VM snapshots for testing
- Implemented port forwarding for External-DNS debugging

### Infrastructure Impact
- **Kubernetes Integration**: Enabled automatic DNS management for Client's container infrastructure
- **Operational Efficiency**: Eliminated manual DNS record management for hundreds of services
- **Development Velocity**: Created reliable local testing environment for complex K8s scenarios
- **API Enhancement**: Identified and documented API limitations, leading to Client improving their DNS API

### The Technical Implementation
The key innovation was the hybrid state management approach:

```go
func (p *Provider) Records(ctx context.Context, zone string) (records []Record, err error) {
    // Try API first (may return partial data)
    var apiRecords []Record
    if apiRecords, err = p.client.ListRecords(ctx, zone); err != nil {
        goto end
    }
    
    // Merge with local cache for complete state
    cachedRecords := p.stateManager.GetCachedRecords(zone)
    records = p.mergeRecordSets(apiRecords, cachedRecords)
    
end:
    return records, err
}
```

### Key Lessons Learned
- **API Design Matters**: Poor API design can force complex workarounds in client code
- **State Management**: Sometimes you need to maintain local state when upstream systems can't provide it reliably
- **OpenAPI Tooling**: Large specs need transformation; consider API subsets for specific use cases
- **Development Environment**: ARM architecture requires different tooling choices and more setup time

---

## Scenario 4: Cross-Platform Build Orchestration for VMware Infrastructure
*Client - DevOps & Infrastructure Automation*

### Background Context
At Client, I worked on their Client's monitoring solution that required building and deploying Go applications across multiple platforms and infrastructure types: Windows Server agents, CentOS-based gateway servers, and Ubuntu-based appliance VMs deployed via VMware vSphere.

### The Infrastructure Challenge
The build and deployment pipeline involved:
- Go applications targeting Windows Server (metrics collection agents)
- Go applications targeting CentOS (gateway servers for data aggregation)
- Jenkins-based CI/CD with complex artifact management
- VMware vSphere automation using govmomi for automated deployment
- jFrog Artifactory for artifact storage and distribution
- Packer for creating custom VM images with embedded Go applications

**The Core Problem:** A critical bare-metal Jenkins server crashed, taking down the entire build pipeline. We needed to recreate the complex build orchestration while improving reliability and moving toward infrastructure-as-code.

### The Technical Challenge
The main challenges were:

1. **Build Complexity**: Multi-platform Go builds with different packaging requirements (RPMs, ISOs, OVAs)
2. **Infrastructure Recreation**: Rebuilding complex Jenkins jobs from minimal documentation
3. **VMware Integration**: Automating VM deployment and testing in vSphere without GUI tools
4. **Artifact Management**: Complex dependency chains between build artifacts
5. **Zero Downtime**: Customer deployments couldn't be interrupted during infrastructure migration

### The Root Cause Analysis
The crashed server revealed several infrastructure anti-patterns:

```bash
# The original process was a fragile chain of manual steps:
# 1. Jenkins builds Go binaries
# 2. Manual packaging into RPMs
# 3. Manual ISO creation with Packer
# 4. Manual OVA conversion
# 5. Manual upload to Artifactory
# 6. Manual deployment testing

# No version control, no reproducibility, no rollback capability
```

The crash happened because everything ran on a single bare-metal server with no backup or version control of the Jenkins configurations.

### The Solution Architecture
I designed and implemented a resilient, automated pipeline:

**1. Infrastructure as Code Migration**
Moved from click-ops to scripted automation:
**2. VMware vSphere Automation**
Used govmomi to automate previously manual deployment testing:
**3. Artifact Pipeline Management**
Created a reliable artifact promotion pipeline:
**4. Packer Integration for Reproducible Images**
Automated the VM image creation process:

### Infrastructure Impact
- **Reliability**: Eliminated single points of failure by moving to containerized Jenkins with persistent volumes
- **Reproducibility**: All builds became version-controlled and reproducible
- **Speed**: Reduced build-to-deployment time from 4+ hours to ~45 minutes
- **Quality**: Automated testing caught integration issues before customer deployment
- **Documentation**: Infrastructure-as-code served as living documentation

### The Technical Implementation
The key innovation was the orchestrated pipeline that handled cross-platform complexity:


### Key Lessons Learned
- **Infrastructure Resilience**: Single points of failure will eventually fail; design for redundancy
- **Automation Pays**: Manual processes are error-prone and don't scale
- **Cross-Platform Complexity**: Go's build system helps, but packaging and deployment vary significantly
- **Version Everything**: Infrastructure, configurations, and artifacts all need version control

---

## Common Themes Across All Four Scenarios

These scenarios demonstrate:
- **Complex Network Debugging**: Understanding multi-layer network stacks and protocols
- **Infrastructure Integration**: Making Go applications work within larger ecosystems
- **Development Tooling**: Building solutions that improve team productivity
- **API Integration Challenges**: Working with imperfect or limited third-party APIs
- **Cross-Platform Deployment**: Managing complexity across different operating systems and environments
- **Automation and Reliability**: Moving from manual processes to automated, resilient systems
- **Open Source Contribution**: Sharing solutions with the broader community

---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*