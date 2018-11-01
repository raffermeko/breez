// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package data is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	ChainStatus
	Account
	Payment
	PaymentsList
	SendNonDepositedCoinsRequest
	PayInvoiceRequest
	InvoiceMemo
	Invoice
	NotificationEvent
	AddFundReply
	FundStatusReply
	RemoveFundRequest
	RemoveFundReply
*/
package data

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account_AccountStatus int32

const (
	Account_WAITING_DEPOSIT              Account_AccountStatus = 0
	Account_WAITING_DEPOSIT_CONFIRMATION Account_AccountStatus = 1
	Account_PROCESSING_BREEZ_CONNECTION  Account_AccountStatus = 2
	Account_PROCESSING_WITHDRAWAL        Account_AccountStatus = 3
	Account_ACTIVE                       Account_AccountStatus = 4
)

var Account_AccountStatus_name = map[int32]string{
	0: "WAITING_DEPOSIT",
	1: "WAITING_DEPOSIT_CONFIRMATION",
	2: "PROCESSING_BREEZ_CONNECTION",
	3: "PROCESSING_WITHDRAWAL",
	4: "ACTIVE",
}
var Account_AccountStatus_value = map[string]int32{
	"WAITING_DEPOSIT":              0,
	"WAITING_DEPOSIT_CONFIRMATION": 1,
	"PROCESSING_BREEZ_CONNECTION":  2,
	"PROCESSING_WITHDRAWAL":        3,
	"ACTIVE":                       4,
}

