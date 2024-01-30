// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: coordinator.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Catagories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catagories []*Category `protobuf:"bytes,1,rep,name=Catagories,proto3" json:"Catagories,omitempty"`
}

func (x *Catagories) Reset() {
	*x = Catagories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catagories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catagories) ProtoMessage() {}

func (x *Catagories) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catagories.ProtoReflect.Descriptor instead.
func (*Catagories) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *Catagories) GetCatagories() []*Category {
	if x != nil {
		return x.Catagories
	}
	return nil
}

type Responce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Responce) Reset() {
	*x = Responce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Responce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Responce) ProtoMessage() {}

func (x *Responce) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Responce.ProtoReflect.Descriptor instead.
func (*Responce) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *Responce) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Responce) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type View struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Page   int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *View) Reset() {
	*x = View{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View) ProtoMessage() {}

func (x *View) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View.ProtoReflect.Descriptor instead.
func (*View) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{2}
}

func (x *View) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *View) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *View) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Packages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Packages) Reset() {
	*x = Packages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packages) ProtoMessage() {}

func (x *Packages) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packages.ProtoReflect.Descriptor instead.
func (*Packages) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{3}
}

type PackagesResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages []*Package `protobuf:"bytes,1,rep,name=Packages,proto3" json:"Packages,omitempty"`
}

func (x *PackagesResponce) Reset() {
	*x = PackagesResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackagesResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackagesResponce) ProtoMessage() {}

func (x *PackagesResponce) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackagesResponce.ProtoReflect.Descriptor instead.
func (*PackagesResponce) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{4}
}

func (x *PackagesResponce) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatagoryId   int64  `protobuf:"varint,1,opt,name=CatagoryId,proto3" json:"CatagoryId,omitempty"`
	CategoryName string `protobuf:"bytes,2,opt,name=categoryName,proto3" json:"categoryName,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{5}
}

func (x *Category) GetCatagoryId() int64 {
	if x != nil {
		return x.CatagoryId
	}
	return 0
}

func (x *Category) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId        int64          `protobuf:"varint,1,opt,name=PackageId,proto3" json:"PackageId,omitempty"`
	Packagename      string         `protobuf:"bytes,2,opt,name=packagename,proto3" json:"packagename,omitempty"`
	CoorinatorId     int64          `protobuf:"varint,3,opt,name=CoorinatorId,proto3" json:"CoorinatorId,omitempty"`
	Startlocation    string         `protobuf:"bytes,4,opt,name=startlocation,proto3" json:"startlocation,omitempty"`
	AvailableSpace   int64          `protobuf:"varint,5,opt,name=availableSpace,proto3" json:"availableSpace,omitempty"`
	Startdate        string         `protobuf:"bytes,6,opt,name=startdate,proto3" json:"startdate,omitempty"`
	Starttime        string         `protobuf:"bytes,7,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Enddate          string         `protobuf:"bytes,8,opt,name=enddate,proto3" json:"enddate,omitempty"`
	Price            int64          `protobuf:"varint,9,opt,name=price,proto3" json:"price,omitempty"`
	Image            string         `protobuf:"bytes,10,opt,name=image,proto3" json:"image,omitempty"`
	DestinationCount int64          `protobuf:"varint,11,opt,name=destinationCount,proto3" json:"destinationCount,omitempty"`
	Destination      string         `protobuf:"bytes,12,opt,name=destination,proto3" json:"destination,omitempty"`
	Description      string         `protobuf:"bytes,13,opt,name=Description,proto3" json:"Description,omitempty"`
	MaxCapacity      int64          `protobuf:"varint,14,opt,name=maxCapacity,proto3" json:"maxCapacity,omitempty"`
	CategoryId       int64          `protobuf:"varint,15,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	Category         *Category      `protobuf:"bytes,16,opt,name=category,proto3" json:"category,omitempty"`
	Destinations     []*Destination `protobuf:"bytes,17,rep,name=Destinations,proto3" json:"Destinations,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{6}
}

