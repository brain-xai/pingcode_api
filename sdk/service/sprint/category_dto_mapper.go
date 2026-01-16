package sprint

import (
	apisprint "github.com/brain-xai/pingcode_api/internal/api/sprint"
	sprintmodel "github.com/brain-xai/pingcode_api/sdk/model/sprint"
)

// ToCreateCategoryRequestDTO 转换创建迭代类别输入为 DTO
func ToCreateCategoryRequestDTO(input sprintmodel.SprintCategoryCreateInput) apisprint.CreateSprintCategoryRequest {
	return apisprint.CreateSprintCategoryRequest{
		Name:    input.Name,
		GroupID: input.GroupID, // API 中称为 section_id
	}
}

// ToUpdateCategoryRequestDTO 转换更新迭代类别输入为 DTO
func ToUpdateCategoryRequestDTO(input sprintmodel.SprintCategoryUpdateInput) apisprint.UpdateSprintCategoryRequest {
	req := apisprint.UpdateSprintCategoryRequest{
		Name: *input.Name,
	}
	if input.GroupID != nil {
		req.GroupID = *input.GroupID
	}
	return req
}
