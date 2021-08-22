package servicenow

type IncidentResultDetail struct {
	Result struct {
		Active                 string `json:"active,omitempty"`
		ActivityDue            string `json:"activity_due,omitempty"`
		AdditionalAssigneeList string `json:"additional_assignee_list,omitempty"`
		Approval               string `json:"approval,omitempty"`
		ApprovalHistory        string `json:"approval_history,omitempty"`
		ApprovalSet            string `json:"approval_set,omitempty"`
		AssignedTo             string `json:"assigned_to,omitempty"`
		AssignmentGroup        struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"assignment_group,omitempty"`
		BusinessDuration string `json:"business_duration,omitempty"`
		BusinessService  string `json:"business_service,omitempty"`
		BusinessStc      string `json:"business_stc,omitempty"`
		CalendarDuration string `json:"calendar_duration,omitempty"`
		CalendarStc      string `json:"calendar_stc,omitempty"`
		CallerID         string `json:"caller_id,omitempty"`
		Category         string `json:"category,omitempty"`
		CausedBy         string `json:"caused_by,omitempty"`
		ChildIncidents   string `json:"child_incidents,omitempty"`
		CloseCode        string `json:"close_code,omitempty"`
		ClosedAt         string `json:"closed_at,omitempty"`
		ClosedBy         string `json:"closed_by,omitempty"`
		CloseNotes       string `json:"close_notes,omitempty"`
		CmdbCi           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"cmdb_ci,omitempty"`
		//CmdbCi	string	`json:"cmdb_ci,omitempty"`
		Comments             string `json:"comments,omitempty"`
		CommentsAndWorkNotes string `json:"comments_and_work_notes,omitempty"`
		Company              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		//Company              string `json:"company,omitempty"`
		ContactType        string `json:"contact_type,omitempty"`
		Contract           string `json:"contract,omitempty"`
		CorrelationDisplay string `json:"correlation_display,omitempty"`
		CorrelationID      string `json:"correlation_id,omitempty"`
		DeliveryPlan       string `json:"delivery_plan,omitempty"`
		DeliveryTask       string `json:"delivery_task,omitempty"`
		Description        string `json:"description,omitempty"`
		DueDate            string `json:"due_date,omitempty"`
		Escalation         string `json:"escalation,omitempty"`
		ExpectedStart      string `json:"expected_start,omitempty"`
		FollowUp           string `json:"follow_up,omitempty"`
		GroupList          string `json:"group_list,omitempty"`
		HoldReason         string `json:"hold_reason,omitempty"`
		Impact             string `json:"impact,omitempty"`
		IncidentState      string `json:"incident_state,omitempty"`
		Knowledge          string `json:"knowledge,omitempty"`
		Location           string `json:"location,omitempty"`
		MadeSLA            string `json:"made_sla,omitempty"`
		Notify             string `json:"notify,omitempty"`
		Number             string `json:"number,omitempty"`
		OpenedAt           string `json:"opened_at,omitempty"`
		OpenedBy           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"opened_by,omitempty"`
		Order             string `json:"order,omitempty"`
		Parent            string `json:"parent,omitempty"`
		ParentIncident    string `json:"parent_incident,omitempty"`
		Priority          string `json:"priority,omitempty"`
		ProblemID         string `json:"problem_id,omitempty"`
		ReassignmentCount string `json:"reassignment_count,omitempty"`
		ReopenCount       string `json:"reopen_count,omitempty"`
		ReopenedBy        string `json:"reopened_by,omitempty"`
		ReopenedTime      string `json:"reopened_time,omitempty"`
		ResolvedAt        string `json:"resolved_at,omitempty"`
		ResolvedBy        string `json:"resolved_by,omitempty"`
		Rfc               string `json:"rfc,omitempty"`
		Severity          string `json:"severity,omitempty"`
		ShortDescription  string `json:"short_description,omitempty"`
		Skills            string `json:"skills,omitempty"`
		SLADue            string `json:"sla_due,omitempty"`
		State             string `json:"state,omitempty"`
		Subcategory       string `json:"subcategory,omitempty"`
		SysClassName      string `json:"sys_class_name,omitempty"`
		SysCreatedBy      string `json:"sys_created_by,omitempty"`
		SysCreatedOn      string `json:"sys_created_on,omitempty"`
		SysDomain         struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath        string `json:"sys_domain_path,omitempty"`
		SysID                string `json:"sys_id,omitempty"`
		SysModCount          string `json:"sys_mod_count,omitempty"`
		SysTags              string `json:"sys_tags,omitempty"`
		SysUpdatedBy         string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn         string `json:"sys_updated_on,omitempty"`
		TaskFor              string `json:"task_for,omitempty"`
		TimeWorked           string `json:"time_worked,omitempty"`
		UponApproval         string `json:"upon_approval,omitempty"`
		UponReject           string `json:"upon_reject,omitempty"`
		Urgency              string `json:"urgency,omitempty"`
		UserInput            string `json:"user_input,omitempty"`
		WatchList            string `json:"watch_list,omitempty"`
		WorkEnd              string `json:"work_end,omitempty"`
		WorkNotes            string `json:"work_notes,omitempty"`
		WorkNotesList        string `json:"work_notes_list,omitempty"`
		WorkStart            string `json:"work_start,omitempty"`
		UPublicCommentsAdded string `json:"u_public_comments_added,omitempty"` // Custom Field
		UVip                 string `json:"u_vip,omitempty"`                   // Custom Field
		UDoesKbNeedUpdated   string `json:"u_does_kb_need_updated,omitempty"`  // Custom Field
		UGlideDateTime1      string `json:"u_glide_date_time_1"`               // Custom Field
		UMajorIncident       string `json:"u_major_incident,omitempty"`        // Custom Field
		UNotificationList    string `json:"u_notification_list,omitempty"`     // Custom Field
		UFootprintsNumber    string `json:"u_footprints_number,omitempty"`     // Custom Field
		USecureComment       string `json:"u_secure_comment,omitempty"`        // Custom Field
		UCustomerVisible     string `json:"u_customer_visible,omitempty"`      // Custom Field
		UReason              string `json:"u_reason,omitempty"`                // Custom Field
		UPublicComment       string `json:"u_public_comment,omitempty"`        // Custom Field
		UHasBreached         string `json:"u_has_breached,omitempty"`          // Custom Field
	} `json:"result,omitempty"`
}

type CaseResultDetail struct {
	Result struct {
		Account struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"account,omitempty"`
		Active                  string `json:"active,omitempty"`
		ActiveAccountEscalation string `json:"active_account_escalation,omitempty"`
		ActiveEscalation        string `json:"active_escalation,omitempty"`
		ActivityDue             string `json:"activity_due,omitempty"`
		AdditionalAssigneeList  string `json:"additional_assignee_list,omitempty"`
		Approval                string `json:"approval,omitempty"`
		ApprovalHistory         string `json:"approval_history,omitempty"`
		ApprovalSet             string `json:"approval_set,omitempty"`
		Asset                   string `json:"asset,omitempty"`
		AssignedTo              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"assigned_to,omitempty"`
		AssignmentGroup struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"assignment_group,omitempty"`
		BusinessDuration string `json:"business_duration,omitempty"`
		BusinessService  string `json:"business_service,omitempty"`
		CalendarDuration string `json:"calendar_duration,omitempty"`
		Case             string `json:"case,omitempty"`
		CaseReport       struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"case_report,omitempty"`
		Category   string `json:"category,omitempty"`
		Cause      string `json:"cause,omitempty"`
		CausedBy   string `json:"caused_by,omitempty"`
		Change     string `json:"change,omitempty"`
		ClosedAt   string `json:"closed_at,omitempty"`
		ClosedBy   string `json:"closed_by,omitempty"`
		CloseNotes string `json:"close_notes,omitempty"`
		CmdbCi     struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"cmdb_ci,omitempty"`
		Comments             string `json:"comments,omitempty"`
		CommentsAndWorkNotes string `json:"comments_and_work_notes,omitempty"`
		Company              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		Consumer string `json:"consumer,omitempty"`
		Contact  struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"contact,omitempty"`
		ContactLocalTime   string `json:"contact_local_time,omitempty"`
		ContactTimeZone    string `json:"contact_time_zone,omitempty"`
		ContactType        string `json:"contact_type,omitempty"`
		Contract           string `json:"contract,omitempty"`
		CorrelationDisplay string `json:"correlation_display,omitempty"`
		CorrelationID      string `json:"correlation_id,omitempty"`
		DeliveryPlan       string `json:"delivery_plan,omitempty"`
		DeliveryTask       string `json:"delivery_task,omitempty"`
		Description        string `json:"description,omitempty"`
		DueDate            string `json:"due_date,omitempty"`
		Entitlement        string `json:"entitlement,omitempty"`
		Escalation         string `json:"escalation,omitempty"`
		ExpectedStart      string `json:"expected_start,omitempty"`
		FirstResponseTime  string `json:"first_response_time,omitempty"`
		FollowTheSun       string `json:"follow_the_sun,omitempty"`
		FollowUp           string `json:"follow_up,omitempty"`
		GroupList          string `json:"group_list,omitempty"`
		Impact             string `json:"impact,omitempty"`
		Incident           string `json:"incident,omitempty"`
		Knowledge          string `json:"knowledge,omitempty"`
		Location           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"location,omitempty"`
		MadeSLA         string `json:"made_sla,omitempty"`
		NotesToComments string `json:"notes_to_comments,omitempty"`
		Notify          string `json:"notify,omitempty"`
		Number          string `json:"number,omitempty"`
		OpenedAt        string `json:"opened_at,omitempty"`
		OpenedBy        struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"opened_by,omitempty"`
		Order                      string `json:"order,omitempty"`
		Parent                     string `json:"parent,omitempty"`
		Partner                    string `json:"partner,omitempty"`
		PartnerContact             string `json:"partner_contact,omitempty"`
		Priority                   string `json:"priority,omitempty"`
		Problem                    string `json:"problem,omitempty"`
		Product                    string `json:"product,omitempty"`
		ReassignmentCount          string `json:"reassignment_count,omitempty"`
		ResolutionCode             string `json:"resolution_code,omitempty"`
		ResolvedAt                 string `json:"resolved_at,omitempty"`
		ResolvedBy                 string `json:"resolved_by,omitempty"`
		ShortDescription           string `json:"short_description,omitempty"`
		Skills                     string `json:"skills,omitempty"`
		SLADue                     string `json:"sla_due,omitempty"`
		SnAppCsSocialSocialProfile string `json:"sn_app_cs_social_social_profile,omitempty"`
		State                      string `json:"state,omitempty"`
		Subcategory                string `json:"subcategory,omitempty"`
		SupportManager             string `json:"support_manager,omitempty"`
		SyncDriver                 string `json:"sync_driver,omitempty"`
		SysClassName               string `json:"sys_class_name,omitempty"`
		SysCreatedBy               string `json:"sys_created_by,omitempty"`
		SysCreatedOn               string `json:"sys_created_on,omitempty"`
		SysDomain                  struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath string `json:"sys_domain_path,omitempty"`
		SysID         string `json:"sys_id,omitempty"`
		SysModCount   string `json:"sys_mod_count,omitempty"`
		SysTags       string `json:"sys_tags,omitempty"`
		SysUpdatedBy  string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn  string `json:"sys_updated_on,omitempty"`
		TaskFor       struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"task_for,omitempty"`
		TimeWorked           string `json:"time_worked,omitempty"`
		UponApproval         string `json:"upon_approval,omitempty"`
		UponReject           string `json:"upon_reject,omitempty"`
		Urgency              string `json:"urgency,omitempty"`
		UserInput            string `json:"user_input,omitempty"`
		WatchList            string `json:"watch_list,omitempty"`
		WorkEnd              string `json:"work_end,omitempty"`
		WorkNotes            string `json:"work_notes,omitempty"`
		WorkNotesList        string `json:"work_notes_list,omitempty"`
		WorkStart            string `json:"work_start,omitempty"`
		UCustomerVisible     string `json:"u_customer_visible,omitempty"`      // Custom Field
		UDoesKbNeedUpdated   string `json:"u_does_kb_need_updated,omitempty"`  // Custom Field
		UGlideDateTime       string `json:"u_glide_date_time,omitempty"`       // Custom Field
		UHasBreached         string `json:"u_has_breached,omitempty"`          // Custom Field
		UNeedsAttention      string `json:"u_needs_attention,omitempty"`       // Custom Field
		UNotificationList    string `json:"u_notification_list,omitempty"`     // Custom Field
		UPublicComment       string `json:"u_public_comment,omitempty"`        // Custom Field
		UPublicCommentsAdded string `json:"u_public_comments_added,omitempty"` // Custom Field
		UReason              string `json:"u_reason,omitempty"`                // Custom Field
		USecureComment       string `json:"u_secure_comment,omitempty"`        // Custom Field
		UToDo                string `json:"u_to_do,omitempty"`                 // Custom Field
		UVip                 string `json:"u_vip,omitempty"`                   // Custom Field
	} `json:"result,omitempty"`
}

