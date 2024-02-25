package acq_files

import (
	"encoding/json"
	"fmt"
)

type ArnTransaction struct {
	TransactionReference string `json:"transaction_reference"`
	ClientTransactionId  string `json:"client_transaction_id"`
	TransactionArn       string `json:"transaction_arn"`
	Franchise            string `json:"franchise"`
	Status               string `json:"status"`
	Created              string `json:"created"`
}

func TestMarcial() {

	eventData := "{\"affiliation_code\":\"152000006378\",\"amount\":{\"currency\":\"CLP\",\"extra_taxes_info\":{},\"subtotal_iva_0\":7100},\"brand\":\"VISA\",\"client_transaction_id\":\"970955ca-0a04-4308-8888-b7ea291798ea\",\"config\":{\"latest_updated_at\":1708317409000,\"region\":\"us-east-1\"},\"country\":\"CHL\",\"created_at\":1708317409,\"merchant_id\":\"6000000000170188354540471776\",\"merchant_information_id\":\"20000000107178058000\",\"message_class\":\"authorizationClass\",\"pos_brand\":\"SUNMI\",\"pos_location\":{},\"pos_terminal_id\":\"1701883545000\",\"source\":\"chip\",\"transaction_mode\":\"Authorization\",\"transaction_reference\":\"2a1f250b-d40b-4422-92f8-ec84530dfb10\",\"transaction_scope\":\"international\",\"transaction_type\":\"Charge\",\"updated_at\":\"1708317413000\",\"vault_token\":\"EQZTOMZTGY3DGMJTHAZTQMZWGY2DGNRWGUZTCMZYGM2DGOBTGU3DIMZYGMZTGNJTGA3DKMZQGM4TMNJTHAZTGNRWGMYTGNBTHEZTAMZZERADMOJWG4YDCOJXGAYDMOJWHE4TSMBRHE3DSNZXGU3DSMBRHE3TGMBQGA3DINRXGY4TSNZWGUYDGMRWGY4TSNRTGE3DKNZTG43TMNJZHBACGMZSGMZTGNJWGU3DEMZRGMZTGNBWGUZTCNRWGM3TGNBWGM3DCNRVGYZDGMJTGIZTSMZSGM2DMMZWGEZTCNRSGY2TGNBTHEZTQMZVGY3SG===\",\"token_type\":\"transaction\",\"masked_pan\":\"476173XXXXXX0060\",\"f3\":\"000000\",\"f4\":\"000000007100\",\"f4_rq\":\"000000007100\",\"f7\":\"0219043649\",\"f11\":\"148957\",\"f12\":\"013649\",\"f13\":\"0219\",\"f18\":\"5555\",\"f19\":\"152\",\"f22\":\"0510\",\"f23\":\"001\",\"f25\":\"00\",\"f32\":\"454411\",\"f37\":\"405004148957\",\"f38\":\"123456\",\"f39\":\"00\",\"f41\":\"83545000\",\"f42\":\"152000006378   \",\"f43\":\"Copec                    SANTIAGO     CL\",\"f44\":\"V         P   2\",\"f48\":\"5CF1F0F4F6F6F2F2F8F2F9\",\"f49\":\"152\",\"f55_tags\":\"82021800950580800080009A032402099F02060000000050059F101706017703A000000F00564953414C3354455354434153459F260826BE74B43F66D1A79F330360F8C89F360200019F370441152B5C\",\"f60\":\"050000100000\",\"f61\":{\"f1\":\"\",\"f2\":\"\",\"f3\":\"\"},\"f62\":{\"f1\":\"\",\"f2\":\"900867429154501\",\"f20\":\"\"},\"f63\":{\"f1\":\"0000\"},\"f123\":\"6700030801f26800230202f0f0030bf4f0f0f1f0f0f7f5f3f3f80410f4f6f2f2f9f4f3f1f2f0f0f0f0f0f0f0\",\"f126\":{\"f13\":\"\",\"f20\":\"F\"},\"fr_message_time\":\"1708317409617\",\"fr_response_time\":\"1708317409624\",\"fr_message_code\":\"0110\",\"historic\":[{\"status\":\"PENDING\",\"updated_at\":1708317413000}],\"transaction_arn\":\"70000014050000001489577\",\"acquirer_bin\":\"454411\",\"transaction_status\":\"PENDING\",\"outgoing_path\":\"\",\"acquirer_key\":\"4761730060700000140500000014895770050000000007100\",\"total_amount\":7100,\"token123\":{\"token_assurance_level\":\"00\",\"token_requestor_id\":\"40010075338\",\"primary_account_number\":\"4622943120000000\"}}"
	var record ArnTransaction
	if err := json.Unmarshal([]byte(eventData), &record); err != nil {
		fmt.Println("Se cayo en el marshall")
	}
	fmt.Println("record: ", record)

}
