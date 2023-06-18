package actions
import (
	"fmt"
	"strings"
)

type RetrieveDataStoresAction struct {}
func (x RetrieveDataStoresAction) GenerateBody() string {
	return `<RetrievePropertiesEx xmlns="urn:vim25">
    <_this type="PropertyCollector">ha-property-collector</_this>
    <specSet>
        <propSet>
            <type>HostSystem</type>
            <all>false</all>
            <pathSet>config.autoStart</pathSet>
        </propSet>
        <objectSet>
            <obj type="HostSystem">ha-host</obj>
            <skip>false</skip>
        </objectSet>
    </specSet>
    <options/>
</RetrievePropertiesEx>`;
}

type RetrieveDataFoldersAction struct {}
func (x RetrieveDataFoldersAction) GeenerateBody() string {
	return `<RetrievePropertiesEx xmlns="urn:vim25">
    <_this type="PropertyCollector">ha-property-collector</_this>
    <specSet>
        <propSet>
            <type>Folder</type>
            <all>false</all>
            <pathSet>childEntity</pathSet>
        </propSet>
        <objectSet>
            <obj type="Folder">ha-folder-vm</obj>
            <skip>false</skip>
        </objectSet>
    </specSet>
    <options/>
</RetrievePropertiesEx>`;
}

type RetrieveVirtualMachinesAction struct {}
func (x RetrieveVirtualMachinesAction) GenerateBody() string {
	return `<RetrievePropertiesEx xmlns="urn:vim25">
    <_this type="PropertyCollector">ha-property-collector</_this>
    <specSet>
        <propSet>
            <type>HostSystem</type>
            <all>false</all>
            <pathSet>config.autoStart</pathSet>
        </propSet>
        <objectSet>
            <obj type="HostSystem">ha-host</obj>
            <skip>false</skip>
        </objectSet>
    </specSet>
    <options/>
</RetrievePropertiesEx>`;
}

type deviceChange struct {
	operation       string
	typeClass       string
	properties      map[string]interface{}
}

func NewDeviceChange(operation string, typeClass string, properties map[string]interface{}) *deviceChange {
	var rst = &deviceChange{};
	rst.operation  = operation;
	rst.typeClass  = typeClass;
	rst.properties = properties;
	return rst;
}

func (x deviceChange) GenerateProperties(properties map[string]interface{}) string {
	var rst = "";
	for k, v := range properties {
		rst += fmt.Sprintf("<%s>", k);
		_, isMap := v.(map[string]interface{}); if isMap {
			rst += x.GenerateProperties(v.(map[string]interface{}));
		} else {
			rst += fmt.Sprintf("%v", v);
		}
		rst += fmt.Sprintf("</%s>", k);
	}
	return rst;
}

func (x deviceChange) GenerateBody() {
	var rst = fmt.Sprintf(`<deviceChange>
<operation>%s</operation>
<device xsi:type="%s">`, x.operation, x.typeClass);
	rst += x.GenerateProperties(x.properties);
	rst += `</device>
        </deviceChange`;
	return rst;
}

// TODO: Implement CRUD operations
type CreateVirtualMachineAction struct {
	properties      map[string]interface{}
	changes         []*deviceChange
}

func (x CreateVirtualMachineAction) GenerateBody() string {
	var rst `<CreateVM_Task xmlns="urn:vim25">
    <_this type="Folder">ha-folder-vm</_this>
    <config>`;
	rst += NewDeviceChange("", "", x.properties).GenerateProperties(x.properties);
	for _, change := range x.changes {
		rst += fmt.Sprintf("%s\n", change.GenerateBody());
	}
	rst += `</config>
    <pool type="ResourcePool">ha-root-pool</pool>
</CreateVM_Task>`;
	return rst;
}

func (x *CreateVirtualMachineAction) SetProperties(properties map[string]interface{}) { x.properties = properties; }
func (x *CreateVirtualMachineAction) SetChanges(changes []*deviceChange) { x.changes = changes; }
func (x CreateVirtualMachineAction) Properties() map[string]interface{} { return x.properties; }
func (x CreateVirtualMachineAction) Changes() []*deviceChange { return x.changes; }

type UpdateVirtualMachineAction struct {
	vm              int
	properties      map[string]interface{} 
	changes         []*deviceChange
}

func (x UpdateVirtualMachineAction) GenerateBody() string {
	var rst fmt.Sprintf(`<ReconfigVM_Task xmlns="urn:vim25">
    <_this type="VirtualMachine">{0}</_this>
    <spec>`, x.vm);
	rst += NewDeviceChange("", "", x.properties).GenerateProperties(x.properties);
    for _, change := range x.changes {
        rst += fmt.Sprintf("%s\n", change.GenerateBody());
    }
    rst += `</spec>
</ReconfigVM_Task>`;
	return rst;
}
func (x *UpdateVirtualMachineAction) SetVM(vm int) { x.vm = vm; }
func (x *UpdateVirtualMachineAction) SetProperties(properties map[string]interface{}) { x.properties = properties; }
func (x *UpdateVirtualMachineAction) SetChanges(changes []*deviceChange) { x.changes = changes; }
func (x UpdateVirtualMachineAction) Properties() map[string]interface{} { return x.properties; }
func (x UpdateVirtualMachineAction) Changes() []*deviceChange { return x.changes; }
func (x UpdateVirtualMachineAction) VM() int { return x.vm; }

type RemoveVirtualMachineAction struct {
	vm             int
}

func (x RemoveVirtualMachineAction) GenerateBody() string {
	return fmt.Sprintf(`<Destroy_Task xmlns="urn:vim25">
    <_this type="VirtualMachine">%d</_this>
</Destroy_Task>`, x.vm);
}

func (x *RemoveVirtualMachineAction) SetVM(vm int) { x.vm = vm; }
func (x RemoveVirtualMachineAction) VM() int { return x.vm; }

