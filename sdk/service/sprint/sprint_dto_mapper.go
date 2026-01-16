package sprint

import (
	apisprint "github.com/brain-xai/pingcode_api/internal/api/sprint"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// ToCreateRequestDTO 转换创建迭代输入为 DTO
func ToCreateRequestDTO(input sprintmodel.SprintCreateInput) apisprint.CreateSprintRequest {
	return apisprint.CreateSprintRequest{
		Name:        input.Name,
		StartAt:     input.StartAt,
		EndAt:       input.EndAt,
		AssigneeID:  input.AssigneeID,
		Description: input.Description,
		Status:      input.Status,
		CategoryIDs: input.CategoryIDs,
	}
}

// ToUpdateRequestDTO 转换更新迭代输入为 DTO
func ToUpdateRequestDTO(input sprintmodel.SprintUpdateInput) apisprint.UpdateSprintRequest {
	return apisprint.UpdateSprintRequest{
		Name:        input.Name,
		StartAt:     input.StartAt,
		EndAt:       input.EndAt,
		AssigneeID:  input.AssigneeID,
		Description: input.Description,
		Status:      input.Status,
		CategoryIDs: input.CategoryIDs,
	}
}

// ToBatchCreateRequestDTO 转换批量创建迭代输入为 DTO
func ToBatchCreateRequestDTO(inputs []sprintmodel.SprintCreateInput) apisprint.BatchCreateSprintsRequest {
	sprints := make([]apisprint.CreateSprintRequest, len(inputs))
	for i, input := range inputs {
		sprints[i] = ToCreateRequestDTO(input)
	}
	return apisprint.BatchCreateSprintsRequest{
		Sprints: sprints,
	}
}