type CmdbCiResultDetail struct {
	Result struct {
		Asset struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"asset,omitempty"`
		AssetTag        string `json:"asset_tag,omitempty"`
		Assigned        string `json:"assigned,omitempty"`
		AssignedTo      string `json:"assigned_to,omitempty"`
		AssignmentGroup string `json:"assignment_group,omitempty"`
		Attributes      string `json:"attributes,omitempty"`
		CanPrint        string `json:"can_print,omitempty"`
		Category        string `json:"category,omitempty"`
		ChangeControl   string `json:"change_control,omitempty"`
		CheckedIn       string `json:"checked_in,omitempty"`
		CheckedOut      string `json:"checked_out,omitempty"`
		Comments        string `json:"comments,omitempty"`
		Company         struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CorrelationID       string `json:"correlation_id,omitempty"`
		Cost                string `json:"cost,omitempty"`
		CostCc              string `json:"cost_cc,omitempty"`
		CostCenter          string `json:"cost_center,omitempty"`
		DeliveryDate        string `json:"delivery_date,omitempty"`
		Department          string `json:"department,omitempty"`
		DiscoverySource     string `json:"discovery_source,omitempty"`
		DNSDomain           string `json:"dns_domain,omitempty"`
		Due                 string `json:"due,omitempty"`
		DueIn               string `json:"due_in,omitempty"`
		FaultCount          string `json:"fault_count,omitempty"`
		FirstDiscovered     string `json:"first_discovered,omitempty"`
		Fqdn                string `json:"fqdn,omitempty"`
		GlAccount           string `json:"gl_account,omitempty"`
		InstallDate         string `json:"install_date,omitempty"`
		InstallStatus       string `json:"install_status,omitempty"`
		InvoiceNumber       string `json:"invoice_number,omitempty"`
		IPAddress           string `json:"ip_address,omitempty"`
		Justification       string `json:"justification,omitempty"`
		LastDiscovered      string `json:"last_discovered,omitempty"`
		LeaseID             string `json:"lease_id,omitempty"`
		Location            string `json:"location,omitempty"`
		MacAddress          string `json:"mac_address,omitempty"`
		MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
		ManagedBy           string `json:"managed_by,omitempty"`
		ManagedDomain       string `json:"managed_domain,omitempty"`
		Manufacturer        string `json:"manufacturer,omitempty"`
		ModelID             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"model_id,omitempty"`
		ModelNumber       string `json:"model_number,omitempty"`
		Monitor           string `json:"monitor,omitempty"`
		Name              string `json:"name,omitempty"`
		OperationalStatus string `json:"operational_status,omitempty"`
		OrderDate         string `json:"order_date,omitempty"`
		OwnedBy           string `json:"owned_by,omitempty"`
		PoNumber          string `json:"po_number,omitempty"`
		PurchaseDate      string `json:"purchase_date,omitempty"`
		Schedule          string `json:"schedule,omitempty"`
		SerialNumber      string `json:"serial_number,omitempty"`
		ShortDescription  string `json:"short_description,omitempty"`
		SkipSync          string `json:"skip_sync,omitempty"`
		StartDate         string `json:"start_date,omitempty"`
		Subcategory       string `json:"subcategory,omitempty"`
		SupportedBy       string `json:"supported_by,omitempty"`
		SupportGroup      string `json:"support_group,omitempty"`
		SysClassName      string `json:"sys_class_name,omitempty"`
		SysClassPath      string `json:"sys_class_path,omitempty"`
		SysCreatedBy      string `json:"sys_created_by,omitempty"`
		SysCreatedOn      string `json:"sys_created_on,omitempty"`
		SysDomain         struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath      string `json:"sys_domain_path,omitempty"`
		SysID              string `json:"sys_id,omitempty"`
		SysModCount        string `json:"sys_mod_count,omitempty"`
		SysTags            string `json:"sys_tags,omitempty"`
		SysUpdatedBy       string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn       string `json:"sys_updated_on,omitempty"`
		Unverified         string `json:"unverified,omitempty"`
		Vendor             string `json:"vendor,omitempty"`
		WarrantyExpiration string `json:"warranty_expiration,omitempty"`
		UApplicationPack   string `json:"u_application_pack,omitempty"`   // Custom Field
		UBackupServer      string `json:"u_backup_server,omitempty"`      // Custom Field
		UChangeAuthorizers string `json:"u_change_authorizers,omitempty"` // Custom Field
		UComments          string `json:"u_comments,omitempty"`           // Custom Field
		UDeviceID          string `json:"u_device_id,omitempty"`          // Custom Field
		UIPAddress         string `json:"u_ip_address,omitempty"`         // Custom Field
		UKbArticle         string `json:"u_kb_article,omitempty"`         // Custom Field
		ULegacySource      string `json:"u_legacy_source,omitempty"`      // Custom Field
		UManagementPack    string `json:"u_management_pack,omitempty"`    // Custom Field
		UParent            string `json:"u_parent,omitempty"`             // Custom Field
		UPollerGroup       string `json:"u_poller_group,omitempty"`       // Custom Field
		URackPosition      string `json:"u_rack_position,omitempty"`      // Custom Field
		URackUnit          string `json:"u_rack_unit,omitempty"`          // Custom Field
		UVip               string `json:"u_vip,omitempty"`                // Custom Field
	} `json:"result,omitempty"`
}

