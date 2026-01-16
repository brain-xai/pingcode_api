package sprint

import (
	apisprint "github.com/brain-xai/pingcode_api/internal/api/sprint"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// ToCreateGroupRequestDTO 转换创建迭代分组输入为 DTO
func ToCreateGroupRequestDTO(input sprintmodel.SprintGroupCreateInput) apisprint.CreateSprintGroupRequest {
	return apisprint.CreateSprintGroupRequest{
		Name: input.Name,
	}
}

// ToUpdateGroupRequestDTO 转换更新迭代分组输入为 DTO
func ToUpdateGroupRequestDTO(input sprintmodel.SprintGroupUpdateInput) apisprint.UpdateSprintGroupRequest {
	return apisprint.UpdateSprintGroupRequest{
		Name: *input.Name,
	}
}
