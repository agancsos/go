package gmonrpc
import (
	common "../common"
	services "../gmon/services"
	models "../gmon/models"
	"strconv"
	context "context"
    grpc "google.golang.org/grpc"
	status "google.golang.org/grpc/status"
	codes "google.golang.org/grpc/codes"
)

func GetAgentProxy(a *models.Agent) (MonitorAgentServiceClient) {
	// Use grpc.WaitForReady(true) in the client calls
    handle, _ := grpc.Dial(a.PublicIP(), grpc.WithInsecure());
    client := NewMonitorAgentServiceClient(handle);
    return client;
}

// ManagementPackService
type rpcManagementPackServiceServer struct {
	mps          *services.LocalManagementPackService
}

func NewRpcManagementPackServer() (*rpcManagementPackServiceServer) {
	var result = &rpcManagementPackServiceServer{};
	result.mps = (services.GetLocalMPServiceInstance()).(*services.LocalManagementPackService);
	return result;
}

func (x rpcManagementPackServiceServer) ImportMP(c context.Context, m *IncomingServerMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.mps.ImportMP(json))}, nil;
}
func (x rpcManagementPackServiceServer) ExportMP(c context.Context, m *IncomingServerMessage) (*JsonMessage, error) {
	var json = *m.Json;
    return &JsonMessage{Json: common.StrToStrPtr(x.mps.ExportMP(json))}, nil;
}
func (x rpcManagementPackServiceServer) AddMP(c context.Context, m *IncomingServerMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var mp = &models.ManagementPack{};
    mp.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.mps.AddMP(mp))}, nil;
}
func (x rpcManagementPackServiceServer) DeleteMP(c context.Context, m *IncomingServerMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var mp = &models.ManagementPack{};
    mp.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.mps.DeleteMP(mp))}, nil;
}
func (x rpcManagementPackServiceServer) GetManagementPacks(c context.Context, m *EmptyMessage) (*JsonMessage, error) {
	var rsp = "{\"managementPacks\":[";
    var packs = x.mps.GetManagementPacks();
    for i, pack := range packs {
        if i > 0 {
            rsp += ",";
        }
        rsp += pack.ToJsonString();
    }
    rsp += "]";
	return &JsonMessage{Json: common.StrToStrPtr(rsp)}, nil;
}
func (x rpcManagementPackServiceServer) GetManagementPack(c context.Context, m *NumericMessage) (*JsonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagementPack not implemented")
}
func (x rpcManagementPackServiceServer) mustEmbedUnimplementedManagementPackServiceServer() {}
/*****************************************************************************/

// AgentService
type rpcAgentServiceServer struct {
	las            *services.LocalAgentService
}

func NewRpcAgentServiceServer() (*rpcAgentServiceServer) {
	var result = &rpcAgentServiceServer{};
    result.las = (services.GetLocalAgentServiceInstance()).(*services.LocalAgentService);
    return result;
}

func (x rpcAgentServiceServer) GetAgent(c context.Context, m *NumericMessage) (*JsonMessage, error) {
    var node = x.las.GetAgent(int(*m.Value));
	return &JsonMessage{Json: common.StrToStrPtr(node.ToJsonString())}, nil;
}
func (x rpcAgentServiceServer) GetAgents(c context.Context, m *EmptyMessage) (*JsonMessage, error) {
	var rsp = "{\"nodes\":[";
    var nodes = x.las.GetAgents();
    for i, node := range nodes {
        if i > 0 {
            rsp += ",";
        }
        rsp += node.ToJsonString();
    }
    rsp += "]}";
	return &JsonMessage{Json: common.StrToStrPtr(rsp)}, nil;
}
func (x rpcAgentServiceServer) AddAgent(c context.Context, m *IncomingAgentMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var agent = &models.Agent{};
    agent.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.las.AddAgent(agent))}, nil;
}
func (x rpcAgentServiceServer) UpdateAgent(c context.Context, m *IncomingAgentMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var agent = &models.Agent{};
    agent.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.las.UpdateAgent(agent))}, nil;
}
func (x rpcAgentServiceServer) RemoveAgent(c context.Context, m *IncomingAgentMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var agent = &models.Agent{};
    agent.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.las.RemoveAgent(agent))}, nil;
}
func (x rpcAgentServiceServer) Contains(c context.Context, m *IncomingAgentMessage) (*BooleanMessage, error) {
	var hostname = *m.Json;
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.las.Contains(hostname))}, nil;
}
func (x rpcAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}
/*****************************************************************************/

// RolerService
type rpcRoleServiceServer struct {
	lrs            *services.LocalRoleService
}

func NewRpcRoleServiceServer() (*rpcRoleServiceServer) {
	var result = &rpcRoleServiceServer{};
    result.lrs = (services.GetLocalRoleServiceInstance()).(*services.LocalRoleService);
    return result;
}

func (x rpcRoleServiceServer) GetRoles(c context.Context, m *EmptyMessage) (*JsonMessage, error) {
	var rsp = "{\"roles\":[";
    var roles = x.lrs.GetRoles();
    for i, role := range roles {
        if i > 0 {
            rsp += ",";
        }
        rsp += role.ToJsonString();
    }
    rsp += "]}";
    return &JsonMessage{Json: common.StrToStrPtr(rsp)}, nil;
}
func (x rpcRoleServiceServer) GetRole(c context.Context, m *IncomingServerMessage) (*JsonMessage, error) {
    id,_ := strconv.Atoi(*m.Json);
    var role = x.lrs.GetRole(id);
	return &JsonMessage{Json: common.StrToStrPtr(role.ToJsonString())}, nil;
}
func (x rpcRoleServiceServer) AddRole(c context.Context, m *IncomingServerMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var role = &models.Role{};
    role.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.lrs.AddRole(role))}, nil;
}
func (x rpcRoleServiceServer) UpdateRole(c context.Context, m *IncomingServerMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var role = &models.Role{};
    role.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.lrs.UpdateRole(role))}, nil;
}
func (x rpcRoleServiceServer) DeleteRole(c context.Context, m *IncomingServerMessage) (*BooleanMessage, error) {
	var json = *m.Json;
    var role = &models.Role{};
    role.ReloadFromJson(json);
    return &BooleanMessage{Status: common.BoolToBoolPtr(x.lrs.DeleteRole(role))}, nil;
}
func (x rpcRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}
/*****************************************************************************/

// MonitorAgentService
type rpcMonitorAgentServiceServer struct {
	las            *services.LocalAgentService
}

func NewRpcMonitorAgentServiceServer() (*rpcMonitorAgentServiceServer) {
	var result = &rpcMonitorAgentServiceServer{};
    result.las = services.GetLocalAgentServiceInstance().(*services.LocalAgentService);
    return result;
}

func (x rpcMonitorAgentServiceServer) IncomingManagementPack(c context.Context, m *IncomingAgentMessage) (*BooleanMessage, error) {
	var json = *m.Json;
	var mp = &models.ManagementPack{};
	mp.ReloadFromJson(json);
	x.las.AcceptIncomingMP(mp);
	return &BooleanMessage{Status: common.BoolToBoolPtr(true)}, nil;
}
func (rpcMonitorAgentServiceServer) mustEmbedRpcMonitorAgentServiceServer() {}
func (rpcMonitorAgentServiceServer) mustEmbedUnimplementedMonitorAgentServiceServer() {}
/*****************************************************************************/