type CmdbCiServerResultDetail struct {
	Result struct {
		Asset struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"asset,omitempty"`
		AssetTag        string `json:"asset_tag,omitempty"`
		Assigned        string `json:"assigned,omitempty"`
		AssignedTo      string `json:"assigned_to,omitempty"`
		AssignmentGroup string `json:"assignment_group,omitempty"`
		Attributes      string `json:"attributes,omitempty"`
		CanPrint        string `json:"can_print,omitempty"`
		Category        string `json:"category,omitempty"`
		CdRom           string `json:"cd_rom,omitempty"`
		CdSpeed         string `json:"cd_speed,omitempty"`
		ChangeControl   string `json:"change_control,omitempty"`
		ChassisType     string `json:"chassis_type,omitempty"`
		CheckedIn       string `json:"checked_in,omitempty"`
		CheckedOut      string `json:"checked_out,omitempty"`
		Classification  string `json:"classification,omitempty"`
		Comments        string `json:"comments,omitempty"`
		Company         struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CorrelationID       string `json:"correlation_id,omitempty"`
		Cost                string `json:"cost,omitempty"`
		CostCc              string `json:"cost_cc,omitempty"`
		CostCenter          string `json:"cost_center,omitempty"`
		CPUCoreCount        string `json:"cpu_core_count,omitempty"`
		CPUCoreThread       string `json:"cpu_core_thread,omitempty"`
		CPUCount            string `json:"cpu_count,omitempty"`
		CPUManufacturer     string `json:"cpu_manufacturer,omitempty"`
		CPUName             string `json:"cpu_name,omitempty"`
		CPUSpeed            string `json:"cpu_speed,omitempty"`
		CPUType             string `json:"cpu_type,omitempty"`
		DefaultGateway      string `json:"default_gateway,omitempty"`
		DeliveryDate        string `json:"delivery_date,omitempty"`
		Department          string `json:"department,omitempty"`
		DiscoverySource     string `json:"discovery_source,omitempty"`
		DiskSpace           string `json:"disk_space,omitempty"`
		DNSDomain           string `json:"dns_domain,omitempty"`
		DrBackup            string `json:"dr_backup,omitempty"`
		Due                 string `json:"due,omitempty"`
		DueIn               string `json:"due_in,omitempty"`
		FaultCount          string `json:"fault_count,omitempty"`
		FirewallStatus      string `json:"firewall_status,omitempty"`
		FirstDiscovered     string `json:"first_discovered,omitempty"`
		Floppy              string `json:"floppy,omitempty"`
		FormFactor          string `json:"form_factor,omitempty"`
		Fqdn                string `json:"fqdn,omitempty"`
		GlAccount           string `json:"gl_account,omitempty"`
		HardwareStatus      string `json:"hardware_status,omitempty"`
		HardwareSubstatus   string `json:"hardware_substatus,omitempty"`
		HostName            string `json:"host_name,omitempty"`
		InstallDate         string `json:"install_date,omitempty"`
		InstallStatus       string `json:"install_status,omitempty"`
		InvoiceNumber       string `json:"invoice_number,omitempty"`
		IPAddress           string `json:"ip_address,omitempty"`
		Justification       string `json:"justification,omitempty"`
		LastDiscovered      string `json:"last_discovered,omitempty"`
		LeaseID             string `json:"lease_id,omitempty"`
		Location            string `json:"location,omitempty"`
		MacAddress          string `json:"mac_address,omitempty"`
		MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
		ManagedBy           string `json:"managed_by,omitempty"`
		ManagedDomain       string `json:"managed_domain,omitempty"`
		Manufacturer        string `json:"manufacturer,omitempty"`
		ModelID             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"model_id,omitempty"`
		ModelNumber       string `json:"model_number,omitempty"`
		Monitor           string `json:"monitor,omitempty"`
		Name              string `json:"name,omitempty"`
		ObjectID          string `json:"object_id,omitempty"`
		OperationalStatus string `json:"operational_status,omitempty"`
		OrderDate         string `json:"order_date,omitempty"`
		Os                string `json:"os,omitempty"`
		OsAddressWidth    string `json:"os_address_width,omitempty"`
		OsDomain          string `json:"os_domain,omitempty"`
		OsServicePack     string `json:"os_service_pack,omitempty"`
		OsVersion         string `json:"os_version,omitempty"`
		OwnedBy           string `json:"owned_by,omitempty"`
		PoNumber          string `json:"po_number,omitempty"`
		PurchaseDate      string `json:"purchase_date,omitempty"`
		RAM               string `json:"ram,omitempty"`
		Schedule          string `json:"schedule,omitempty"`
		SerialNumber      string `json:"serial_number,omitempty"`
		ShortDescription  string `json:"short_description,omitempty"`
		SkipSync          string `json:"skip_sync,omitempty"`
		StartDate         string `json:"start_date,omitempty"`
		Subcategory       string `json:"subcategory,omitempty"`
		SupportedBy       string `json:"supported_by,omitempty"`
		SupportGroup      string `json:"support_group,omitempty"`
		SysClassName      string `json:"sys_class_name,omitempty"`
		SysClassPath      string `json:"sys_class_path,omitempty"`
		SysCreatedBy      string `json:"sys_created_by,omitempty"`
		SysCreatedOn      string `json:"sys_created_on,omitempty"`
		SysDomain         struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath       string `json:"sys_domain_path,omitempty"`
		SysID               string `json:"sys_id,omitempty"`
		SysModCount         string `json:"sys_mod_count,omitempty"`
		SysTags             string `json:"sys_tags,omitempty"`
		SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
		Unverified          string `json:"unverified,omitempty"`
		UsedFor             string `json:"used_for,omitempty"`
		Vendor              string `json:"vendor,omitempty"`
		Virtual             string `json:"virtual,omitempty"`
		WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
		UApplicationPack    string `json:"u_application_pack,omitempty"`     // Custom Field
		UBackupRetention    string `json:"u_backup_retention,omitempty"`     // Custom Field
		UBackupSchedule     string `json:"u_backup_schedule,omitempty"`      // Custom Field
		UBackupServer       string `json:"u_backup_server,omitempty"`        // Custom Field
		UBackupSystem       string `json:"u_backup_system,omitempty"`        // Custom Field
		UBuildDate          string `json:"u_build_date,omitempty"`           // Custom Field
		UChangeAuthorizers  string `json:"u_change_authorizers,omitempty"`   // Custom Field
		UCloudManagement    string `json:"u_cloud_management,omitempty"`     // Custom Field
		UComments           string `json:"u_comments,omitempty"`             // Custom Field
		UCustomerDeviceName string `json:"u_customer_device_name,omitempty"` // Custom Field
		UDeviceID           string `json:"u_device_id,omitempty"`            // Custom Field
		UDeviceLocation     string `json:"u_device_location,omitempty"`      // Custom Field
		UDeviceStatus       string `json:"u_device_status,omitempty"`        // Custom Field
		UDeviceType         string `json:"u_device_type,omitempty"`          // Custom Field
		UExcludeFromAnsible string `json:"u_exclude_from_ansible,omitempty"` // Custom Field
		UIPAddress          string `json:"u_ip_address,omitempty"`           // Custom Field
		UKbArticle          string `json:"u_kb_article,omitempty"`           // Custom Field
		ULegacySource       string `json:"u_legacy_source,omitempty"`        // Custom Field
		ULoginPort          string `json:"u_login_port,omitempty"`           // Custom Field
		ULuksErv            string `json:"u_luks_erv,omitempty"`             // Custom Field
		ULuksPmpEntry       string `json:"u_luks_pmp_entry,omitempty"`       // Custom Field
		UManagementPack     string `json:"u_management_pack,omitempty"`      // Custom Field
		UMotherboard        string `json:"u_motherboard,omitempty"`          // Custom Field
		UOsFamily           string `json:"u_os_family,omitempty"`            // Custom Field
		UParent             string `json:"u_parent,omitempty"`               // Custom Field
		UPollerGroup        string `json:"u_poller_group,omitempty"`         // Custom Field
		URack               string `json:"u_rack,omitempty"`                 // Custom Field
		URackPosition       string `json:"u_rack_position,omitempty"`        // Custom Field
		URackUnit           string `json:"u_rack_unit,omitempty"`            // Custom Field
		URAMDescription     string `json:"u_ram_description,omitempty"`      // Custom Field
		URAMHardwareCount   string `json:"u_ram_hardware_count,omitempty"`   // Custom Field
		URoom               string `json:"u_room,omitempty"`                 // Custom Field
		USerialID           string `json:"u_serial_id,omitempty"`            // Custom Field
		UVip                string `json:"u_vip,omitempty"`                  // Custom Field
		UVirtualCPUCount    string `json:"u_virtual_cpu_count,omitempty"`    // Custom Field
		UVirtualDiskSize    string `json:"u_virtual_disk_size,omitempty"`    // Custom Field
		UVirtualRAMSize     string `json:"u_virtual_ram_size,omitempty"`     // Custom Field
	} `json:"result,omitempty"`
}

