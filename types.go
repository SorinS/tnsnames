package main

import (
	"bramp.net/antlr4/tnsnames"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	guuid "github.com/google/uuid"
	"strings"
)

type Address struct {
	Id       string
	Host     string
	Port     string
	Protocol string
}

// CustomTnsNamesListener is a complete listener for a parse tree produced by tnsnamesParser.
type CustomTnsNamesListener struct {
	Addresses    []Address
	CrtAddress   Address
	OnNewAddress bool
}

func ToValue(kv string) string {
	s := strings.Split(kv, "=")
	if len(s) > 1 {
		return strings.TrimSuffix(s[1], ")")
	}
	return kv
}

// VisitTerminal is called when a terminal node is visited.
func (s *CustomTnsNamesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CustomTnsNamesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *CustomTnsNamesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//	fmt.Printf("%s\n", ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (s *CustomTnsNamesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTnsnames is called when production tnsnames is entered.
func (s *CustomTnsNamesListener) EnterTnsnames(ctx *tnsnames.TnsnamesContext) {}

// ExitTnsnames is called when production tnsnames is exited.
func (s *CustomTnsNamesListener) ExitTnsnames(ctx *tnsnames.TnsnamesContext) {}

// EnterTns_entry is called when production tns_entry is entered.
func (s *CustomTnsNamesListener) EnterTns_entry(ctx *tnsnames.Tns_entryContext) {}

// ExitTns_entry is called when production tns_entry is exited.
func (s *CustomTnsNamesListener) ExitTns_entry(ctx *tnsnames.Tns_entryContext) {}

// EnterIfile is called when production ifile is entered.
func (s *CustomTnsNamesListener) EnterIfile(ctx *tnsnames.IfileContext) {}

// ExitIfile is called when production ifile is exited.
func (s *CustomTnsNamesListener) ExitIfile(ctx *tnsnames.IfileContext) {}

// EnterLsnr_entry is called when production lsnr_entry is entered.
func (s *CustomTnsNamesListener) EnterLsnr_entry(ctx *tnsnames.Lsnr_entryContext) {}

// ExitLsnr_entry is called when production lsnr_entry is exited.
func (s *CustomTnsNamesListener) ExitLsnr_entry(ctx *tnsnames.Lsnr_entryContext) {}

// EnterLsnr_description is called when production lsnr_description is entered.
func (s *CustomTnsNamesListener) EnterLsnr_description(ctx *tnsnames.Lsnr_descriptionContext) {}

// ExitLsnr_description is called when production lsnr_description is exited.
func (s *CustomTnsNamesListener) ExitLsnr_description(ctx *tnsnames.Lsnr_descriptionContext) {}

// EnterAlias_list is called when production alias_list is entered.
func (s *CustomTnsNamesListener) EnterAlias_list(ctx *tnsnames.Alias_listContext) {}

// ExitAlias_list is called when production alias_list is exited.
func (s *CustomTnsNamesListener) ExitAlias_list(ctx *tnsnames.Alias_listContext) {}

// EnterAlias is called when production alias is entered.
func (s *CustomTnsNamesListener) EnterAlias(ctx *tnsnames.AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *CustomTnsNamesListener) ExitAlias(ctx *tnsnames.AliasContext) {}

// EnterDescription_list is called when production description_list is entered.
func (s *CustomTnsNamesListener) EnterDescription_list(ctx *tnsnames.Description_listContext) {}

// ExitDescription_list is called when production description_list is exited.
func (s *CustomTnsNamesListener) ExitDescription_list(ctx *tnsnames.Description_listContext) {}

// EnterDl_params is called when production dl_params is entered.
func (s *CustomTnsNamesListener) EnterDl_params(ctx *tnsnames.Dl_paramsContext) {}

// ExitDl_params is called when production dl_params is exited.
func (s *CustomTnsNamesListener) ExitDl_params(ctx *tnsnames.Dl_paramsContext) {}

// EnterDl_parameter is called when production dl_parameter is entered.
func (s *CustomTnsNamesListener) EnterDl_parameter(ctx *tnsnames.Dl_parameterContext) {}

// ExitDl_parameter is called when production dl_parameter is exited.
func (s *CustomTnsNamesListener) ExitDl_parameter(ctx *tnsnames.Dl_parameterContext) {}

// EnterDescription is called when production description is entered.
func (s *CustomTnsNamesListener) EnterDescription(ctx *tnsnames.DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *CustomTnsNamesListener) ExitDescription(ctx *tnsnames.DescriptionContext) {}

// EnterD_params is called when production d_params is entered.
func (s *CustomTnsNamesListener) EnterD_params(ctx *tnsnames.D_paramsContext) {}

// ExitD_params is called when production d_params is exited.
func (s *CustomTnsNamesListener) ExitD_params(ctx *tnsnames.D_paramsContext) {}

// EnterD_parameter is called when production d_parameter is entered.
func (s *CustomTnsNamesListener) EnterD_parameter(ctx *tnsnames.D_parameterContext) {}

// ExitD_parameter is called when production d_parameter is exited.
func (s *CustomTnsNamesListener) ExitD_parameter(ctx *tnsnames.D_parameterContext) {}

// EnterD_enable is called when production d_enable is entered.
func (s *CustomTnsNamesListener) EnterD_enable(ctx *tnsnames.D_enableContext) {}

// ExitD_enable is called when production d_enable is exited.
func (s *CustomTnsNamesListener) ExitD_enable(ctx *tnsnames.D_enableContext) {}

// EnterD_sdu is called when production d_sdu is entered.
func (s *CustomTnsNamesListener) EnterD_sdu(ctx *tnsnames.D_sduContext) {}

// ExitD_sdu is called when production d_sdu is exited.
func (s *CustomTnsNamesListener) ExitD_sdu(ctx *tnsnames.D_sduContext) {}

// EnterD_recv_buf is called when production d_recv_buf is entered.
func (s *CustomTnsNamesListener) EnterD_recv_buf(ctx *tnsnames.D_recv_bufContext) {}

// ExitD_recv_buf is called when production d_recv_buf is exited.
func (s *CustomTnsNamesListener) ExitD_recv_buf(ctx *tnsnames.D_recv_bufContext) {}

// EnterD_send_buf is called when production d_send_buf is entered.
func (s *CustomTnsNamesListener) EnterD_send_buf(ctx *tnsnames.D_send_bufContext) {}

// ExitD_send_buf is called when production d_send_buf is exited.
func (s *CustomTnsNamesListener) ExitD_send_buf(ctx *tnsnames.D_send_bufContext) {}

// EnterD_service_type is called when production d_service_type is entered.
func (s *CustomTnsNamesListener) EnterD_service_type(ctx *tnsnames.D_service_typeContext) {}

// ExitD_service_type is called when production d_service_type is exited.
func (s *CustomTnsNamesListener) ExitD_service_type(ctx *tnsnames.D_service_typeContext) {}

// EnterD_security is called when production d_security is entered.
func (s *CustomTnsNamesListener) EnterD_security(ctx *tnsnames.D_securityContext) {}

// ExitD_security is called when production d_security is exited.
func (s *CustomTnsNamesListener) ExitD_security(ctx *tnsnames.D_securityContext) {}

// EnterD_conn_timeout is called when production d_conn_timeout is entered.
func (s *CustomTnsNamesListener) EnterD_conn_timeout(ctx *tnsnames.D_conn_timeoutContext) {}

// ExitD_conn_timeout is called when production d_conn_timeout is exited.
func (s *CustomTnsNamesListener) ExitD_conn_timeout(ctx *tnsnames.D_conn_timeoutContext) {}

// EnterD_retry_count is called when production d_retry_count is entered.
func (s *CustomTnsNamesListener) EnterD_retry_count(ctx *tnsnames.D_retry_countContext) {}

// ExitD_retry_count is called when production d_retry_count is exited.
func (s *CustomTnsNamesListener) ExitD_retry_count(ctx *tnsnames.D_retry_countContext) {}

// EnterD_tct is called when production d_tct is entered.
func (s *CustomTnsNamesListener) EnterD_tct(ctx *tnsnames.D_tctContext) {}

// ExitD_tct is called when production d_tct is exited.
func (s *CustomTnsNamesListener) ExitD_tct(ctx *tnsnames.D_tctContext) {}

// EnterDs_parameter is called when production ds_parameter is entered.
func (s *CustomTnsNamesListener) EnterDs_parameter(ctx *tnsnames.Ds_parameterContext) {}

// ExitDs_parameter is called when production ds_parameter is exited.
func (s *CustomTnsNamesListener) ExitDs_parameter(ctx *tnsnames.Ds_parameterContext) {}

// EnterAddress_list is called when production address_list is entered.
func (s *CustomTnsNamesListener) EnterAddress_list(ctx *tnsnames.Address_listContext) {}

// ExitAddress_list is called when production address_list is exited.
func (s *CustomTnsNamesListener) ExitAddress_list(ctx *tnsnames.Address_listContext) {}

// EnterAl_params is called when production al_params is entered.
func (s *CustomTnsNamesListener) EnterAl_params(ctx *tnsnames.Al_paramsContext) {}

// ExitAl_params is called when production al_params is exited.
func (s *CustomTnsNamesListener) ExitAl_params(ctx *tnsnames.Al_paramsContext) {}

// EnterAl_parameter is called when production al_parameter is entered.
func (s *CustomTnsNamesListener) EnterAl_parameter(ctx *tnsnames.Al_parameterContext) {}

// ExitAl_parameter is called when production al_parameter is exited.
func (s *CustomTnsNamesListener) ExitAl_parameter(ctx *tnsnames.Al_parameterContext) {}

// EnterAl_failover is called when production al_failover is entered.
func (s *CustomTnsNamesListener) EnterAl_failover(ctx *tnsnames.Al_failoverContext) {}

// ExitAl_failover is called when production al_failover is exited.
func (s *CustomTnsNamesListener) ExitAl_failover(ctx *tnsnames.Al_failoverContext) {}

// EnterAl_load_balance is called when production al_load_balance is entered.
func (s *CustomTnsNamesListener) EnterAl_load_balance(ctx *tnsnames.Al_load_balanceContext) {}

// ExitAl_load_balance is called when production al_load_balance is exited.
func (s *CustomTnsNamesListener) ExitAl_load_balance(ctx *tnsnames.Al_load_balanceContext) {}

// EnterAl_source_route is called when production al_source_route is entered.
func (s *CustomTnsNamesListener) EnterAl_source_route(ctx *tnsnames.Al_source_routeContext) {}

// ExitAl_source_route is called when production al_source_route is exited.
func (s *CustomTnsNamesListener) ExitAl_source_route(ctx *tnsnames.Al_source_routeContext) {}

// EnterAddress is called when production address is entered.
func (s *CustomTnsNamesListener) EnterAddress(ctx *tnsnames.AddressContext) {
	s.OnNewAddress = true
	s.CrtAddress = Address{
		Id:       guuid.New().String(),
		Host:     "",
		Port:     "",
		Protocol: "oracle",
	}
}

// ExitAddress is called when production address is exited.
func (s *CustomTnsNamesListener) ExitAddress(ctx *tnsnames.AddressContext) {
	s.OnNewAddress = false
	s.Addresses = append(s.Addresses, s.CrtAddress)
}

// EnterA_params is called when production a_params is entered.
func (s *CustomTnsNamesListener) EnterA_params(ctx *tnsnames.A_paramsContext) {}

// ExitA_params is called when production a_params is exited.
func (s *CustomTnsNamesListener) ExitA_params(ctx *tnsnames.A_paramsContext) {}

// EnterA_parameter is called when production a_parameter is entered.
func (s *CustomTnsNamesListener) EnterA_parameter(ctx *tnsnames.A_parameterContext) {}

// ExitA_parameter is called when production a_parameter is exited.
func (s *CustomTnsNamesListener) ExitA_parameter(ctx *tnsnames.A_parameterContext) {}

// EnterProtocol_info is called when production protocol_info is entered.
func (s *CustomTnsNamesListener) EnterProtocol_info(ctx *tnsnames.Protocol_infoContext) {}

// ExitProtocol_info is called when production protocol_info is exited.
func (s *CustomTnsNamesListener) ExitProtocol_info(ctx *tnsnames.Protocol_infoContext) {}

// EnterTcp_protocol is called when production tcp_protocol is entered.
func (s *CustomTnsNamesListener) EnterTcp_protocol(ctx *tnsnames.Tcp_protocolContext) {}

// ExitTcp_protocol is called when production tcp_protocol is exited.
func (s *CustomTnsNamesListener) ExitTcp_protocol(ctx *tnsnames.Tcp_protocolContext) {}

// EnterTcp_params is called when production tcp_params is entered.
func (s *CustomTnsNamesListener) EnterTcp_params(ctx *tnsnames.Tcp_paramsContext) {}

// ExitTcp_params is called when production tcp_params is exited.
func (s *CustomTnsNamesListener) ExitTcp_params(ctx *tnsnames.Tcp_paramsContext) {}

// EnterTcp_parameter is called when production tcp_parameter is entered.
func (s *CustomTnsNamesListener) EnterTcp_parameter(ctx *tnsnames.Tcp_parameterContext) {}

// ExitTcp_parameter is called when production tcp_parameter is exited.
func (s *CustomTnsNamesListener) ExitTcp_parameter(ctx *tnsnames.Tcp_parameterContext) {}

// EnterTcp_host is called when production tcp_host is entered.
func (s *CustomTnsNamesListener) EnterTcp_host(ctx *tnsnames.Tcp_hostContext) {
	if s.OnNewAddress {
		s.CrtAddress.Host = ToValue(ctx.GetText())
	}
}

// ExitTcp_host is called when production tcp_host is exited.
func (s *CustomTnsNamesListener) ExitTcp_host(ctx *tnsnames.Tcp_hostContext) {}

// EnterTcp_port is called when production tcp_port is entered.
func (s *CustomTnsNamesListener) EnterTcp_port(ctx *tnsnames.Tcp_portContext) {
	if s.OnNewAddress {
		s.CrtAddress.Port = ToValue(ctx.GetText())
	}
}

// ExitTcp_port is called when production tcp_port is exited.
func (s *CustomTnsNamesListener) ExitTcp_port(ctx *tnsnames.Tcp_portContext) {}

// EnterTcp_tcp is called when production tcp_tcp is entered.
func (s *CustomTnsNamesListener) EnterTcp_tcp(ctx *tnsnames.Tcp_tcpContext) {}

// ExitTcp_tcp is called when production tcp_tcp is exited.
func (s *CustomTnsNamesListener) ExitTcp_tcp(ctx *tnsnames.Tcp_tcpContext) {}

// EnterHost is called when production host is entered.
func (s *CustomTnsNamesListener) EnterHost(ctx *tnsnames.HostContext) {}

// ExitHost is called when production host is exited.
func (s *CustomTnsNamesListener) ExitHost(ctx *tnsnames.HostContext) {}

// EnterPort is called when production port is entered.
func (s *CustomTnsNamesListener) EnterPort(ctx *tnsnames.PortContext) {}

// ExitPort is called when production port is exited.
func (s *CustomTnsNamesListener) ExitPort(ctx *tnsnames.PortContext) {}

// EnterIpc_protocol is called when production ipc_protocol is entered.
func (s *CustomTnsNamesListener) EnterIpc_protocol(ctx *tnsnames.Ipc_protocolContext) {}

// ExitIpc_protocol is called when production ipc_protocol is exited.
func (s *CustomTnsNamesListener) ExitIpc_protocol(ctx *tnsnames.Ipc_protocolContext) {}

// EnterIpc_params is called when production ipc_params is entered.
func (s *CustomTnsNamesListener) EnterIpc_params(ctx *tnsnames.Ipc_paramsContext) {}

// ExitIpc_params is called when production ipc_params is exited.
func (s *CustomTnsNamesListener) ExitIpc_params(ctx *tnsnames.Ipc_paramsContext) {}

// EnterIpc_parameter is called when production ipc_parameter is entered.
func (s *CustomTnsNamesListener) EnterIpc_parameter(ctx *tnsnames.Ipc_parameterContext) {}

// ExitIpc_parameter is called when production ipc_parameter is exited.
func (s *CustomTnsNamesListener) ExitIpc_parameter(ctx *tnsnames.Ipc_parameterContext) {}

// EnterIpc_ipc is called when production ipc_ipc is entered.
func (s *CustomTnsNamesListener) EnterIpc_ipc(ctx *tnsnames.Ipc_ipcContext) {}

// ExitIpc_ipc is called when production ipc_ipc is exited.
func (s *CustomTnsNamesListener) ExitIpc_ipc(ctx *tnsnames.Ipc_ipcContext) {}

// EnterIpc_key is called when production ipc_key is entered.
func (s *CustomTnsNamesListener) EnterIpc_key(ctx *tnsnames.Ipc_keyContext) {}

// ExitIpc_key is called when production ipc_key is exited.
func (s *CustomTnsNamesListener) ExitIpc_key(ctx *tnsnames.Ipc_keyContext) {}

// EnterSpx_protocol is called when production spx_protocol is entered.
func (s *CustomTnsNamesListener) EnterSpx_protocol(ctx *tnsnames.Spx_protocolContext) {}

// ExitSpx_protocol is called when production spx_protocol is exited.
func (s *CustomTnsNamesListener) ExitSpx_protocol(ctx *tnsnames.Spx_protocolContext) {}

// EnterSpx_params is called when production spx_params is entered.
func (s *CustomTnsNamesListener) EnterSpx_params(ctx *tnsnames.Spx_paramsContext) {}

// ExitSpx_params is called when production spx_params is exited.
func (s *CustomTnsNamesListener) ExitSpx_params(ctx *tnsnames.Spx_paramsContext) {}

// EnterSpx_parameter is called when production spx_parameter is entered.
func (s *CustomTnsNamesListener) EnterSpx_parameter(ctx *tnsnames.Spx_parameterContext) {}

// ExitSpx_parameter is called when production spx_parameter is exited.
func (s *CustomTnsNamesListener) ExitSpx_parameter(ctx *tnsnames.Spx_parameterContext) {}

// EnterSpx_spx is called when production spx_spx is entered.
func (s *CustomTnsNamesListener) EnterSpx_spx(ctx *tnsnames.Spx_spxContext) {}

// ExitSpx_spx is called when production spx_spx is exited.
func (s *CustomTnsNamesListener) ExitSpx_spx(ctx *tnsnames.Spx_spxContext) {}

// EnterSpx_service is called when production spx_service is entered.
func (s *CustomTnsNamesListener) EnterSpx_service(ctx *tnsnames.Spx_serviceContext) {}

// ExitSpx_service is called when production spx_service is exited.
func (s *CustomTnsNamesListener) ExitSpx_service(ctx *tnsnames.Spx_serviceContext) {}

// EnterNmp_protocol is called when production nmp_protocol is entered.
func (s *CustomTnsNamesListener) EnterNmp_protocol(ctx *tnsnames.Nmp_protocolContext) {}

// ExitNmp_protocol is called when production nmp_protocol is exited.
func (s *CustomTnsNamesListener) ExitNmp_protocol(ctx *tnsnames.Nmp_protocolContext) {}

// EnterNmp_params is called when production nmp_params is entered.
func (s *CustomTnsNamesListener) EnterNmp_params(ctx *tnsnames.Nmp_paramsContext) {}

// ExitNmp_params is called when production nmp_params is exited.
func (s *CustomTnsNamesListener) ExitNmp_params(ctx *tnsnames.Nmp_paramsContext) {}

// EnterNmp_parameter is called when production nmp_parameter is entered.
func (s *CustomTnsNamesListener) EnterNmp_parameter(ctx *tnsnames.Nmp_parameterContext) {}

// ExitNmp_parameter is called when production nmp_parameter is exited.
func (s *CustomTnsNamesListener) ExitNmp_parameter(ctx *tnsnames.Nmp_parameterContext) {}

// EnterNmp_nmp is called when production nmp_nmp is entered.
func (s *CustomTnsNamesListener) EnterNmp_nmp(ctx *tnsnames.Nmp_nmpContext) {}

// ExitNmp_nmp is called when production nmp_nmp is exited.
func (s *CustomTnsNamesListener) ExitNmp_nmp(ctx *tnsnames.Nmp_nmpContext) {}

// EnterNmp_server is called when production nmp_server is entered.
func (s *CustomTnsNamesListener) EnterNmp_server(ctx *tnsnames.Nmp_serverContext) {}

// ExitNmp_server is called when production nmp_server is exited.
func (s *CustomTnsNamesListener) ExitNmp_server(ctx *tnsnames.Nmp_serverContext) {}

// EnterNmp_pipe is called when production nmp_pipe is entered.
func (s *CustomTnsNamesListener) EnterNmp_pipe(ctx *tnsnames.Nmp_pipeContext) {}

// ExitNmp_pipe is called when production nmp_pipe is exited.
func (s *CustomTnsNamesListener) ExitNmp_pipe(ctx *tnsnames.Nmp_pipeContext) {}

// EnterBeq_protocol is called when production beq_protocol is entered.
func (s *CustomTnsNamesListener) EnterBeq_protocol(ctx *tnsnames.Beq_protocolContext) {}

// ExitBeq_protocol is called when production beq_protocol is exited.
func (s *CustomTnsNamesListener) ExitBeq_protocol(ctx *tnsnames.Beq_protocolContext) {}

// EnterBeq_params is called when production beq_params is entered.
func (s *CustomTnsNamesListener) EnterBeq_params(ctx *tnsnames.Beq_paramsContext) {}

// ExitBeq_params is called when production beq_params is exited.
func (s *CustomTnsNamesListener) ExitBeq_params(ctx *tnsnames.Beq_paramsContext) {}

// EnterBeq_parameter is called when production beq_parameter is entered.
func (s *CustomTnsNamesListener) EnterBeq_parameter(ctx *tnsnames.Beq_parameterContext) {}

// ExitBeq_parameter is called when production beq_parameter is exited.
func (s *CustomTnsNamesListener) ExitBeq_parameter(ctx *tnsnames.Beq_parameterContext) {}

// EnterBeq_beq is called when production beq_beq is entered.
func (s *CustomTnsNamesListener) EnterBeq_beq(ctx *tnsnames.Beq_beqContext) {}

// ExitBeq_beq is called when production beq_beq is exited.
func (s *CustomTnsNamesListener) ExitBeq_beq(ctx *tnsnames.Beq_beqContext) {}

// EnterBeq_program is called when production beq_program is entered.
func (s *CustomTnsNamesListener) EnterBeq_program(ctx *tnsnames.Beq_programContext) {}

// ExitBeq_program is called when production beq_program is exited.
func (s *CustomTnsNamesListener) ExitBeq_program(ctx *tnsnames.Beq_programContext) {}

// EnterBeq_argv0 is called when production beq_argv0 is entered.
func (s *CustomTnsNamesListener) EnterBeq_argv0(ctx *tnsnames.Beq_argv0Context) {}

// ExitBeq_argv0 is called when production beq_argv0 is exited.
func (s *CustomTnsNamesListener) ExitBeq_argv0(ctx *tnsnames.Beq_argv0Context) {}

// EnterBeq_args is called when production beq_args is entered.
func (s *CustomTnsNamesListener) EnterBeq_args(ctx *tnsnames.Beq_argsContext) {}

// ExitBeq_args is called when production beq_args is exited.
func (s *CustomTnsNamesListener) ExitBeq_args(ctx *tnsnames.Beq_argsContext) {}

// EnterBa_parameter is called when production ba_parameter is entered.
func (s *CustomTnsNamesListener) EnterBa_parameter(ctx *tnsnames.Ba_parameterContext) {}

// ExitBa_parameter is called when production ba_parameter is exited.
func (s *CustomTnsNamesListener) ExitBa_parameter(ctx *tnsnames.Ba_parameterContext) {}

// EnterBa_description is called when production ba_description is entered.
func (s *CustomTnsNamesListener) EnterBa_description(ctx *tnsnames.Ba_descriptionContext) {}

// ExitBa_description is called when production ba_description is exited.
func (s *CustomTnsNamesListener) ExitBa_description(ctx *tnsnames.Ba_descriptionContext) {}

// EnterBad_params is called when production bad_params is entered.
func (s *CustomTnsNamesListener) EnterBad_params(ctx *tnsnames.Bad_paramsContext) {}

// ExitBad_params is called when production bad_params is exited.
func (s *CustomTnsNamesListener) ExitBad_params(ctx *tnsnames.Bad_paramsContext) {}

// EnterBad_parameter is called when production bad_parameter is entered.
func (s *CustomTnsNamesListener) EnterBad_parameter(ctx *tnsnames.Bad_parameterContext) {}

// ExitBad_parameter is called when production bad_parameter is exited.
func (s *CustomTnsNamesListener) ExitBad_parameter(ctx *tnsnames.Bad_parameterContext) {}

// EnterBad_local is called when production bad_local is entered.
func (s *CustomTnsNamesListener) EnterBad_local(ctx *tnsnames.Bad_localContext) {}

// ExitBad_local is called when production bad_local is exited.
func (s *CustomTnsNamesListener) ExitBad_local(ctx *tnsnames.Bad_localContext) {}

// EnterBad_address is called when production bad_address is entered.
func (s *CustomTnsNamesListener) EnterBad_address(ctx *tnsnames.Bad_addressContext) {}

// ExitBad_address is called when production bad_address is exited.
func (s *CustomTnsNamesListener) ExitBad_address(ctx *tnsnames.Bad_addressContext) {}

// EnterConnect_data is called when production connect_data is entered.
func (s *CustomTnsNamesListener) EnterConnect_data(ctx *tnsnames.Connect_dataContext) {}

// ExitConnect_data is called when production connect_data is exited.
func (s *CustomTnsNamesListener) ExitConnect_data(ctx *tnsnames.Connect_dataContext) {}

// EnterCd_params is called when production cd_params is entered.
func (s *CustomTnsNamesListener) EnterCd_params(ctx *tnsnames.Cd_paramsContext) {}

// ExitCd_params is called when production cd_params is exited.
func (s *CustomTnsNamesListener) ExitCd_params(ctx *tnsnames.Cd_paramsContext) {}

// EnterCd_parameter is called when production cd_parameter is entered.
func (s *CustomTnsNamesListener) EnterCd_parameter(ctx *tnsnames.Cd_parameterContext) {}

// ExitCd_parameter is called when production cd_parameter is exited.
func (s *CustomTnsNamesListener) ExitCd_parameter(ctx *tnsnames.Cd_parameterContext) {}

// EnterCd_service_name is called when production cd_service_name is entered.
func (s *CustomTnsNamesListener) EnterCd_service_name(ctx *tnsnames.Cd_service_nameContext) {}

// ExitCd_service_name is called when production cd_service_name is exited.
func (s *CustomTnsNamesListener) ExitCd_service_name(ctx *tnsnames.Cd_service_nameContext) {}

// EnterCd_sid is called when production cd_sid is entered.
func (s *CustomTnsNamesListener) EnterCd_sid(ctx *tnsnames.Cd_sidContext) {}

// ExitCd_sid is called when production cd_sid is exited.
func (s *CustomTnsNamesListener) ExitCd_sid(ctx *tnsnames.Cd_sidContext) {}

// EnterCd_instance_name is called when production cd_instance_name is entered.
func (s *CustomTnsNamesListener) EnterCd_instance_name(ctx *tnsnames.Cd_instance_nameContext) {}

// ExitCd_instance_name is called when production cd_instance_name is exited.
func (s *CustomTnsNamesListener) ExitCd_instance_name(ctx *tnsnames.Cd_instance_nameContext) {}

// EnterCd_failover_mode is called when production cd_failover_mode is entered.
func (s *CustomTnsNamesListener) EnterCd_failover_mode(ctx *tnsnames.Cd_failover_modeContext) {}

// ExitCd_failover_mode is called when production cd_failover_mode is exited.
func (s *CustomTnsNamesListener) ExitCd_failover_mode(ctx *tnsnames.Cd_failover_modeContext) {}

// EnterCd_global_name is called when production cd_global_name is entered.
func (s *CustomTnsNamesListener) EnterCd_global_name(ctx *tnsnames.Cd_global_nameContext) {}

// ExitCd_global_name is called when production cd_global_name is exited.
func (s *CustomTnsNamesListener) ExitCd_global_name(ctx *tnsnames.Cd_global_nameContext) {}

// EnterCd_hs is called when production cd_hs is entered.
func (s *CustomTnsNamesListener) EnterCd_hs(ctx *tnsnames.Cd_hsContext) {}

// ExitCd_hs is called when production cd_hs is exited.
func (s *CustomTnsNamesListener) ExitCd_hs(ctx *tnsnames.Cd_hsContext) {}

// EnterCd_rdb_database is called when production cd_rdb_database is entered.
func (s *CustomTnsNamesListener) EnterCd_rdb_database(ctx *tnsnames.Cd_rdb_databaseContext) {}

// ExitCd_rdb_database is called when production cd_rdb_database is exited.
func (s *CustomTnsNamesListener) ExitCd_rdb_database(ctx *tnsnames.Cd_rdb_databaseContext) {}

// EnterCd_server is called when production cd_server is entered.
func (s *CustomTnsNamesListener) EnterCd_server(ctx *tnsnames.Cd_serverContext) {}

// ExitCd_server is called when production cd_server is exited.
func (s *CustomTnsNamesListener) ExitCd_server(ctx *tnsnames.Cd_serverContext) {}

// EnterCd_ur is called when production cd_ur is entered.
func (s *CustomTnsNamesListener) EnterCd_ur(ctx *tnsnames.Cd_urContext) {}

// ExitCd_ur is called when production cd_ur is exited.
func (s *CustomTnsNamesListener) ExitCd_ur(ctx *tnsnames.Cd_urContext) {}

// EnterFo_params is called when production fo_params is entered.
func (s *CustomTnsNamesListener) EnterFo_params(ctx *tnsnames.Fo_paramsContext) {}

// ExitFo_params is called when production fo_params is exited.
func (s *CustomTnsNamesListener) ExitFo_params(ctx *tnsnames.Fo_paramsContext) {}

// EnterFo_parameter is called when production fo_parameter is entered.
func (s *CustomTnsNamesListener) EnterFo_parameter(ctx *tnsnames.Fo_parameterContext) {}

// ExitFo_parameter is called when production fo_parameter is exited.
func (s *CustomTnsNamesListener) ExitFo_parameter(ctx *tnsnames.Fo_parameterContext) {}

// EnterFo_type is called when production fo_type is entered.
func (s *CustomTnsNamesListener) EnterFo_type(ctx *tnsnames.Fo_typeContext) {}

// ExitFo_type is called when production fo_type is exited.
func (s *CustomTnsNamesListener) ExitFo_type(ctx *tnsnames.Fo_typeContext) {}

// EnterFo_backup is called when production fo_backup is entered.
func (s *CustomTnsNamesListener) EnterFo_backup(ctx *tnsnames.Fo_backupContext) {}

// ExitFo_backup is called when production fo_backup is exited.
func (s *CustomTnsNamesListener) ExitFo_backup(ctx *tnsnames.Fo_backupContext) {}

// EnterFo_method is called when production fo_method is entered.
func (s *CustomTnsNamesListener) EnterFo_method(ctx *tnsnames.Fo_methodContext) {}

// ExitFo_method is called when production fo_method is exited.
func (s *CustomTnsNamesListener) ExitFo_method(ctx *tnsnames.Fo_methodContext) {}

// EnterFo_retries is called when production fo_retries is entered.
func (s *CustomTnsNamesListener) EnterFo_retries(ctx *tnsnames.Fo_retriesContext) {}

// ExitFo_retries is called when production fo_retries is exited.
func (s *CustomTnsNamesListener) ExitFo_retries(ctx *tnsnames.Fo_retriesContext) {}

// EnterFo_delay is called when production fo_delay is entered.
func (s *CustomTnsNamesListener) EnterFo_delay(ctx *tnsnames.Fo_delayContext) {}

// ExitFo_delay is called when production fo_delay is exited.
func (s *CustomTnsNamesListener) ExitFo_delay(ctx *tnsnames.Fo_delayContext) {}