func (x Account_AccountStatus) String() string {
	return proto.EnumName(Account_AccountStatus_name, int32(x))
}
func (Account_AccountStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Payment_PaymentType int32

const (
	Payment_DEPOSIT    Payment_PaymentType = 0
	Payment_WITHDRAWAL Payment_PaymentType = 1
	Payment_SENT       Payment_PaymentType = 2
	Payment_RECEIVED   Payment_PaymentType = 3
)

var Payment_PaymentType_name = map[int32]string{
	0: "DEPOSIT",
	1: "WITHDRAWAL",
	2: "SENT",
	3: "RECEIVED",
}
var Payment_PaymentType_value = map[string]int32{
	"DEPOSIT":    0,
	"WITHDRAWAL": 1,
	"SENT":       2,
	"RECEIVED":   3,
}

func (x Payment_PaymentType) String() string {
	return proto.EnumName(Payment_PaymentType_name, int32(x))
}
func (Payment_PaymentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type NotificationEvent_NotificationType int32

const (
	NotificationEvent_READY                           NotificationEvent_NotificationType = 0
	NotificationEvent_INITIALIZATION_FAILED           NotificationEvent_NotificationType = 1
	NotificationEvent_ACCOUNT_CHANGED                 NotificationEvent_NotificationType = 2
	NotificationEvent_INVOICE_PAID                    NotificationEvent_NotificationType = 3
	NotificationEvent_ROUTING_NODE_CONNECTION_CHANGED NotificationEvent_NotificationType = 4
	NotificationEvent_LIGHTNING_SERVICE_DOWN          NotificationEvent_NotificationType = 5
)

var NotificationEvent_NotificationType_name = map[int32]string{
	0: "READY",
	1: "INITIALIZATION_FAILED",
	2: "ACCOUNT_CHANGED",
	3: "INVOICE_PAID",
	4: "ROUTING_NODE_CONNECTION_CHANGED",
	5: "LIGHTNING_SERVICE_DOWN",
}
var NotificationEvent_NotificationType_value = map[string]int32{
	"READY":                           0,
	"INITIALIZATION_FAILED":           1,
	"ACCOUNT_CHANGED":                 2,
	"INVOICE_PAID":                    3,
	"ROUTING_NODE_CONNECTION_CHANGED": 4,
	"LIGHTNING_SERVICE_DOWN":          5,
}

func (x NotificationEvent_NotificationType) String() string {
	return proto.EnumName(NotificationEvent_NotificationType_name, int32(x))
}
func (NotificationEvent_NotificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

type FundStatusReply_FundStatus int32

const (
	FundStatusReply_NO_FUND              FundStatusReply_FundStatus = 0
	FundStatusReply_WAITING_CONFIRMATION FundStatusReply_FundStatus = 1
	FundStatusReply_CONFIRMED            FundStatusReply_FundStatus = 2
)

var FundStatusReply_FundStatus_name = map[int32]string{
	0: "NO_FUND",
	1: "WAITING_CONFIRMATION",
	2: "CONFIRMED",
}
var FundStatusReply_FundStatus_value = map[string]int32{
	"NO_FUND":              0,
	"WAITING_CONFIRMATION": 1,
	"CONFIRMED":            2,
}

func (x FundStatusReply_FundStatus) String() string {
	return proto.EnumName(FundStatusReply_FundStatus_name, int32(x))
}
func (FundStatusReply_FundStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

type ChainStatus struct {
	BlockHeight   uint32 `protobuf:"varint,1,opt,name=blockHeight" json:"blockHeight,omitempty"`
	SyncedToChain bool   `protobuf:"varint,2,opt,name=syncedToChain" json:"syncedToChain,omitempty"`
}

func (m *ChainStatus) Reset()                    { *m = ChainStatus{} }
func (m *ChainStatus) String() string            { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()               {}
func (*ChainStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChainStatus) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ChainStatus) GetSyncedToChain() bool {
	if m != nil {
		return m.SyncedToChain
	}
	return false
}

type Account struct {
	Id                    string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Balance               int64                 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	NonDepositableBalance int64                 `protobuf:"varint,3,opt,name=nonDepositableBalance" json:"nonDepositableBalance,omitempty"`
	Status                Account_AccountStatus `protobuf:"varint,4,opt,name=status,enum=data.Account_AccountStatus" json:"status,omitempty"`
	// maximum payment this node can receive via lightning
	MaxAllowedToReceive int64 `protobuf:"varint,5,opt,name=maxAllowedToReceive" json:"maxAllowedToReceive,omitempty"`
	// maximum payment this node can pay via lightning
	MaxAllowedToPay int64 `protobuf:"varint,6,opt,name=maxAllowedToPay" json:"maxAllowedToPay,omitempty"`
	// The lightning absolute payment amount
	MaxPaymentAmount int64 `protobuf:"varint,7,opt,name=maxPaymentAmount" json:"maxPaymentAmount,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetNonDepositableBalance() int64 {
	if m != nil {
		return m.NonDepositableBalance
	}
	return 0
}

func (m *Account) GetStatus() Account_AccountStatus {
	if m != nil {
		return m.Status
	}
	return Account_WAITING_DEPOSIT
}

func (m *Account) GetMaxAllowedToReceive() int64 {
	if m != nil {
		return m.MaxAllowedToReceive
	}
	return 0
}

func (m *Account) GetMaxAllowedToPay() int64 {
	if m != nil {
		return m.MaxAllowedToPay
	}
	return 0
}

func (m *Account) GetMaxPaymentAmount() int64 {
	if m != nil {
		return m.MaxPaymentAmount
	}
	return 0
}

type Payment struct {
	Type              Payment_PaymentType `protobuf:"varint,1,opt,name=type,enum=data.Payment_PaymentType" json:"type,omitempty"`
	Amount            int64               `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	CreationTimestamp int64               `protobuf:"varint,4,opt,name=creationTimestamp" json:"creationTimestamp,omitempty"`
	InvoiceMemo       *InvoiceMemo        `protobuf:"bytes,6,opt,name=invoiceMemo" json:"invoiceMemo,omitempty"`
	RedeemTxID        string              `protobuf:"bytes,7,opt,name=redeemTxID" json:"redeemTxID,omitempty"`
	PaymentHash       string              `protobuf:"bytes,8,opt,name=paymentHash" json:"paymentHash,omitempty"`
	Destination       string              `protobuf:"bytes,9,opt,name=destination" json:"destination,omitempty"`
}

func (m *Payment) Reset()                    { *m = Payment{} }
func (m *Payment) String() string            { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()               {}
func (*Payment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Payment) GetType() Payment_PaymentType {
	if m != nil {
		return m.Type
	}
	return Payment_DEPOSIT
}

func (m *Payment) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Payment) GetCreationTimestamp() int64 {
	if m != nil {
		return m.CreationTimestamp
	}
	return 0
}

func (m *Payment) GetInvoiceMemo() *InvoiceMemo {
	if m != nil {
		return m.InvoiceMemo
	}
	return nil
}

func (m *Payment) GetRedeemTxID() string {
	if m != nil {
		return m.RedeemTxID
	}
	return ""
}

func (m *Payment) GetPaymentHash() string {
	if m != nil {
		return m.PaymentHash
	}
	return ""
}

func (m *Payment) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type PaymentsList struct {
	PaymentsList []*Payment `protobuf:"bytes,1,rep,name=paymentsList" json:"paymentsList,omitempty"`
}

func (m *PaymentsList) Reset()                    { *m = PaymentsList{} }
func (m *PaymentsList) String() string            { return proto.CompactTextString(m) }
func (*PaymentsList) ProtoMessage()               {}
func (*PaymentsList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PaymentsList) GetPaymentsList() []*Payment {
	if m != nil {
		return m.PaymentsList
	}
	return nil
}

type SendNonDepositedCoinsRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *SendNonDepositedCoinsRequest) Reset()                    { *m = SendNonDepositedCoinsRequest{} }
func (m *SendNonDepositedCoinsRequest) String() string            { return proto.CompactTextString(m) }
func (*SendNonDepositedCoinsRequest) ProtoMessage()               {}
func (*SendNonDepositedCoinsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SendNonDepositedCoinsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PayInvoiceRequest struct {
	Amount         int64  `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	PaymentRequest string `protobuf:"bytes,2,opt,name=paymentRequest" json:"paymentRequest,omitempty"`
}

func (m *PayInvoiceRequest) Reset()                    { *m = PayInvoiceRequest{} }
func (m *PayInvoiceRequest) String() string            { return proto.CompactTextString(m) }
func (*PayInvoiceRequest) ProtoMessage()               {}
func (*PayInvoiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PayInvoiceRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PayInvoiceRequest) GetPaymentRequest() string {
	if m != nil {
		return m.PaymentRequest
	}
	return ""
}

type InvoiceMemo struct {
	Description     string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Amount          int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	PayeeName       string `protobuf:"bytes,3,opt,name=payeeName" json:"payeeName,omitempty"`
	PayeeImageURL   string `protobuf:"bytes,4,opt,name=payeeImageURL" json:"payeeImageURL,omitempty"`
	PayerName       string `protobuf:"bytes,5,opt,name=payerName" json:"payerName,omitempty"`
	PayerImageURL   string `protobuf:"bytes,6,opt,name=payerImageURL" json:"payerImageURL,omitempty"`
	TransferRequest bool   `protobuf:"varint,7,opt,name=transferRequest" json:"transferRequest,omitempty"`
	Expiry          int64  `protobuf:"varint,8,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *InvoiceMemo) Reset()                    { *m = InvoiceMemo{} }
func (m *InvoiceMemo) String() string            { return proto.CompactTextString(m) }
func (*InvoiceMemo) ProtoMessage()               {}
func (*InvoiceMemo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InvoiceMemo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvoiceMemo) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InvoiceMemo) GetPayeeName() string {
	if m != nil {
		return m.PayeeName
	}
	return ""
}

func (m *InvoiceMemo) GetPayeeImageURL() string {
	if m != nil {
		return m.PayeeImageURL
	}
	return ""
}

func (m *InvoiceMemo) GetPayerName() string {
	if m != nil {
		return m.PayerName
	}
	return ""
}

func (m *InvoiceMemo) GetPayerImageURL() string {
	if m != nil {
		return m.PayerImageURL
	}
	return ""
}

func (m *InvoiceMemo) GetTransferRequest() bool {
	if m != nil {
		return m.TransferRequest
	}
	return false
}

func (m *InvoiceMemo) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type Invoice struct {
	Memo    *InvoiceMemo `protobuf:"bytes,1,opt,name=memo" json:"memo,omitempty"`
	Settled bool         `protobuf:"varint,2,opt,name=settled" json:"settled,omitempty"`
	AmtPaid int64        `protobuf:"varint,3,opt,name=amtPaid" json:"amtPaid,omitempty"`
}

func (m *Invoice) Reset()                    { *m = Invoice{} }
func (m *Invoice) String() string            { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()               {}
func (*Invoice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Invoice) GetMemo() *InvoiceMemo {
	if m != nil {
		return m.Memo
	}
	return nil
}

func (m *Invoice) GetSettled() bool {
	if m != nil {
		return m.Settled
	}
	return false
}

func (m *Invoice) GetAmtPaid() int64 {
	if m != nil {
		return m.AmtPaid
	}
	return 0
}

type NotificationEvent struct {
	Type NotificationEvent_NotificationType `protobuf:"varint,1,opt,name=type,enum=data.NotificationEvent_NotificationType" json:"type,omitempty"`
}

func (m *NotificationEvent) Reset()                    { *m = NotificationEvent{} }
func (m *NotificationEvent) String() string            { return proto.CompactTextString(m) }
func (*NotificationEvent) ProtoMessage()               {}
func (*NotificationEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NotificationEvent) GetType() NotificationEvent_NotificationType {
	if m != nil {
		return m.Type
	}
	return NotificationEvent_READY
}

type AddFundReply struct {
	Address           string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	MaxAllowedDeposit int64  `protobuf:"varint,2,opt,name=maxAllowedDeposit" json:"maxAllowedDeposit,omitempty"`
	ErrorMessage      string `protobuf:"bytes,3,opt,name=errorMessage" json:"errorMessage,omitempty"`
}

func (m *AddFundReply) Reset()                    { *m = AddFundReply{} }
func (m *AddFundReply) String() string            { return proto.CompactTextString(m) }
func (*AddFundReply) ProtoMessage()               {}
func (*AddFundReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddFundReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddFundReply) GetMaxAllowedDeposit() int64 {
	if m != nil {
		return m.MaxAllowedDeposit
	}
	return 0
}

func (m *AddFundReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type FundStatusReply struct {
	Status FundStatusReply_FundStatus `protobuf:"varint,1,opt,name=status,enum=data.FundStatusReply_FundStatus" json:"status,omitempty"`
}

func (m *FundStatusReply) Reset()                    { *m = FundStatusReply{} }
func (m *FundStatusReply) String() string            { return proto.CompactTextString(m) }
func (*FundStatusReply) ProtoMessage()               {}
func (*FundStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FundStatusReply) GetStatus() FundStatusReply_FundStatus {
	if m != nil {
		return m.Status
	}
	return FundStatusReply_NO_FUND
}

type RemoveFundRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Amount  int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *RemoveFundRequest) Reset()                    { *m = RemoveFundRequest{} }
func (m *RemoveFundRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveFundRequest) ProtoMessage()               {}
func (*RemoveFundRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RemoveFundRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RemoveFundRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type RemoveFundReply struct {
	Txid         string `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage" json:"errorMessage,omitempty"`
}

func (m *RemoveFundReply) Reset()                    { *m = RemoveFundReply{} }
func (m *RemoveFundReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveFundReply) ProtoMessage()               {}
func (*RemoveFundReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RemoveFundReply) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *RemoveFundReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ChainStatus)(nil), "data.ChainStatus")
	proto.RegisterType((*Account)(nil), "data.Account")
	proto.RegisterType((*Payment)(nil), "data.Payment")
	proto.RegisterType((*PaymentsList)(nil), "data.PaymentsList")
	proto.RegisterType((*SendNonDepositedCoinsRequest)(nil), "data.SendNonDepositedCoinsRequest")
	proto.RegisterType((*PayInvoiceRequest)(nil), "data.PayInvoiceRequest")
	proto.RegisterType((*InvoiceMemo)(nil), "data.InvoiceMemo")
	proto.RegisterType((*Invoice)(nil), "data.Invoice")
	proto.RegisterType((*NotificationEvent)(nil), "data.NotificationEvent")
	proto.RegisterType((*AddFundReply)(nil), "data.AddFundReply")
	proto.RegisterType((*FundStatusReply)(nil), "data.FundStatusReply")
	proto.RegisterType((*RemoveFundRequest)(nil), "data.RemoveFundRequest")
	proto.RegisterType((*RemoveFundReply)(nil), "data.RemoveFundReply")
	proto.RegisterEnum("data.Account_AccountStatus", Account_AccountStatus_name, Account_AccountStatus_value)
	proto.RegisterEnum("data.Payment_PaymentType", Payment_PaymentType_name, Payment_PaymentType_value)
	proto.RegisterEnum("data.NotificationEvent_NotificationType", NotificationEvent_NotificationType_name, NotificationEvent_NotificationType_value)
	proto.RegisterEnum("data.FundStatusReply_FundStatus", FundStatusReply_FundStatus_name, FundStatusReply_FundStatus_value)
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1019 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x4e, 0xe3, 0x46,
	0x14, 0xc6, 0x49, 0x48, 0xc8, 0x49, 0x08, 0x66, 0xb6, 0xbb, 0xf2, 0x76, 0x51, 0x37, 0x72, 0x7f,
	0x14, 0x55, 0x2d, 0x6a, 0xa1, 0x17, 0x7b, 0xd1, 0x1b, 0x63, 0x1b, 0x18, 0x29, 0x38, 0xd1, 0xc4,
	0x80, 0xba, 0x37, 0xd1, 0x10, 0xcf, 0x82, 0xd5, 0xf8, 0xa7, 0xb6, 0xa1, 0x49, 0x1f, 0xa2, 0x55,
	0xdf, 0xa0, 0x6f, 0xd1, 0x87, 0xe9, 0x6b, 0xf4, 0x01, 0xaa, 0x19, 0x8f, 0x13, 0x3b, 0x59, 0xd4,
	0xab, 0x70, 0xbe, 0xf3, 0xf9, 0x9c, 0x99, 0xef, 0xfc, 0x0c, 0xd0, 0x0b, 0x58, 0x9a, 0xd2, 0x7b,
	0x96, 0x1e, 0xc7, 0x49, 0x94, 0x45, 0xa8, 0xe1, 0xd1, 0x8c, 0xea, 0xd7, 0xd0, 0x31, 0x1f, 0xa8,
	0x1f, 0x4e, 0x32, 0x9a, 0x3d, 0xa6, 0xa8, 0x0f, 0x9d, 0xbb, 0x79, 0x34, 0xfb, 0xf9, 0x92, 0xf9,
	0xf7, 0x0f, 0x99, 0xa6, 0xf4, 0x95, 0xc1, 0x3e, 0x29, 0x43, 0xe8, 0x0b, 0xd8, 0x4f, 0x97, 0xe1,
	0x8c, 0x79, 0x6e, 0x24, 0x3e, 0xd4, 0x6a, 0x7d, 0x65, 0xb0, 0x47, 0xaa, 0xa0, 0xfe, 0x77, 0x1d,
	0x5a, 0xc6, 0x6c, 0x16, 0x3d, 0x86, 0x19, 0xea, 0x41, 0xcd, 0xf7, 0x44, 0xa8, 0x36, 0xa9, 0xf9,
	0x1e, 0xd2, 0xa0, 0x75, 0x47, 0xe7, 0x34, 0x9c, 0x31, 0xf1, 0x6d, 0x9d, 0x14, 0x26, 0xfa, 0x01,
	0x5e, 0x86, 0x51, 0x68, 0xb1, 0x38, 0x4a, 0xfd, 0x8c, 0xde, 0xcd, 0xd9, 0x99, 0xe4, 0xd5, 0x05,
	0xef, 0xe3, 0x4e, 0x74, 0x0a, 0xcd, 0x54, 0x9c, 0x5e, 0x6b, 0xf4, 0x95, 0x41, 0xef, 0xe4, 0xcd,
	0x31, 0xbf, 0xd9, 0xb1, 0x4c, 0x5f, 0xfc, 0xe6, 0x17, 0x24, 0x92, 0x8a, 0xbe, 0x83, 0x17, 0x01,
	0x5d, 0x18, 0xf3, 0x79, 0xf4, 0x2b, 0x3f, 0x35, 0x61, 0x33, 0xe6, 0x3f, 0x31, 0x6d, 0x57, 0x24,
	0xfa, 0x98, 0x0b, 0x0d, 0xe0, 0xa0, 0x0c, 0x8f, 0xe9, 0x52, 0x6b, 0x0a, 0xf6, 0x26, 0x8c, 0xbe,
	0x06, 0x35, 0xa0, 0x8b, 0x31, 0x5d, 0x06, 0x2c, 0xcc, 0x8c, 0x80, 0x67, 0xd7, 0x5a, 0x82, 0xba,
	0x85, 0xeb, 0xbf, 0x2b, 0xb0, 0x5f, 0x39, 0x21, 0x7a, 0x01, 0x07, 0xb7, 0x06, 0x76, 0xb1, 0x73,
	0x31, 0xb5, 0xec, 0xf1, 0x68, 0x82, 0x5d, 0x75, 0x07, 0xf5, 0xe1, 0x68, 0x03, 0x9c, 0x9a, 0x23,
	0xe7, 0x1c, 0x93, 0x2b, 0xc3, 0xc5, 0x23, 0x47, 0x55, 0xd0, 0x5b, 0x78, 0x33, 0x26, 0x23, 0xd3,
	0x9e, 0x4c, 0x38, 0xe9, 0x8c, 0xd8, 0xf6, 0x7b, 0x4e, 0x71, 0x6c, 0x53, 0x10, 0x6a, 0xe8, 0x35,
	0xbc, 0x2c, 0x11, 0x6e, 0xb1, 0x7b, 0x69, 0x11, 0xe3, 0xd6, 0x18, 0xaa, 0x75, 0x04, 0xd0, 0x34,
	0x4c, 0x17, 0xdf, 0xd8, 0x6a, 0x43, 0xff, 0xa7, 0x06, 0x2d, 0x79, 0x44, 0xf4, 0x2d, 0x34, 0xb2,
	0x65, 0xcc, 0x44, 0xed, 0x7a, 0x27, 0xaf, 0x73, 0x5d, 0xa5, 0xb3, 0xf8, 0x75, 0x97, 0x31, 0x23,
	0x82, 0x86, 0x5e, 0x41, 0x93, 0xe6, 0xb7, 0xcd, 0xeb, 0x25, 0x2d, 0xf4, 0x0d, 0x1c, 0xce, 0x12,
	0x46, 0x33, 0x3f, 0x0a, 0x5d, 0x3f, 0x60, 0x69, 0x46, 0x83, 0x58, 0xd4, 0xaa, 0x4e, 0xb6, 0x1d,
	0xe8, 0x14, 0x3a, 0x7e, 0xf8, 0x14, 0xf9, 0x33, 0x76, 0xc5, 0x82, 0x48, 0x68, 0xdc, 0x39, 0x39,
	0xcc, 0x73, 0xe3, 0xb5, 0x83, 0x94, 0x59, 0xe8, 0x33, 0x80, 0x84, 0x79, 0x8c, 0x05, 0xee, 0x02,
	0x5b, 0x42, 0xec, 0x36, 0x29, 0x21, 0xbc, 0xaf, 0xe3, 0xfc, 0xbc, 0x97, 0x34, 0x7d, 0xd0, 0xf6,
	0x04, 0xa1, 0x0c, 0x71, 0x86, 0xc7, 0xd2, 0xcc, 0x0f, 0xc5, 0x71, 0xb4, 0x76, 0xce, 0x28, 0x41,
	0xfa, 0x19, 0x74, 0x4a, 0x77, 0x46, 0x1d, 0x68, 0xad, 0xeb, 0xd3, 0x03, 0x28, 0x29, 0xaa, 0xa0,
	0x3d, 0x68, 0x4c, 0x6c, 0xc7, 0x55, 0x6b, 0xa8, 0x0b, 0x7b, 0xc4, 0x36, 0x6d, 0x7c, 0x63, 0x5b,
	0x6a, 0x5d, 0x37, 0xa0, 0x2b, 0x63, 0xa4, 0x43, 0x3f, 0xcd, 0xd0, 0xf7, 0xd0, 0x8d, 0x4b, 0xb6,
	0xa6, 0xf4, 0xeb, 0x83, 0xce, 0xc9, 0x7e, 0x45, 0x69, 0x52, 0xa1, 0xe8, 0xef, 0xe0, 0x68, 0xc2,
	0x42, 0xcf, 0x59, 0xcd, 0x02, 0xf3, 0xcc, 0xc8, 0x0f, 0x53, 0xc2, 0x7e, 0x79, 0x64, 0x69, 0xc6,
	0xc7, 0x8b, 0x7a, 0x5e, 0xc2, 0xd2, 0x54, 0xce, 0x5c, 0x61, 0xea, 0x13, 0x38, 0x1c, 0xd3, 0xa5,
	0xd4, 0xb0, 0xa0, 0xaf, 0x8b, 0xa6, 0x54, 0x8a, 0xf6, 0x15, 0xf4, 0x64, 0x5a, 0xc9, 0x14, 0xc3,
	0xda, 0x26, 0x1b, 0xa8, 0xfe, 0x67, 0x0d, 0x3a, 0xa5, 0xb2, 0x48, 0x1d, 0x67, 0x89, 0x1f, 0x0b,
	0x1d, 0x95, 0x95, 0x8e, 0x05, 0x54, 0xca, 0x58, 0xab, 0x64, 0x3c, 0x82, 0x76, 0x4c, 0x97, 0x8c,
	0x39, 0x34, 0xc8, 0x27, 0xbe, 0x4d, 0xd6, 0x00, 0xdf, 0x3b, 0xc2, 0xc0, 0x01, 0xbd, 0x67, 0xd7,
	0x64, 0x28, 0x1a, 0xa8, 0x4d, 0xaa, 0x60, 0x11, 0x23, 0x11, 0x31, 0x76, 0xd7, 0x31, 0x92, 0x72,
	0x8c, 0x64, 0x15, 0xa3, 0xb9, 0x8e, 0xb1, 0x02, 0xf9, 0xa0, 0x67, 0x09, 0x0d, 0xd3, 0x0f, 0x2c,
	0x29, 0xae, 0xde, 0x12, 0x3b, 0x6e, 0x13, 0xe6, 0x37, 0x61, 0x8b, 0xd8, 0x4f, 0x96, 0xa2, 0xa1,
	0xea, 0x44, 0x5a, 0xba, 0x07, 0x2d, 0x29, 0x09, 0xfa, 0x12, 0x1a, 0x01, 0x6f, 0x63, 0xe5, 0xb9,
	0x36, 0x16, 0x6e, 0x5e, 0xb4, 0x94, 0x65, 0xd9, 0x9c, 0x79, 0x72, 0x9f, 0x16, 0xa6, 0x28, 0x67,
	0x90, 0x8d, 0xa9, 0xef, 0xc9, 0xa9, 0x2a, 0x4c, 0xfd, 0x5f, 0x05, 0x0e, 0x9d, 0x28, 0xf3, 0x3f,
	0xf8, 0x33, 0xd1, 0xa0, 0xf6, 0x13, 0x9f, 0xd9, 0x1f, 0x2b, 0x33, 0x3b, 0xc8, 0x13, 0x6e, 0xd1,
	0x2a, 0xc8, 0x7a, 0x84, 0xf5, 0xbf, 0x14, 0x50, 0x37, 0x5d, 0xa8, 0x0d, 0xbb, 0xc4, 0x36, 0xac,
	0x9f, 0xd4, 0x1d, 0xbe, 0x44, 0xb0, 0x83, 0x5d, 0x6c, 0x0c, 0xf1, 0x7b, 0xb1, 0x79, 0xa6, 0xe7,
	0x06, 0x1e, 0xda, 0x96, 0xaa, 0xf0, 0xbd, 0x65, 0x98, 0xe6, 0xe8, 0xda, 0x71, 0xa7, 0xe6, 0xa5,
	0xe1, 0x5c, 0xd8, 0x96, 0x5a, 0x43, 0x2a, 0x74, 0xb1, 0x73, 0x33, 0xc2, 0xa6, 0x3d, 0x1d, 0x1b,
	0xd8, 0x52, 0xeb, 0xe8, 0x73, 0x78, 0x4b, 0x46, 0xd7, 0x62, 0x93, 0x39, 0x23, 0xcb, 0x2e, 0xed,
	0xa8, 0xd5, 0x67, 0x0d, 0xf4, 0x29, 0xbc, 0x1a, 0xe2, 0x8b, 0x4b, 0xd7, 0xe1, 0xb4, 0x89, 0x4d,
	0x6e, 0x78, 0x00, 0x6b, 0x74, 0xeb, 0xa8, 0xbb, 0xfa, 0x6f, 0xd0, 0x35, 0x3c, 0xef, 0xfc, 0x31,
	0xf4, 0x08, 0x8b, 0xe7, 0xcb, 0xe7, 0xfb, 0x9d, 0xef, 0x9d, 0xf5, 0x6a, 0x96, 0xc3, 0x22, 0x7b,
	0x6e, 0xdb, 0x81, 0x74, 0xe8, 0xb2, 0x24, 0x89, 0x92, 0xab, 0xfc, 0x99, 0x94, 0x1d, 0x58, 0xc1,
	0xf4, 0x3f, 0x14, 0x38, 0xe0, 0x99, 0xe5, 0x63, 0x22, 0xf2, 0xbf, 0x5b, 0x3d, 0x3f, 0xb9, 0xe4,
	0xfd, 0x5c, 0xf2, 0x0d, 0x5a, 0xd9, 0x96, 0x7c, 0xfd, 0x0c, 0x60, 0x8d, 0xf2, 0x7d, 0xe2, 0x8c,
	0xa6, 0xe7, 0xd7, 0x8e, 0xa5, 0xee, 0x20, 0x0d, 0x3e, 0x29, 0xf6, 0xfd, 0xc6, 0x9e, 0xdf, 0x87,
	0xb6, 0x44, 0xb8, 0xc0, 0xba, 0x0d, 0x87, 0x84, 0x05, 0xd1, 0x13, 0xcb, 0x05, 0xf9, 0x9f, 0x15,
	0xf0, 0xdc, 0xec, 0xe9, 0x18, 0x0e, 0xca, 0x61, 0xf8, 0xbd, 0x10, 0x34, 0xb2, 0xc5, 0xea, 0xe1,
	0x16, 0x7f, 0x6f, 0x69, 0x54, 0xdb, 0xd6, 0xe8, 0xae, 0x29, 0xfe, 0xbd, 0x38, 0xfd, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x98, 0x8c, 0xbb, 0x63, 0x70, 0x08, 0x00, 0x00,
}
