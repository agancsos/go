package actions

type InstanceInfoAction struct {}

func (x InstanceInfoAction) GenerateBody() string {
        return `<RetrieveServiceContent xmlns="urn:vim25">
<_this type="ServiceInstance">ServiceInstance</_this>
</RetrieveServiceContent>`;
}

type SessionManagerAction struct {}
func (x SessionManagerAction) GenerateBody() string {
	return `<RetrievePropertiesEx xmlns="urn:vim25">
            <_this type="PropertyCollector">ha-property-collector</_this>
            <specSet>
                <propSet>
                    <type>SessionManager</type>
                    <all>false</all>
                    <pathSet>currentSession</pathSet>
                </propSet>
                <objectSet>
                    <obj type="SessionManager">ha-sessionmgr</obj>
                    <skip>false</skip>
                </objectSet>
            </specSet>
            <options/>
        </RetrievePropertiesEx>`;
}

