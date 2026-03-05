package transitions

import (
	"math"

	"github.com/kuetix/engine/pkg/domain"
	"github.com/kuetix/engine/pkg/domain/interfaces"
	"github.com/kuetix/engine/pkg/helpers"
	"github.com/kuetix/engine/pkg/workflow"
)

type paginationTransitions struct {
	workflow.BaseServiceTransition
}

func NewPaginationTransitions() interfaces.ServiceTransitions {
	return &paginationTransitions{}
}

//goland:noinspection GoUnusedParameter
func (pt *paginationTransitions) PaginationTotal(pagination *domain.Pagination, total int) (r domain.FlowStepResult) {
	if pagination == nil {
		pagination = &domain.Pagination{
			Page:             1,
			PageSize:         999999999,
			TotalPage:        0,
			TotalRecords:     0,
			CurrentPageTotal: 0,
		}
	}
	pagination.TotalRecords = total
	pagination.TotalPage = int(math.Ceil(float64(total) / float64(pagination.PageSize)))

	r.Success = true
	return
}

//goland:noinspection GoUnusedParameter
func (pt *paginationTransitions) CurrentPageTotal(pagination *domain.Pagination, entities []interface{}) (r domain.FlowStepResult) {
	pagination.CurrentPageTotal = helpers.Len(entities)

	r.Success = true
	return
}