type CmdbCiIpRouterResultDetail struct {
	Result struct {
		Asset struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"asset,omitempty"`
		AssetTag          string `json:"asset_tag,omitempty"`
		Assigned          string `json:"assigned,omitempty"`
		AssignedTo        string `json:"assigned_to,omitempty"`
		AssignmentGroup   string `json:"assignment_group,omitempty"`
		Attributes        string `json:"attributes,omitempty"`
		Bandwidth         string `json:"bandwidth,omitempty"`
		CanHub            string `json:"can_hub,omitempty"`
		CanPartitionvlans string `json:"can_partitionvlans,omitempty"`
		CanPrint          string `json:"can_print,omitempty"`
		CanRoute          string `json:"can_route,omitempty"`
		CanSwitch         string `json:"can_switch,omitempty"`
		Category          string `json:"category,omitempty"`
		ChangeControl     string `json:"change_control,omitempty"`
		Channels          string `json:"channels,omitempty"`
		CheckedIn         string `json:"checked_in,omitempty"`
		CheckedOut        string `json:"checked_out,omitempty"`
		Comments          string `json:"comments,omitempty"`
		Company           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CorrelationID        string `json:"correlation_id,omitempty"`
		Cost                 string `json:"cost,omitempty"`
		CostCc               string `json:"cost_cc,omitempty"`
		CostCenter           string `json:"cost_center,omitempty"`
		CPUCount             string `json:"cpu_count,omitempty"`
		CPUManufacturer      string `json:"cpu_manufacturer,omitempty"`
		CPUSpeed             string `json:"cpu_speed,omitempty"`
		CPUType              string `json:"cpu_type,omitempty"`
		DefaultGateway       string `json:"default_gateway,omitempty"`
		DeliveryDate         string `json:"delivery_date,omitempty"`
		Department           string `json:"department,omitempty"`
		DeviceType           string `json:"device_type,omitempty"`
		DiscoveryProtoID     string `json:"discovery_proto_id,omitempty"`
		DiscoveryProtoType   string `json:"discovery_proto_type,omitempty"`
		DiscoverySource      string `json:"discovery_source,omitempty"`
		DiskSpace            string `json:"disk_space,omitempty"`
		DNSDomain            string `json:"dns_domain,omitempty"`
		Due                  string `json:"due,omitempty"`
		DueIn                string `json:"due_in,omitempty"`
		FaultCount           string `json:"fault_count,omitempty"`
		FirmwareManufacturer string `json:"firmware_manufacturer,omitempty"`
		FirmwareVersion      string `json:"firmware_version,omitempty"`
		FirstDiscovered      string `json:"first_discovered,omitempty"`
		Fqdn                 string `json:"fqdn,omitempty"`
		GlAccount            string `json:"gl_account,omitempty"`
		HardwareStatus       string `json:"hardware_status,omitempty"`
		HardwareSubstatus    string `json:"hardware_substatus,omitempty"`
		InstallDate          string `json:"install_date,omitempty"`
		InstallStatus        string `json:"install_status,omitempty"`
		InvoiceNumber        string `json:"invoice_number,omitempty"`
		IPAddress            string `json:"ip_address,omitempty"`
		Justification        string `json:"justification,omitempty"`
		LastDiscovered       string `json:"last_discovered,omitempty"`
		LeaseID              string `json:"lease_id,omitempty"`
		Location             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"location,omitempty"`
		MacAddress          string `json:"mac_address,omitempty"`
		MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
		ManagedBy           string `json:"managed_by,omitempty"`
		ManagedDomain       string `json:"managed_domain,omitempty"`
		Manufacturer        string `json:"manufacturer,omitempty"`
		ModelID             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"model_id,omitempty"`
		ModelNumber            string `json:"model_number,omitempty"`
		Monitor                string `json:"monitor,omitempty"`
		Name                   string `json:"name,omitempty"`
		OperationalStatus      string `json:"operational_status,omitempty"`
		OrderDate              string `json:"order_date,omitempty"`
		OwnedBy                string `json:"owned_by,omitempty"`
		PhysicalInterfaceCount string `json:"physical_interface_count,omitempty"`
		PoNumber               string `json:"po_number,omitempty"`
		Ports                  string `json:"ports,omitempty"`
		PurchaseDate           string `json:"purchase_date,omitempty"`
		RAM                    string `json:"ram,omitempty"`
		Range                  string `json:"range,omitempty"`
		Schedule               string `json:"schedule,omitempty"`
		SerialNumber           string `json:"serial_number,omitempty"`
		ShortDescription       string `json:"short_description,omitempty"`
		SkipSync               string `json:"skip_sync,omitempty"`
		SnmpSysLocation        string `json:"snmp_sys_location,omitempty"`
		StartDate              string `json:"start_date,omitempty"`
		Subcategory            string `json:"subcategory,omitempty"`
		SupportedBy            string `json:"supported_by,omitempty"`
		SupportGroup           string `json:"support_group,omitempty"`
		SysClassName           string `json:"sys_class_name,omitempty"`
		SysClassPath           string `json:"sys_class_path,omitempty"`
		SysCreatedBy           string `json:"sys_created_by,omitempty"`
		SysCreatedOn           string `json:"sys_created_on,omitempty"`
		SysDomain              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath       string `json:"sys_domain_path,omitempty"`
		SysID               string `json:"sys_id,omitempty"`
		SysModCount         string `json:"sys_mod_count,omitempty"`
		SysTags             string `json:"sys_tags,omitempty"`
		SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
		Unverified          string `json:"unverified,omitempty"`
		Vendor              string `json:"vendor,omitempty"`
		WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
		UApplicationPack    string `json:"u_application_pack,omitempty"`     // Custom Field
		UBackupServer       string `json:"u_backup_server,omitempty"`        // Custom Field
		UBuildDate          string `json:"u_build_date,omitempty"`           // Custom Field
		UChangeAuthorizers  string `json:"u_change_authorizers,omitempty"`   // Custom Field
		UComments           string `json:"u_comments,omitempty"`             // Custom Field
		UDeviceID           string `json:"u_device_id,omitempty"`            // Custom Field
		UDeviceLocation     string `json:"u_device_location,omitempty"`      // Custom Field
		UExcludeFromAnsible string `json:"u_exclude_from_ansible,omitempty"` // Custom Field
		UHostName           string `json:"u_host_name,omitempty"`            // Custom Field
		UIPAddress          string `json:"u_ip_address,omitempty"`           // Custom Field
		UKbArticle          string `json:"u_kb_article,omitempty"`           // Custom Field
		ULegacySource       string `json:"u_legacy_source,omitempty"`        // Custom Field
		ULoginPort          string `json:"u_login_port,omitempty"`           // Custom Field
		UManagementLevel    string `json:"u_management_level,omitempty"`     // Custom Field
		UManagementPack     string `json:"u_management_pack,omitempty"`      // Custom Field
		UParent             string `json:"u_parent,omitempty"`               // Custom Field
		UPollerGroup        string `json:"u_poller_group,omitempty"`         // Custom Field
		URack               string `json:"u_rack,omitempty"`                 // Custom Field
		URackPosition       string `json:"u_rack_position,omitempty"`        // Custom Field
		URackUnit           string `json:"u_rack_unit,omitempty"`            // Custom Field
		URAMDescription     string `json:"u_ram_description,omitempty"`      // Custom Field
		URAMHardwareCount   string `json:"u_ram_hardware_count,omitempty"`   // Custom Field
		URoom               string `json:"u_room,omitempty"`                 // Custom Field
		UVip                string `json:"u_vip,omitempty"`                  // Custom Field
		UVirtualDiskSize    string `json:"u_virtual_disk_size,omitempty"`    // Custom Field
		UVirtualRAMSize     string `json:"u_virtual_ram_size,omitempty"`     // Custom Field
	} `json:"result,omitempty"`
}

type CmdbCiIpSwitchResultDetail struct {
	Result struct {
		Asset struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"asset,omitempty"`
		AssetTag          string `json:"asset_tag,omitempty"`
		Assigned          string `json:"assigned,omitempty"`
		AssignedTo        string `json:"assigned_to,omitempty"`
		AssignmentGroup   string `json:"assignment_group,omitempty"`
		Attributes        string `json:"attributes,omitempty"`
		Bandwidth         string `json:"bandwidth,omitempty"`
		CanHub            string `json:"can_hub,omitempty"`
		CanPartitionvlans string `json:"can_partitionvlans,omitempty"`
		CanPrint          string `json:"can_print,omitempty"`
		CanRoute          string `json:"can_route,omitempty"`
		CanSwitch         string `json:"can_switch,omitempty"`
		Category          string `json:"category,omitempty"`
		ChangeControl     string `json:"change_control,omitempty"`
		Channels          string `json:"channels,omitempty"`
		CheckedIn         string `json:"checked_in,omitempty"`
		CheckedOut        string `json:"checked_out,omitempty"`
		Comments          string `json:"comments,omitempty"`
		Company           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CorrelationID        string `json:"correlation_id,omitempty"`
		Cost                 string `json:"cost,omitempty"`
		CostCc               string `json:"cost_cc,omitempty"`
		CostCenter           string `json:"cost_center,omitempty"`
		CPUCount             string `json:"cpu_count,omitempty"`
		CPUManufacturer      string `json:"cpu_manufacturer,omitempty"`
		CPUSpeed             string `json:"cpu_speed,omitempty"`
		CPUType              string `json:"cpu_type,omitempty"`
		DefaultGateway       string `json:"default_gateway,omitempty"`
		DeliveryDate         string `json:"delivery_date,omitempty"`
		Department           string `json:"department,omitempty"`
		DeviceType           string `json:"device_type,omitempty"`
		DiscoveryProtoID     string `json:"discovery_proto_id,omitempty"`
		DiscoveryProtoType   string `json:"discovery_proto_type,omitempty"`
		DiscoverySource      string `json:"discovery_source,omitempty"`
		DiskSpace            string `json:"disk_space,omitempty"`
		DNSDomain            string `json:"dns_domain,omitempty"`
		Due                  string `json:"due,omitempty"`
		DueIn                string `json:"due_in,omitempty"`
		FaultCount           string `json:"fault_count,omitempty"`
		FirmwareManufacturer string `json:"firmware_manufacturer,omitempty"`
		FirmwareVersion      string `json:"firmware_version,omitempty"`
		FirstDiscovered      string `json:"first_discovered,omitempty"`
		Fqdn                 string `json:"fqdn,omitempty"`
		GlAccount            string `json:"gl_account,omitempty"`
		HardwareStatus       string `json:"hardware_status,omitempty"`
		HardwareSubstatus    string `json:"hardware_substatus,omitempty"`
		InstallDate          string `json:"install_date,omitempty"`
		InstallStatus        string `json:"install_status,omitempty"`
		InvoiceNumber        string `json:"invoice_number,omitempty"`
		IPAddress            string `json:"ip_address,omitempty"`
		Justification        string `json:"justification,omitempty"`
		LastDiscovered       string `json:"last_discovered,omitempty"`
		LeaseID              string `json:"lease_id,omitempty"`
		Location             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"location,omitempty"`
		MacAddress          string `json:"mac_address,omitempty"`
		MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
		ManagedBy           string `json:"managed_by,omitempty"`
		ManagedDomain       string `json:"managed_domain,omitempty"`
		Manufacturer        string `json:"manufacturer,omitempty"`
		ModelID             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"model_id,omitempty"`
		ModelNumber            string `json:"model_number,omitempty"`
		Monitor                string `json:"monitor,omitempty"`
		Name                   string `json:"name,omitempty"`
		OperationalStatus      string `json:"operational_status,omitempty"`
		OrderDate              string `json:"order_date,omitempty"`
		OwnedBy                string `json:"owned_by,omitempty"`
		PhysicalInterfaceCount string `json:"physical_interface_count,omitempty"`
		PoNumber               string `json:"po_number,omitempty"`
		Ports                  string `json:"ports,omitempty"`
		PurchaseDate           string `json:"purchase_date,omitempty"`
		RAM                    string `json:"ram,omitempty"`
		Range                  string `json:"range,omitempty"`
		Schedule               string `json:"schedule,omitempty"`
		SerialNumber           string `json:"serial_number,omitempty"`
		ShortDescription       string `json:"short_description,omitempty"`
		SkipSync               string `json:"skip_sync,omitempty"`
		SnmpSysLocation        string `json:"snmp_sys_location,omitempty"`
		Stack                  string `json:"stack,omitempty"`
		StackMode              string `json:"stack_mode,omitempty"`
		StartDate              string `json:"start_date,omitempty"`
		Subcategory            string `json:"subcategory,omitempty"`
		SupportedBy            string `json:"supported_by,omitempty"`
		SupportGroup           string `json:"support_group,omitempty"`
		SysClassName           string `json:"sys_class_name,omitempty"`
		SysClassPath           string `json:"sys_class_path,omitempty"`
		SysCreatedBy           string `json:"sys_created_by,omitempty"`
		SysCreatedOn           string `json:"sys_created_on,omitempty"`
		SysDomain              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath       string `json:"sys_domain_path,omitempty"`
		SysID               string `json:"sys_id,omitempty"`
		SysModCount         string `json:"sys_mod_count,omitempty"`
		SysTags             string `json:"sys_tags,omitempty"`
		SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
		Unverified          string `json:"unverified,omitempty"`
		Vendor              string `json:"vendor,omitempty"`
		WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
		UApplicationPack    string `json:"u_application_pack,omitempty"`     // Custom Field
		UBackupServer       string `json:"u_backup_server,omitempty"`        // Custom Field
		UBuildDate          string `json:"u_build_date,omitempty"`           // Custom Field
		UChangeAuthorizers  string `json:"u_change_authorizers,omitempty"`   // Custom Field
		UComments           string `json:"u_comments,omitempty"`             // Custom Field
		UDeviceID           string `json:"u_device_id,omitempty"`            // Custom Field
		UDeviceLocation     string `json:"u_device_location,omitempty"`      // Custom Field
		UExcludeFromAnsible string `json:"u_exclude_from_ansible,omitempty"` // Custom Field
		UHostName           string `json:"u_host_name,omitempty"`            // Custom Field
		UIPAddress          string `json:"u_ip_address,omitempty"`           // Custom Field
		UKbArticle          string `json:"u_kb_article,omitempty"`           // Custom Field
		ULegacySource       string `json:"u_legacy_source,omitempty"`        // Custom Field
		ULoginPort          string `json:"u_login_port,omitempty"`           // Custom Field
		UManagementLevel    string `json:"u_management_level,omitempty"`     // Custom Field
		UManagementPack     string `json:"u_management_pack,omitempty"`      // Custom Field
		UParent             string `json:"u_parent,omitempty"`               // Custom Field
		UPollerGroup        string `json:"u_poller_group,omitempty"`         // Custom Field
		URack               struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"u_rack,omitempty"` // Custom Field
		URackPosition     string `json:"u_rack_position,omitempty"`      // Custom Field
		URackUnit         string `json:"u_rack_unit,omitempty"`          // Custom Field
		URAMDescription   string `json:"u_ram_description,omitempty"`    // Custom Field
		URAMHardwareCount string `json:"u_ram_hardware_count,omitempty"` // Custom Field
		URoom             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"u_room,omitempty"` // Custom Field
		UVip             string `json:"u_vip,omitempty"`               // Custom Field
		UVirtualDiskSize string `json:"u_virtual_disk_size,omitempty"` // Custom Field
		UVirtualRAMSize  string `json:"u_virtual_ram_size,omitempty"`  // Custom Field
	} `json:"result,omitempty"`
}