func (x *Package) GetPackageId() int64 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

func (x *Package) GetPackagename() string {
	if x != nil {
		return x.Packagename
	}
	return ""
}

func (x *Package) GetCoorinatorId() int64 {
	if x != nil {
		return x.CoorinatorId
	}
	return 0
}

func (x *Package) GetStartlocation() string {
	if x != nil {
		return x.Startlocation
	}
	return ""
}

func (x *Package) GetAvailableSpace() int64 {
	if x != nil {
		return x.AvailableSpace
	}
	return 0
}

func (x *Package) GetStartdate() string {
	if x != nil {
		return x.Startdate
	}
	return ""
}

func (x *Package) GetStarttime() string {
	if x != nil {
		return x.Starttime
	}
	return ""
}

func (x *Package) GetEnddate() string {
	if x != nil {
		return x.Enddate
	}
	return ""
}

func (x *Package) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Package) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Package) GetDestinationCount() int64 {
	if x != nil {
		return x.DestinationCount
	}
	return 0
}

func (x *Package) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Package) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Package) GetMaxCapacity() int64 {
	if x != nil {
		return x.MaxCapacity
	}
	return 0
}

func (x *Package) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Package) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Package) GetDestinations() []*Destination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type Destination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationId   int64       `protobuf:"varint,1,opt,name=destinationId,proto3" json:"destinationId,omitempty"`
	DestinationName string      `protobuf:"bytes,2,opt,name=DestinationName,proto3" json:"DestinationName,omitempty"`
	Description     string      `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Minprice        int64       `protobuf:"varint,4,opt,name=minprice,proto3" json:"minprice,omitempty"`
	MaxCapacity     int64       `protobuf:"varint,5,opt,name=maxCapacity,proto3" json:"maxCapacity,omitempty"`
	Image           string      `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Activity        []*Activity `protobuf:"bytes,7,rep,name=activity,proto3" json:"activity,omitempty"`
	PackageID       int64       `protobuf:"varint,8,opt,name=PackageID,proto3" json:"PackageID,omitempty"`
}

func (x *Destination) Reset() {
	*x = Destination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Destination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Destination) ProtoMessage() {}

func (x *Destination) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destination.ProtoReflect.Descriptor instead.
func (*Destination) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{7}
}

func (x *Destination) GetDestinationId() int64 {
	if x != nil {
		return x.DestinationId
	}
	return 0
}

func (x *Destination) GetDestinationName() string {
	if x != nil {
		return x.DestinationName
	}
	return ""
}

func (x *Destination) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Destination) GetMinprice() int64 {
	if x != nil {
		return x.Minprice
	}
	return 0
}

func (x *Destination) GetMaxCapacity() int64 {
	if x != nil {
		return x.MaxCapacity
	}
	return 0
}

func (x *Destination) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Destination) GetActivity() []*Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

func (x *Destination) GetPackageID() int64 {
	if x != nil {
		return x.PackageID
	}
	return 0
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId    int64  `protobuf:"varint,1,opt,name=activityId,proto3" json:"activityId,omitempty"`
	Activityname  string `protobuf:"bytes,2,opt,name=activityname,proto3" json:"activityname,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Location      string `protobuf:"bytes,4,opt,name=Location,proto3" json:"Location,omitempty"`
	ActivityType  string `protobuf:"bytes,5,opt,name=ActivityType,proto3" json:"ActivityType,omitempty"`
	Amount        int64  `protobuf:"varint,6,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Date          string `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Time          string `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	DestinationId int64  `protobuf:"varint,9,opt,name=destinationId,proto3" json:"destinationId,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coordinator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_coordinator_proto_rawDescGZIP(), []int{8}
}

func (x *Activity) GetActivityId() int64 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *Activity) GetActivityname() string {
	if x != nil {
		return x.Activityname
	}
	return ""
}

func (x *Activity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Activity) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Activity) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *Activity) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Activity) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Activity) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Activity) GetDestinationId() int64 {
	if x != nil {
		return x.DestinationId
	}
	return 0
}

