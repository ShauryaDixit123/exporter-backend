package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateInstance(
	ctx *gin.Context,
) {
	var req rdbms.CreateWorkflowInstanceI
	if er := ctx.ShouldBindQuery(&req); er != nil {
		ctx.JSON(500, er)
	}
	if _, er := h.workflowService.CreateWorkflowInstance(req); er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, gin.H{"message": "created successfully"})
}

// {
//     "name" : "default_standard_workflow_one",
//     "type" : "default_standard_workflow_one",
//     "flows" : [
//         {
//             "type" : "default_standard_inbuilt_purchase_order",
//             "title" :  "Purchase order",
//             "description" : "purchase order creation",
//             "order" : 1,
//             "tat" : 48,
//             "flow_params" : [
//                 {
//                     "name" : "created by buyer",
//                     "type" : "default_flow_params",
//                     "mandatory" : true
//                 },
//                 {
//                     "name" : "received by seller",
//                     "type" : "default_flow_params",
//                     "mandatory" : true
//                 },
//                 {
//                     "name" : "seller approved",
//                     "type" : "default_flow_params",
//                     "mandatory" : true
//                 }
//             ]
//         },
//         {
//             "type" : "default_standard_inbuilt_sales_order",
//             "title" :  "Purchase order",
//             "description" : "purchase order creation",
//             "order" : 2,
//             "tat" : 24,
//             "flow_params" : [
//                 {
//                     "name" : "created by seller",
//                     "type" : "default_flow_params",
//                     "mandatory" : true
//                 },
//                 {
//                     "name" : "received by buyer",
//                     "type" : "default_flow_params",
//                     "mandatory" : true
//                 },
//                 {
//                     "name" : "buyer approved",
//                     "type" : "default_flow_params",
//                     "mandatory" : true
//                 }
//             ]
//         }
//     ]
// }