type CmdbCiFirewallNetworkResultDetail struct {
	Result struct {
		Asset             string `json:"asset,omitempty"`
		AssetTag          string `json:"asset_tag,omitempty"`
		Assigned          string `json:"assigned,omitempty"`
		AssignedTo        string `json:"assigned_to,omitempty"`
		AssignmentGroup   string `json:"assignment_group,omitempty"`
		Attributes        string `json:"attributes,omitempty"`
		Bandwidth         string `json:"bandwidth,omitempty"`
		CanHub            string `json:"can_hub,omitempty"`
		CanPartitionvlans string `json:"can_partitionvlans,omitempty"`
		CanPrint          string `json:"can_print,omitempty"`
		CanRoute          string `json:"can_route,omitempty"`
		CanSwitch         string `json:"can_switch,omitempty"`
		Category          string `json:"category,omitempty"`
		ChangeControl     string `json:"change_control,omitempty"`
		Channels          string `json:"channels,omitempty"`
		CheckedIn         string `json:"checked_in,omitempty"`
		CheckedOut        string `json:"checked_out,omitempty"`
		Comments          string `json:"comments,omitempty"`
		Company           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CorrelationID        string `json:"correlation_id,omitempty"`
		Cost                 string `json:"cost,omitempty"`
		CostCc               string `json:"cost_cc,omitempty"`
		CostCenter           string `json:"cost_center,omitempty"`
		CPUCount             string `json:"cpu_count,omitempty"`
		CPUManufacturer      string `json:"cpu_manufacturer,omitempty"`
		CPUSpeed             string `json:"cpu_speed,omitempty"`
		CPUType              string `json:"cpu_type,omitempty"`
		DefaultGateway       string `json:"default_gateway,omitempty"`
		DeliveryDate         string `json:"delivery_date,omitempty"`
		Department           string `json:"department,omitempty"`
		DeviceType           string `json:"device_type,omitempty"`
		DiscoveryProtoID     string `json:"discovery_proto_id,omitempty"`
		DiscoveryProtoType   string `json:"discovery_proto_type,omitempty"`
		DiscoverySource      string `json:"discovery_source,omitempty"`
		DiskSpace            string `json:"disk_space,omitempty"`
		DNSDomain            string `json:"dns_domain,omitempty"`
		Due                  string `json:"due,omitempty"`
		DueIn                string `json:"due_in,omitempty"`
		FaultCount           string `json:"fault_count,omitempty"`
		FirmwareManufacturer string `json:"firmware_manufacturer,omitempty"`
		FirmwareVersion      string `json:"firmware_version,omitempty"`
		FirstDiscovered      string `json:"first_discovered,omitempty"`
		Fqdn                 string `json:"fqdn,omitempty"`
		GlAccount            string `json:"gl_account,omitempty"`
		HardwareStatus       string `json:"hardware_status,omitempty"`
		HardwareSubstatus    string `json:"hardware_substatus,omitempty"`
		InstallDate          string `json:"install_date,omitempty"`
		InstallStatus        string `json:"install_status,omitempty"`
		InvoiceNumber        string `json:"invoice_number,omitempty"`
		IPAddress            string `json:"ip_address,omitempty"`
		Justification        string `json:"justification,omitempty"`
		LastDiscovered       string `json:"last_discovered,omitempty"`
		LeaseID              string `json:"lease_id,omitempty"`
		Location             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"location,omitempty"`
		MacAddress             string `json:"mac_address,omitempty"`
		MaintenanceSchedule    string `json:"maintenance_schedule,omitempty"`
		ManagedBy              string `json:"managed_by,omitempty"`
		ManagedDomain          string `json:"managed_domain,omitempty"`
		Manufacturer           string `json:"manufacturer,omitempty"`
		ModelID                string `json:"model_id,omitempty"`
		ModelNumber            string `json:"model_number,omitempty"`
		Monitor                string `json:"monitor,omitempty"`
		Name                   string `json:"name,omitempty"`
		OperationalStatus      string `json:"operational_status,omitempty"`
		OrderDate              string `json:"order_date,omitempty"`
		OwnedBy                string `json:"owned_by,omitempty"`
		PhysicalInterfaceCount string `json:"physical_interface_count,omitempty"`
		PoNumber               string `json:"po_number,omitempty"`
		Ports                  string `json:"ports,omitempty"`
		PurchaseDate           string `json:"purchase_date,omitempty"`
		RAM                    string `json:"ram,omitempty"`
		Range                  string `json:"range,omitempty"`
		Schedule               string `json:"schedule,omitempty"`
		SerialNumber           string `json:"serial_number,omitempty"`
		ShortDescription       string `json:"short_description,omitempty"`
		SkipSync               string `json:"skip_sync,omitempty"`
		SnmpSysLocation        string `json:"snmp_sys_location,omitempty"`
		StartDate              string `json:"start_date,omitempty"`
		Subcategory            string `json:"subcategory,omitempty"`
		SupportedBy            string `json:"supported_by,omitempty"`
		SupportGroup           string `json:"support_group,omitempty"`
		SysClassName           string `json:"sys_class_name,omitempty"`
		SysClassPath           string `json:"sys_class_path,omitempty"`
		SysCreatedBy           string `json:"sys_created_by,omitempty"`
		SysCreatedOn           string `json:"sys_created_on,omitempty"`
		SysDomain              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath       string `json:"sys_domain_path,omitempty"`
		SysID               string `json:"sys_id,omitempty"`
		SysModCount         string `json:"sys_mod_count,omitempty"`
		SysTags             string `json:"sys_tags,omitempty"`
		SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
		Unverified          string `json:"unverified,omitempty"`
		Vendor              string `json:"vendor,omitempty"`
		WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
		UApplicationPack    string `json:"u_application_pack,omitempty"`     // Custom Field
		UBackupServer       string `json:"u_backup_server,omitempty"`        // Custom Field
		UBuildDate          string `json:"u_build_date,omitempty"`           // Custom Field
		UChangeAuthorizers  string `json:"u_change_authorizers,omitempty"`   // Custom Field
		UComments           string `json:"u_comments,omitempty"`             // Custom Field
		UCustomerDeviceName string `json:"u_customer_device_name,omitempty"` // Custom Field
		UDeviceID           string `json:"u_device_id,omitempty"`            // Custom Field
		UDeviceLocation     string `json:"u_device_location,omitempty"`      // Custom Field
		UExcludeFromAnsible string `json:"u_exclude_from_ansible,omitempty"` // Custom Field
		UHostName           string `json:"u_host_name,omitempty"`            // Custom Field
		UIPAddress          string `json:"u_ip_address,omitempty"`           // Custom Field
		UKbArticle          string `json:"u_kb_article,omitempty"`           // Custom Field
		ULegacySource       string `json:"u_legacy_source,omitempty"`        // Custom Field
		ULoginPort          string `json:"u_login_port,omitempty"`           // Custom Field
		UManagementLevel    string `json:"u_management_level,omitempty"`     // Custom Field
		UManagementPack     string `json:"u_management_pack,omitempty"`      // Custom Field
		UParent             string `json:"u_parent,omitempty"`               // Custom Field
		UPollerGroup        string `json:"u_poller_group,omitempty"`         // Custom Field
		URack               string `json:"u_rack,omitempty"`                 // Custom Field
		URackPosition       string `json:"u_rack_position,omitempty"`        // Custom Field
		URackUnit           string `json:"u_rack_unit,omitempty"`            // Custom Field
		URAMDescription     string `json:"u_ram_description,omitempty"`      // Custom Field
		URAMHardwareCount   string `json:"u_ram_hardware_count,omitempty"`   // Custom Field
		URoom               struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"u_room,omitempty"` // Custom Field
		UVip             string `json:"u_vip,omitempty"`               // Custom Field
		UVirtualDiskSize string `json:"u_virtual_disk_size,omitempty"` // Custom Field
		UVirtualRAMSize  string `json:"u_virtual_ram_size,omitempty"`  // Custom Field
	} `json:"result,omitempty"`
}