var File_coordinator_proto protoreflect.FileDescriptor

var file_coordinator_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x3a, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x61, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x61, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x61, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x42, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x3b, 0x0a, 0x10, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4e,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61,
	0x74, 0x61, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x43, 0x61, 0x74, 0x61, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xce,
	0x04, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f,
	0x6f, 0x72, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x43, 0x6f, 0x6f, 0x72, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x0c, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x9b, 0x02, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x22, 0x96, 0x02,
	0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xe4, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74,
	0x61, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x11, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x16, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x61, 0x63, 0x61, 0x6b, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x74,
	0x61, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x37, 0x0a, 0x1a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x56, 0x69, 0x65, 0x77, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x17, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coordinator_proto_rawDescOnce sync.Once
	file_coordinator_proto_rawDescData = file_coordinator_proto_rawDesc
)

func file_coordinator_proto_rawDescGZIP() []byte {
	file_coordinator_proto_rawDescOnce.Do(func() {
		file_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_coordinator_proto_rawDescData)
	})
	return file_coordinator_proto_rawDescData
}

var file_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_coordinator_proto_goTypes = []interface{}{
	(*Catagories)(nil),       // 0: pb.Catagories
	(*Responce)(nil),         // 1: pb.Responce
	(*View)(nil),             // 2: pb.View
	(*Packages)(nil),         // 3: pb.Packages
	(*PackagesResponce)(nil), // 4: pb.PackagesResponce
	(*Category)(nil),         // 5: pb.Category
	(*Package)(nil),          // 6: pb.Package
	(*Destination)(nil),      // 7: pb.Destination
	(*Activity)(nil),         // 8: pb.Activity
}
var file_coordinator_proto_depIdxs = []int32{
	5,  // 0: pb.Catagories.Catagories:type_name -> pb.Category
	6,  // 1: pb.PackagesResponce.Packages:type_name -> pb.Package
	5,  // 2: pb.Package.category:type_name -> pb.Category
	7,  // 3: pb.Package.Destinations:type_name -> pb.Destination
	8,  // 4: pb.Destination.activity:type_name -> pb.Activity
	5,  // 5: pb.Coordinator.AddCatagory:input_type -> pb.Category
	2,  // 6: pb.Coordinator.AvailablePackages:input_type -> pb.View
	2,  // 7: pb.Coordinator.CoordinatorViewPackage:input_type -> pb.View
	2,  // 8: pb.Coordinator.AdminPacakgeStatus:input_type -> pb.View
	2,  // 9: pb.Coordinator.ViewCatagories:input_type -> pb.View
	2,  // 10: pb.Coordinator.CoordinatorViewDestination:input_type -> pb.View
	2,  // 11: pb.Coordinator.CoordinatorViewActivity:input_type -> pb.View
	1,  // 12: pb.Coordinator.AddCatagory:output_type -> pb.Responce
	4,  // 13: pb.Coordinator.AvailablePackages:output_type -> pb.PackagesResponce
	6,  // 14: pb.Coordinator.CoordinatorViewPackage:output_type -> pb.Package
	1,  // 15: pb.Coordinator.AdminPacakgeStatus:output_type -> pb.Responce
	0,  // 16: pb.Coordinator.ViewCatagories:output_type -> pb.Catagories
	7,  // 17: pb.Coordinator.CoordinatorViewDestination:output_type -> pb.Destination
	8,  // 18: pb.Coordinator.CoordinatorViewActivity:output_type -> pb.Activity
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_coordinator_proto_init() }
func file_coordinator_proto_init() {
	if File_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catagories); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Responce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackagesResponce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Destination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coordinator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coordinator_proto_goTypes,
		DependencyIndexes: file_coordinator_proto_depIdxs,
		MessageInfos:      file_coordinator_proto_msgTypes,
	}.Build()
	File_coordinator_proto = out.File
	file_coordinator_proto_rawDesc = nil
	file_coordinator_proto_goTypes = nil
	file_coordinator_proto_depIdxs = nil
}
