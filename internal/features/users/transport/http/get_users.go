package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/mrhexvel/golang-first-app/internal/core/logger"
	core_http_response "github.com/mrhexvel/golang-first-app/internal/core/transport/http/response"
	core_http_utils "github.com/mrhexvel/golang-first-app/internal/core/transport/http/utils"
)

type GetUsersRespones []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responesHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responesHandler.ErrorResponse(err,
			"failed to get 'limit'/'offset' query param",
		)

		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responesHandler.ErrorResponse(err, "failed to get users")

		return
	}

	response := GetUsersRespones(usersDTOFromDomains(userDomains))

	responesHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParams(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParams(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
