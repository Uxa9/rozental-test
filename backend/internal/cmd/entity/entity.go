package entity

var User = struct {
	Table string
	name  string
}{}

var InventoryOperation = struct {
	Table       string
	srcExecutor string
	dstExecutor string
	requestTime string
	statusTime  string
	status      string
}{}

var Inventory = struct {
	Table string
}{}

var InventoryOperationStatus = struct {
	Table  string
	status string
}{}

var InventoryOperationDetail = struct {
	Table string
}{}

func Init() {
	User.Table = "user"
	User.name = "name"

	InventoryOperation.Table = "inventory_operations"
	InventoryOperation.srcExecutor = "src_executor"
	InventoryOperation.dstExecutor = "dst_executor"
	InventoryOperation.requestTime = "request_time"
	InventoryOperation.statusTime = "status_time"
	InventoryOperation.status = "status"

	InventoryOperationDetail.Table = "inventory_operations_detail"

	InventoryOperationStatus.Table = "inventory_operations_status"

	Inventory.Table = "inventory"
}
