package dto

type CompletedModuleRequest struct {
	ModuleId string `json:"module_id"`
	UserId   string `json:"user_id"`
}

type CompletedModuleResponse struct {
	UserId   string `json:"user_id"`
	ModuleId string `json:"module_id"`
}

type CompletedModuleListResponse struct {
	CompletedModules []CompletedModuleResponse `json:"completed_modules"`
}