type CmdbCiNetgearResultDetail struct {
	Result struct {
		Asset struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"asset,omitempty"`
		AssetTag          string `json:"asset_tag,omitempty"`
		Assigned          string `json:"assigned,omitempty"`
		AssignedTo        string `json:"assigned_to,omitempty"`
		AssignmentGroup   string `json:"assignment_group,omitempty"`
		Attributes        string `json:"attributes,omitempty"`
		Bandwidth         string `json:"bandwidth,omitempty"`
		CanHub            string `json:"can_hub,omitempty"`
		CanPartitionvlans string `json:"can_partitionvlans,omitempty"`
		CanPrint          string `json:"can_print,omitempty"`
		CanRoute          string `json:"can_route,omitempty"`
		CanSwitch         string `json:"can_switch,omitempty"`
		Category          string `json:"category,omitempty"`
		ChangeControl     string `json:"change_control,omitempty"`
		Channels          string `json:"channels,omitempty"`
		CheckedIn         string `json:"checked_in,omitempty"`
		CheckedOut        string `json:"checked_out,omitempty"`
		Comments          string `json:"comments,omitempty"`
		Company           struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CorrelationID        string `json:"correlation_id,omitempty"`
		Cost                 string `json:"cost,omitempty"`
		CostCc               string `json:"cost_cc,omitempty"`
		CostCenter           string `json:"cost_center,omitempty"`
		CPUCount             string `json:"cpu_count,omitempty"`
		CPUManufacturer      string `json:"cpu_manufacturer,omitempty"`
		CPUSpeed             string `json:"cpu_speed,omitempty"`
		CPUType              string `json:"cpu_type,omitempty"`
		DefaultGateway       string `json:"default_gateway,omitempty"`
		DeliveryDate         string `json:"delivery_date,omitempty"`
		Department           string `json:"department,omitempty"`
		DeviceType           string `json:"device_type,omitempty"`
		DiscoveryProtoID     string `json:"discovery_proto_id,omitempty"`
		DiscoveryProtoType   string `json:"discovery_proto_type,omitempty"`
		DiscoverySource      string `json:"discovery_source,omitempty"`
		DiskSpace            string `json:"disk_space,omitempty"`
		DNSDomain            string `json:"dns_domain,omitempty"`
		Due                  string `json:"due,omitempty"`
		DueIn                string `json:"due_in,omitempty"`
		FaultCount           string `json:"fault_count,omitempty"`
		FirmwareManufacturer string `json:"firmware_manufacturer,omitempty"`
		FirmwareVersion      string `json:"firmware_version,omitempty"`
		FirstDiscovered      string `json:"first_discovered,omitempty"`
		Fqdn                 string `json:"fqdn,omitempty"`
		GlAccount            string `json:"gl_account,omitempty"`
		HardwareStatus       string `json:"hardware_status,omitempty"`
		HardwareSubstatus    string `json:"hardware_substatus,omitempty"`
		InstallDate          string `json:"install_date,omitempty"`
		InstallStatus        string `json:"install_status,omitempty"`
		InvoiceNumber        string `json:"invoice_number,omitempty"`
		IPAddress            string `json:"ip_address,omitempty"`
		Justification        string `json:"justification,omitempty"`
		LastDiscovered       string `json:"last_discovered,omitempty"`
		LeaseID              string `json:"lease_id,omitempty"`
		Location             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"location,omitempty"`
		MacAddress          string `json:"mac_address,omitempty"`
		MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
		ManagedBy           string `json:"managed_by,omitempty"`
		ManagedDomain       string `json:"managed_domain,omitempty"`
		Manufacturer        string `json:"manufacturer,omitempty"`
		ModelID             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"model_id,omitempty"`
		ModelNumber            string `json:"model_number,omitempty"`
		Monitor                string `json:"monitor,omitempty"`
		Name                   string `json:"name,omitempty"`
		OperationalStatus      string `json:"operational_status,omitempty"`
		OrderDate              string `json:"order_date,omitempty"`
		OwnedBy                string `json:"owned_by,omitempty"`
		PhysicalInterfaceCount string `json:"physical_interface_count,omitempty"`
		PoNumber               string `json:"po_number,omitempty"`
		Ports                  string `json:"ports,omitempty"`
		PurchaseDate           string `json:"purchase_date,omitempty"`
		RAM                    string `json:"ram,omitempty"`
		Range                  string `json:"range,omitempty"`
		Schedule               string `json:"schedule,omitempty"`
		SerialNumber           string `json:"serial_number,omitempty"`
		ShortDescription       string `json:"short_description,omitempty"`
		SkipSync               string `json:"skip_sync,omitempty"`
		SnmpSysLocation        string `json:"snmp_sys_location,omitempty"`
		StartDate              string `json:"start_date,omitempty"`
		Subcategory            string `json:"subcategory,omitempty"`
		SupportedBy            string `json:"supported_by,omitempty"`
		SupportGroup           string `json:"support_group,omitempty"`
		SysClassName           string `json:"sys_class_name,omitempty"`
		SysClassPath           string `json:"sys_class_path,omitempty"`
		SysCreatedBy           string `json:"sys_created_by,omitempty"`
		SysCreatedOn           string `json:"sys_created_on,omitempty"`
		SysDomain              struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath       string `json:"sys_domain_path,omitempty"`
		SysID               string `json:"sys_id,omitempty"`
		SysModCount         string `json:"sys_mod_count,omitempty"`
		SysTags             string `json:"sys_tags,omitempty"`
		SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
		Unverified          string `json:"unverified,omitempty"`
		Vendor              string `json:"vendor,omitempty"`
		WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
		UApplicationPack    string `json:"u_application_pack,omitempty"`     // Custom Field
		UBackupServer       string `json:"u_backup_server,omitempty"`        // Custom Field
		UBuildDate          string `json:"u_build_date,omitempty"`           // Custom Field
		UChangeAuthorizers  string `json:"u_change_authorizers,omitempty"`   // Custom Field
		UComments           string `json:"u_comments,omitempty"`             // Custom Field
		UDeviceID           string `json:"u_device_id,omitempty"`            // Custom Field
		UDeviceLocation     string `json:"u_device_location,omitempty"`      // Custom Field
		UExcludeFromAnsible string `json:"u_exclude_from_ansible,omitempty"` // Custom Field
		UHostName           string `json:"u_host_name,omitempty"`            // Custom Field
		UIPAddress          string `json:"u_ip_address,omitempty"`           // Custom Field
		UKbArticle          string `json:"u_kb_article,omitempty"`           // Custom Field
		ULegacySource       string `json:"u_legacy_source,omitempty"`        // Custom Field
		ULoginPort          string `json:"u_login_port,omitempty"`           // Custom Field
		UManagementLevel    string `json:"u_management_level,omitempty"`     // Custom Field
		UManagementPack     string `json:"u_management_pack,omitempty"`      // Custom Field
		UParent             string `json:"u_parent,omitempty"`               // Custom Field
		UPollerGroup        string `json:"u_poller_group,omitempty"`         // Custom Field
		URack               string `json:"u_rack,omitempty"`                 // Custom Field
		URackPosition       string `json:"u_rack_position,omitempty"`        // Custom Field
		URackUnit           string `json:"u_rack_unit,omitempty"`            // Custom Field
		URAMDescription     string `json:"u_ram_description,omitempty"`      // Custom Field
		URAMHardwareCount   string `json:"u_ram_hardware_count,omitempty"`   // Custom Field
		URoom               struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"u_room,omitempty"` // Custom Field
		UVip             string `json:"u_vip,omitempty"`               // Custom Field
		UVirtualDiskSize string `json:"u_virtual_disk_size,omitempty"` // Custom Field
		UVirtualRAMSize  string `json:"u_virtual_ram_size,omitempty"`  // Custom Field
	} `json:"result,omitempty"`
}

