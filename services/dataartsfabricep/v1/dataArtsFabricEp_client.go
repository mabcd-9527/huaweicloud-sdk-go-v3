package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsfabricep/v1/model"
)

type DataArtsFabricEpClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDataArtsFabricEpClient(hcClient *httpclient.HcHttpClient) *DataArtsFabricEpClient {
	return &DataArtsFabricEpClient{HcClient: hcClient}
}

func DataArtsFabricEpClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CancelJobRun 取消作业运行
//
// 取消作业运行。主要在取消运行Ray job/python job/spark job等job场景使用；输入workspace_id和run_id；输出为接口运行成功或失败的响应消息，无具体的返回值内容。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) CancelJobRun(request *model.CancelJobRunRequest) (*model.CancelJobRunResponse, error) {
	requestDef := GenReqDefForCancelJobRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelJobRunResponse), nil
	}
}

// CancelJobRunInvoker 取消作业运行
func (c *DataArtsFabricEpClient) CancelJobRunInvoker(request *model.CancelJobRunRequest) *CancelJobRunInvoker {
	requestDef := GenReqDefForCancelJobRun()
	return &CancelJobRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ClearJobRun 删除端点下所有的作业
//
// 删除端点下所有的job运行记录。主要在删除指定端点下所有的Ray job/python job/spark job等job场景使用；输入参数workspace_id和endpoint_id；输出为接口运行成功或失败的响应消息，无具体的返回值内容。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ClearJobRun(request *model.ClearJobRunRequest) (*model.ClearJobRunResponse, error) {
	requestDef := GenReqDefForClearJobRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClearJobRunResponse), nil
	}
}

// ClearJobRunInvoker 删除端点下所有的作业
func (c *DataArtsFabricEpClient) ClearJobRunInvoker(request *model.ClearJobRunRequest) *ClearJobRunInvoker {
	requestDef := GenReqDefForClearJobRun()
	return &ClearJobRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJobRun 删除作业
//
// 删除作业。主要在删除Ray job/python job/spark job等job场景使用；输入workspace_id和run_id；输出为接口运行成功或失败的响应消息，无具体的返回值内容。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) DeleteJobRun(request *model.DeleteJobRunRequest) (*model.DeleteJobRunResponse, error) {
	requestDef := GenReqDefForDeleteJobRun()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobRunResponse), nil
	}
}

// DeleteJobRunInvoker 删除作业
func (c *DataArtsFabricEpClient) DeleteJobRunInvoker(request *model.DeleteJobRunRequest) *DeleteJobRunInvoker {
	requestDef := GenReqDefForDeleteJobRun()
	return &DeleteJobRunInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobRuns 查看作业运行列表
//
// 查看作业运行列表。主要在查询所有Ray job/python job/spark job等job场景使用；输入workspace_id，job运行id(可选)，job运行名称（可选），分页查询参数limit和offset，job id（可选），endpoint id（可选），job version id（可选），job状态（可选），job类型（可选），排序参数及升序/降序排序方式；输出job运行信息列表。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ListJobRuns(request *model.ListJobRunsRequest) (*model.ListJobRunsResponse, error) {
	requestDef := GenReqDefForListJobRuns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobRunsResponse), nil
	}
}

// ListJobRunsInvoker 查看作业运行列表
func (c *DataArtsFabricEpClient) ListJobRunsInvoker(request *model.ListJobRunsRequest) *ListJobRunsInvoker {
	requestDef := GenReqDefForListJobRuns()
	return &ListJobRunsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunJob 运行作业
//
// 运行作业。主要在运行Ray job/python job/spark job等job场景使用；输入workspace_id，job信息（jobId、jobVersionId），job运行名称，job运行端点id；输出作业运行Id。此接口为异步接口，配套使用接口showJobRunDetail接口或ListJobRuns接口查询job运行状态，无特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) RunJob(request *model.RunJobRequest) (*model.RunJobResponse, error) {
	requestDef := GenReqDefForRunJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunJobResponse), nil
	}
}

