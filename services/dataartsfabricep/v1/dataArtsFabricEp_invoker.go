package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsfabricep/v1/model"
)

type CancelJobRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelJobRunInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelJobRunInvoker) Invoke() (*model.CancelJobRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelJobRunResponse), nil
	}
}

type ClearJobRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *ClearJobRunInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ClearJobRunInvoker) Invoke() (*model.ClearJobRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearJobRunResponse), nil
	}
}

type DeleteJobRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobRunInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteJobRunInvoker) Invoke() (*model.DeleteJobRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobRunResponse), nil
	}
}

type ListJobRunsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobRunsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobRunsInvoker) Invoke() (*model.ListJobRunsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobRunsResponse), nil
	}
}

type RunJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunJobInvoker) Invoke() (*model.RunJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunJobResponse), nil
	}
}

type ShowJobRunDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobRunDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobRunDetailInvoker) Invoke() (*model.ShowJobRunDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobRunDetailResponse), nil
	}
}

type DeleteServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceInstanceInvoker) Invoke() (*model.DeleteServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceInstanceResponse), nil
	}
}

type DeployServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeployServiceInstanceInvoker) Invoke() (*model.DeployServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployServiceInstanceResponse), nil
	}
}

type ListServicesInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicesInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicesInstancesInvoker) Invoke() (*model.ListServicesInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicesInstancesResponse), nil
	}
}

type ShowServiceInstanceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServiceInstanceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServiceInstanceDetailInvoker) Invoke() (*model.ShowServiceInstanceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServiceInstanceDetailResponse), nil
	}
}

type UpdateServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServiceInstanceInvoker) Invoke() (*model.UpdateServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServiceInstanceResponse), nil
	}
}