type CoreCompanyResultDetail struct {
	Result struct {
		Active           string `json:"active,omitempty"`
		AppleIcon        string `json:"apple_icon,omitempty"`
		BannerImage      string `json:"banner_image,omitempty"`
		BannerImageLight string `json:"banner_image_light,omitempty"`
		BannerText       string `json:"banner_text,omitempty"`
		City             string `json:"city,omitempty"`
		Contact          string `json:"contact,omitempty"`
		Country          string `json:"country,omitempty"`
		Customer         string `json:"customer,omitempty"`
		Discount         string `json:"discount,omitempty"`
		FaxPhone         string `json:"fax_phone,omitempty"`
		FiscalYear       string `json:"fiscal_year,omitempty"`
		Latitude         string `json:"latitude,omitempty"`
		LatLongError     string `json:"lat_long_error,omitempty"`
		Longitude        string `json:"longitude,omitempty"`
		Manufacturer     string `json:"manufacturer,omitempty"`
		MarketCap        string `json:"market_cap,omitempty"`
		Name             string `json:"name,omitempty"`
		Notes            string `json:"notes,omitempty"`
		NumEmployees     string `json:"num_employees,omitempty"`
		Parent           string `json:"parent,omitempty"`
		Phone            string `json:"phone,omitempty"`
		Primary          string `json:"primary,omitempty"`
		Profits          string `json:"profits,omitempty"`
		PubliclyTraded   string `json:"publicly_traded,omitempty"`
		RankTier         string `json:"rank_tier,omitempty"`
		RevenuePerYear   string `json:"revenue_per_year,omitempty"`
		SsoSource        string `json:"sso_source,omitempty"`
		State            string `json:"state,omitempty"`
		StockPrice       string `json:"stock_price,omitempty"`
		StockSymbol      string `json:"stock_symbol,omitempty"`
		Street           string `json:"street,omitempty"`
		SysClassName     string `json:"sys_class_name,omitempty"`
		SysCreatedBy     string `json:"sys_created_by,omitempty"`
		SysCreatedOn     string `json:"sys_created_on,omitempty"`
		SysDomain        struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath        string `json:"sys_domain_path,omitempty"`
		SysID                string `json:"sys_id,omitempty"`
		SysModCount          string `json:"sys_mod_count,omitempty"`
		SysTags              string `json:"sys_tags,omitempty"`
		SysUpdatedBy         string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn         string `json:"sys_updated_on,omitempty"`
		Theme                string `json:"theme,omitempty"`
		Vendor               string `json:"vendor,omitempty"`
		VendorManager        string `json:"vendor_manager,omitempty"`
		VendorType           string `json:"vendor_type,omitempty"`
		Website              string `json:"website,omitempty"`
		Zip                  string `json:"zip,omitempty"`
		UAccountManager      string `json:"u_account_manager,omitempty"`      // Custom Field
		UAccountStatus       string `json:"u_account_status,omitempty"`       // Custom Field
		UAutoClose           string `json:"u_auto_close,omitempty"`           // Custom Field
		UChangeAuthorizers   string `json:"u_change_authorizers,omitempty"`   // Custom Field
		UClientNumber        string `json:"u_client_number,omitempty"`        // Custom Field
		UContactEmail        string `json:"u_contact_email,omitempty"`        // Custom Field
		UCustomSop           string `json:"u_custom_sop,omitempty"`           // Custom Field
		UIndustry            string `json:"u_industry,omitempty"`             // Custom Field
		ULegacySource        string `json:"u_legacy_source,omitempty"`        // Custom Field
		UMaintenanceSchedule string `json:"u_maintenance_schedule,omitempty"` // Custom Field
		UNetworkAlias        string `json:"u_network_alias,omitempty"`        // Custom Field
		UNotificationList    string `json:"u_notification_list,omitempty"`    // Custom Field
		UProduct             string `json:"u_product,omitempty"`              // Custom Field
		USalesforceID        string `json:"u_salesforce_id,omitempty"`        // Custom Field
		USalesPerson         string `json:"u_sales_person,omitempty"`         // Custom Field
		UTechnicalContacts   string `json:"u_technical_contacts,omitempty"`   // Custom Field
		UUberClientID        string `json:"u_uber_client_id,omitempty"`       // Custom Field
		UVip                 string `json:"u_vip,omitempty"`                  // Custom Field
		UWatchList           string `json:"u_watch_list,omitempty"`           // Custom Field
	} `json:"result,omitempty"`
}

