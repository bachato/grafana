/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */
package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/middleware/requestmeta"
	contextmodel "github.com/grafana/grafana/pkg/services/contexthandler/model"
	"github.com/grafana/grafana/pkg/services/ngalert/metrics"
	"github.com/grafana/grafana/pkg/web"
)

type ConvertPrometheusApi interface {
	RouteConvertPrometheusCortexDeleteNamespace(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusCortexDeleteRuleGroup(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusCortexGetNamespace(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusCortexGetRuleGroup(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusCortexGetRules(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusCortexPostRuleGroup(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusCortexPostRuleGroups(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusDeleteAlertmanagerConfig(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusDeleteNamespace(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusDeleteRuleGroup(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusGetAlertmanagerConfig(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusGetNamespace(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusGetRuleGroup(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusGetRules(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusPostAlertmanagerConfig(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusPostRuleGroup(*contextmodel.ReqContext) response.Response
	RouteConvertPrometheusPostRuleGroups(*contextmodel.ReqContext) response.Response
}

func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexDeleteNamespace(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	return f.handleRouteConvertPrometheusCortexDeleteNamespace(ctx, namespaceTitleParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexDeleteRuleGroup(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	groupParam := web.Params(ctx.Req)[":Group"]
	return f.handleRouteConvertPrometheusCortexDeleteRuleGroup(ctx, namespaceTitleParam, groupParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexGetNamespace(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	return f.handleRouteConvertPrometheusCortexGetNamespace(ctx, namespaceTitleParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexGetRuleGroup(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	groupParam := web.Params(ctx.Req)[":Group"]
	return f.handleRouteConvertPrometheusCortexGetRuleGroup(ctx, namespaceTitleParam, groupParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexGetRules(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusCortexGetRules(ctx)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexPostRuleGroup(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	return f.handleRouteConvertPrometheusCortexPostRuleGroup(ctx, namespaceTitleParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusCortexPostRuleGroups(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusCortexPostRuleGroups(ctx)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusDeleteAlertmanagerConfig(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusDeleteAlertmanagerConfig(ctx)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusDeleteNamespace(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	return f.handleRouteConvertPrometheusDeleteNamespace(ctx, namespaceTitleParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusDeleteRuleGroup(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	groupParam := web.Params(ctx.Req)[":Group"]
	return f.handleRouteConvertPrometheusDeleteRuleGroup(ctx, namespaceTitleParam, groupParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusGetAlertmanagerConfig(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusGetAlertmanagerConfig(ctx)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusGetNamespace(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	return f.handleRouteConvertPrometheusGetNamespace(ctx, namespaceTitleParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusGetRuleGroup(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	groupParam := web.Params(ctx.Req)[":Group"]
	return f.handleRouteConvertPrometheusGetRuleGroup(ctx, namespaceTitleParam, groupParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusGetRules(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusGetRules(ctx)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusPostAlertmanagerConfig(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusPostAlertmanagerConfig(ctx)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusPostRuleGroup(ctx *contextmodel.ReqContext) response.Response {
	// Parse Path Parameters
	namespaceTitleParam := web.Params(ctx.Req)[":NamespaceTitle"]
	return f.handleRouteConvertPrometheusPostRuleGroup(ctx, namespaceTitleParam)
}
func (f *ConvertPrometheusApiHandler) RouteConvertPrometheusPostRuleGroups(ctx *contextmodel.ReqContext) response.Response {
	return f.handleRouteConvertPrometheusPostRuleGroups(ctx)
}

func (api *API) RegisterConvertPrometheusApiEndpoints(srv ConvertPrometheusApi, m *metrics.API) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Delete(
			toMacaronPath("/api/convert/api/prom/rules/{NamespaceTitle}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodDelete, "/api/convert/api/prom/rules/{NamespaceTitle}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/convert/api/prom/rules/{NamespaceTitle}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexDeleteNamespace),
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/convert/api/prom/rules/{NamespaceTitle}/{Group}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodDelete, "/api/convert/api/prom/rules/{NamespaceTitle}/{Group}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/convert/api/prom/rules/{NamespaceTitle}/{Group}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexDeleteRuleGroup),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/api/prom/rules/{NamespaceTitle}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/api/prom/rules/{NamespaceTitle}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/api/prom/rules/{NamespaceTitle}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexGetNamespace),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/api/prom/rules/{NamespaceTitle}/{Group}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/api/prom/rules/{NamespaceTitle}/{Group}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/api/prom/rules/{NamespaceTitle}/{Group}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexGetRuleGroup),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/api/prom/rules"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/api/prom/rules"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/api/prom/rules",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexGetRules),
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/convert/api/prom/rules/{NamespaceTitle}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodPost, "/api/convert/api/prom/rules/{NamespaceTitle}"),
			metrics.Instrument(
				http.MethodPost,
				"/api/convert/api/prom/rules/{NamespaceTitle}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexPostRuleGroup),
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/convert/api/prom/rules"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodPost, "/api/convert/api/prom/rules"),
			metrics.Instrument(
				http.MethodPost,
				"/api/convert/api/prom/rules",
				api.Hooks.Wrap(srv.RouteConvertPrometheusCortexPostRuleGroups),
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/convert/api/v1/alerts"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodDelete, "/api/convert/api/v1/alerts"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/convert/api/v1/alerts",
				api.Hooks.Wrap(srv.RouteConvertPrometheusDeleteAlertmanagerConfig),
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/convert/prometheus/config/v1/rules/{NamespaceTitle}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodDelete, "/api/convert/prometheus/config/v1/rules/{NamespaceTitle}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/convert/prometheus/config/v1/rules/{NamespaceTitle}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusDeleteNamespace),
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/convert/prometheus/config/v1/rules/{NamespaceTitle}/{Group}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodDelete, "/api/convert/prometheus/config/v1/rules/{NamespaceTitle}/{Group}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/convert/prometheus/config/v1/rules/{NamespaceTitle}/{Group}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusDeleteRuleGroup),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/api/v1/alerts"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/api/v1/alerts"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/api/v1/alerts",
				api.Hooks.Wrap(srv.RouteConvertPrometheusGetAlertmanagerConfig),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/prometheus/config/v1/rules/{NamespaceTitle}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/prometheus/config/v1/rules/{NamespaceTitle}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/prometheus/config/v1/rules/{NamespaceTitle}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusGetNamespace),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/prometheus/config/v1/rules/{NamespaceTitle}/{Group}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/prometheus/config/v1/rules/{NamespaceTitle}/{Group}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/prometheus/config/v1/rules/{NamespaceTitle}/{Group}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusGetRuleGroup),
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/convert/prometheus/config/v1/rules"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodGet, "/api/convert/prometheus/config/v1/rules"),
			metrics.Instrument(
				http.MethodGet,
				"/api/convert/prometheus/config/v1/rules",
				api.Hooks.Wrap(srv.RouteConvertPrometheusGetRules),
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/convert/api/v1/alerts"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodPost, "/api/convert/api/v1/alerts"),
			metrics.Instrument(
				http.MethodPost,
				"/api/convert/api/v1/alerts",
				api.Hooks.Wrap(srv.RouteConvertPrometheusPostAlertmanagerConfig),
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/convert/prometheus/config/v1/rules/{NamespaceTitle}"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodPost, "/api/convert/prometheus/config/v1/rules/{NamespaceTitle}"),
			metrics.Instrument(
				http.MethodPost,
				"/api/convert/prometheus/config/v1/rules/{NamespaceTitle}",
				api.Hooks.Wrap(srv.RouteConvertPrometheusPostRuleGroup),
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/convert/prometheus/config/v1/rules"),
			requestmeta.SetOwner(requestmeta.TeamAlerting),
			requestmeta.SetSLOGroup(requestmeta.SLOGroupHighSlow),
			api.authorize(http.MethodPost, "/api/convert/prometheus/config/v1/rules"),
			metrics.Instrument(
				http.MethodPost,
				"/api/convert/prometheus/config/v1/rules",
				api.Hooks.Wrap(srv.RouteConvertPrometheusPostRuleGroups),
				m,
			),
		)
	}, middleware.ReqSignedIn)
}