// RunJobInvoker 运行作业
func (c *DataArtsFabricEpClient) RunJobInvoker(request *model.RunJobRequest) *RunJobInvoker {
	requestDef := GenReqDefForRunJob()
	return &RunJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobRunDetail 查看作业运行详情
//
// 查看指定作业运行详情。主要在Ray job/python job/spark job等job场景用户查询job运行详细信息；输入workspace_id和run_id；输出job的详细信息，包括id、创建时间、更新时间、状态、运行时长、job类型、错误码、错误信息、错误解决方案、远端信息，其中只有job错误时才会有错误码、错误信息等。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ShowJobRunDetail(request *model.ShowJobRunDetailRequest) (*model.ShowJobRunDetailResponse, error) {
	requestDef := GenReqDefForShowJobRunDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobRunDetailResponse), nil
	}
}

// ShowJobRunDetailInvoker 查看作业运行详情
func (c *DataArtsFabricEpClient) ShowJobRunDetailInvoker(request *model.ShowJobRunDetailRequest) *ShowJobRunDetailInvoker {
	requestDef := GenReqDefForShowJobRunDetail()
	return &ShowJobRunDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceInstance 删除Service实例
//
// 删除Service实例，释放该实例的资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) DeleteServiceInstance(request *model.DeleteServiceInstanceRequest) (*model.DeleteServiceInstanceResponse, error) {
	requestDef := GenReqDefForDeleteServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceInstanceResponse), nil
	}
}

// DeleteServiceInstanceInvoker 删除Service实例
func (c *DataArtsFabricEpClient) DeleteServiceInstanceInvoker(request *model.DeleteServiceInstanceRequest) *DeleteServiceInstanceInvoker {
	requestDef := GenReqDefForDeleteServiceInstance()
	return &DeleteServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployServiceInstance 部署服务
//
// 部署一个Service实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) DeployServiceInstance(request *model.DeployServiceInstanceRequest) (*model.DeployServiceInstanceResponse, error) {
	requestDef := GenReqDefForDeployServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployServiceInstanceResponse), nil
	}
}

// DeployServiceInstanceInvoker 部署服务
func (c *DataArtsFabricEpClient) DeployServiceInstanceInvoker(request *model.DeployServiceInstanceRequest) *DeployServiceInstanceInvoker {
	requestDef := GenReqDefForDeployServiceInstance()
	return &DeployServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServicesInstances 列举已部署的Service实例
//
// 列举已部署的Service实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ListServicesInstances(request *model.ListServicesInstancesRequest) (*model.ListServicesInstancesResponse, error) {
	requestDef := GenReqDefForListServicesInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesInstancesResponse), nil
	}
}

// ListServicesInstancesInvoker 列举已部署的Service实例
func (c *DataArtsFabricEpClient) ListServicesInstancesInvoker(request *model.ListServicesInstancesRequest) *ListServicesInstancesInvoker {
	requestDef := GenReqDefForListServicesInstances()
	return &ListServicesInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServiceInstanceDetail 查看部署的Service实例详情
//
// 查看部署后的Service实例的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ShowServiceInstanceDetail(request *model.ShowServiceInstanceDetailRequest) (*model.ShowServiceInstanceDetailResponse, error) {
	requestDef := GenReqDefForShowServiceInstanceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceInstanceDetailResponse), nil
	}
}

// ShowServiceInstanceDetailInvoker 查看部署的Service实例详情
func (c *DataArtsFabricEpClient) ShowServiceInstanceDetailInvoker(request *model.ShowServiceInstanceDetailRequest) *ShowServiceInstanceDetailInvoker {
	requestDef := GenReqDefForShowServiceInstanceDetail()
	return &ShowServiceInstanceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServiceInstance 更新Service实例
//
// 更新已部署的Service实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) UpdateServiceInstance(request *model.UpdateServiceInstanceRequest) (*model.UpdateServiceInstanceResponse, error) {
	requestDef := GenReqDefForUpdateServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceInstanceResponse), nil
	}
}

// UpdateServiceInstanceInvoker 更新Service实例
func (c *DataArtsFabricEpClient) UpdateServiceInstanceInvoker(request *model.UpdateServiceInstanceRequest) *UpdateServiceInstanceInvoker {
	requestDef := GenReqDefForUpdateServiceInstance()
	return &UpdateServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