type SysUserResultDetail struct {
	Result struct {
		Active              string `json:"active,omitempty"`
		Avatar              string `json:"avatar,omitempty"`
		Building            string `json:"building,omitempty"`
		CalendarIntegration string `json:"calendar_integration,omitempty"`
		City                string `json:"city,omitempty"`
		Company             struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"company,omitempty"`
		CostCenter              string `json:"cost_center,omitempty"`
		Country                 string `json:"country,omitempty"`
		DateFormat              string `json:"date_format,omitempty"`
		DefaultPerspective      string `json:"default_perspective,omitempty"`
		Department              string `json:"department,omitempty"`
		Email                   string `json:"email,omitempty"`
		EmployeeNumber          string `json:"employee_number,omitempty"`
		EnableMultifactorAuthn  string `json:"enable_multifactor_authn,omitempty"`
		FailedAttempts          string `json:"failed_attempts,omitempty"`
		FirstName               string `json:"first_name,omitempty"`
		Gender                  string `json:"gender,omitempty"`
		HomePhone               string `json:"home_phone,omitempty"`
		InternalIntegrationUser string `json:"internal_integration_user,omitempty"`
		Introduction            string `json:"introduction,omitempty"`
		LastLogin               string `json:"last_login,omitempty"`
		LastLoginTime           string `json:"last_login_time,omitempty"`
		LastName                string `json:"last_name,omitempty"`
		LdapServer              string `json:"ldap_server,omitempty"`
		Location                string `json:"location,omitempty"`
		LockedOut               string `json:"locked_out,omitempty"`
		ManagedDomain           string `json:"managed_domain,omitempty"`
		Manager                 string `json:"manager,omitempty"`
		MiddleName              string `json:"middle_name,omitempty"`
		MobilePhone             string `json:"mobile_phone,omitempty"`
		Name                    string `json:"name,omitempty"`
		Notification            string `json:"notification,omitempty"`
		PasswordNeedsReset      string `json:"password_needs_reset,omitempty"`
		Phone                   string `json:"phone,omitempty"`
		Photo                   string `json:"photo,omitempty"`
		PreferredLanguage       string `json:"preferred_language,omitempty"`
		Roles                   string `json:"roles,omitempty"`
		Schedule                string `json:"schedule,omitempty"`
		Source                  string `json:"source,omitempty"`
		SsoSource               string `json:"sso_source,omitempty"`
		State                   string `json:"state,omitempty"`
		Street                  string `json:"street,omitempty"`
		SysClassName            string `json:"sys_class_name,omitempty"`
		SysCreatedBy            string `json:"sys_created_by,omitempty"`
		SysCreatedOn            string `json:"sys_created_on,omitempty"`
		SysDomain               struct {
			Link  string `json:"link,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"sys_domain,omitempty"`
		SysDomainPath        string `json:"sys_domain_path,omitempty"`
		SysID                string `json:"sys_id,omitempty"`
		SysModCount          string `json:"sys_mod_count,omitempty"`
		SysTags              string `json:"sys_tags,omitempty"`
		SysUpdatedBy         string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn         string `json:"sys_updated_on,omitempty"`
		TimeFormat           string `json:"time_format,omitempty"`
		TimeSheetPolicy      string `json:"time_sheet_policy,omitempty"`
		TimeZone             string `json:"time_zone,omitempty"`
		Title                string `json:"title,omitempty"`
		UserName             string `json:"user_name,omitempty"`
		UserPassword         string `json:"user_password,omitempty"`
		Vip                  string `json:"vip,omitempty"`
		WebServiceAccessOnly string `json:"web_service_access_only,omitempty"`
		Zip                  string `json:"zip,omitempty"`
		UAccess              string `json:"u_access,omitempty"`                // Custom Field
		UActiveCompanyFilter string `json:"u_active_company_filter,omitempty"` // Custom Field
		ULastActive          string `json:"u_last_active,omitempty"`           // Custom Field
		ULegacySource        string `json:"u_legacy_source,omitempty"`         // Custom Field
		UPassphrase          string `json:"u_passphrase,omitempty"`            // Custom Field
		UUberClientID        string `json:"u_uber_client_id,omitempty"`        // Custom Field
		UUberID              string `json:"u_uber_id,omitempty"`               // Custom Field
	} `json:"result,omitempty"`
}

type UApplicationPackResultDetail struct {
	Result struct {
		SysCreatedBy     string `json:"sys_created_by,omitempty"`
		SysCreatedOn     string `json:"sys_created_on,omitempty"`
		SysID            string `json:"sys_id,omitempty"`
		SysModCount      string `json:"sys_mod_count,omitempty"`
		SysTags          string `json:"sys_tags,omitempty"`
		SysUpdatedBy     string `json:"sys_updated_by,omitempty"`
		SysUpdatedOn     string `json:"sys_updated_on,omitempty"`
		UActive          string `json:"u_active,omitempty"`           // Custom Field
		UApplicationPack string `json:"u_application_pack,omitempty"` // Custom Field
		UElement         string `json:"u_element,omitempty"`          // Custom Fieldx
	} `json:"result"`
}
