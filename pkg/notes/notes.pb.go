// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notes.proto

package notes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Query struct {
	IDs                  []string             `protobuf:"bytes,1,rep,name=IDs,proto3" json:"IDs,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Team                 string               `protobuf:"bytes,3,opt,name=team,proto3" json:"team,omitempty"`
	Authors              []string             `protobuf:"bytes,4,rep,name=authors,proto3" json:"authors,omitempty"`
	Todate               *timestamp.Timestamp `protobuf:"bytes,5,opt,name=todate,proto3" json:"todate,omitempty"`
	Fromdate             *timestamp.Timestamp `protobuf:"bytes,6,opt,name=fromdate,proto3" json:"fromdate,omitempty"`
	Tags                 []string             `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetIDs() []string {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *Query) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Query) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *Query) GetAuthors() []string {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Query) GetTodate() *timestamp.Timestamp {
	if m != nil {
		return m.Todate
	}
	return nil
}

func (m *Query) GetFromdate() *timestamp.Timestamp {
	if m != nil {
		return m.Fromdate
	}
	return nil
}

func (m *Query) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type NoteList struct {
	Notes                []*Note  `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoteList) Reset()         { *m = NoteList{} }
func (m *NoteList) String() string { return proto.CompactTextString(m) }
func (*NoteList) ProtoMessage()    {}
func (*NoteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{1}
}

func (m *NoteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoteList.Unmarshal(m, b)
}
func (m *NoteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoteList.Marshal(b, m, deterministic)
}
func (m *NoteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoteList.Merge(m, src)
}
func (m *NoteList) XXX_Size() int {
	return xxx_messageInfo_NoteList.Size(m)
}
func (m *NoteList) XXX_DiscardUnknown() {
	xxx_messageInfo_NoteList.DiscardUnknown(m)
}

var xxx_messageInfo_NoteList proto.InternalMessageInfo

func (m *NoteList) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

type Tag struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Color                string   `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{2}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Tag) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

type Note struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Team                 string               `protobuf:"bytes,4,opt,name=team,proto3" json:"team,omitempty"`
	Body                 string               `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Type                 string               `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,7,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	DateModified         *timestamp.Timestamp `protobuf:"bytes,8,opt,name=dateModified,proto3" json:"dateModified,omitempty"`
	Link                 string               `protobuf:"bytes,9,opt,name=link,proto3" json:"link,omitempty"`
	Image                []byte               `protobuf:"bytes,10,opt,name=image,proto3" json:"image,omitempty"`
	Tags                 []*Tag               `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	Score                float32              `protobuf:"fixed32,12,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{3}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Note) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *Note) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Note) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Note) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *Note) GetDateModified() *timestamp.Timestamp {
	if m != nil {
		return m.DateModified
	}
	return nil
}

func (m *Note) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Note) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Note) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Note) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Query)(nil), "notes.Query")
	proto.RegisterType((*NoteList)(nil), "notes.NoteList")
	proto.RegisterType((*Tag)(nil), "notes.Tag")
	proto.RegisterType((*Note)(nil), "notes.Note")
	proto.RegisterType((*Empty)(nil), "notes.Empty")
}

func init() { proto.RegisterFile("notes.proto", fileDescriptor_ffefd935cd6c4a4a) }

var fileDescriptor_ffefd935cd6c4a4a = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0x49, 0x26, 0x99, 0x3f, 0x95, 0x41, 0xa5, 0x11, 0x69, 0xe6, 0xa0, 0x31, 0x07, 0x89,
	0x07, 0x33, 0x30, 0x82, 0x27, 0xf1, 0xe2, 0x8a, 0x08, 0x2a, 0x18, 0xc6, 0x07, 0xe8, 0xd9, 0xd4,
	0xc4, 0xc6, 0x64, 0x3b, 0x24, 0x35, 0xb0, 0x79, 0x25, 0xdf, 0x49, 0x9f, 0x45, 0xba, 0x3a, 0x59,
	0xe2, 0x41, 0xb2, 0xb7, 0xfa, 0xba, 0xbf, 0x6a, 0xbe, 0xfa, 0x55, 0x02, 0xd1, 0x8d, 0x21, 0xec,
	0xb2, 0xa6, 0x35, 0x64, 0x44, 0xc8, 0x62, 0xf7, 0xac, 0x34, 0xa6, 0xac, 0x70, 0xcf, 0x87, 0xa7,
	0xcb, 0x79, 0x4f, 0xba, 0xc6, 0x8e, 0x54, 0xdd, 0x38, 0x5f, 0xf2, 0xdb, 0x83, 0xf0, 0xdb, 0x05,
	0xdb, 0x5e, 0x3c, 0x82, 0xc5, 0xa7, 0xab, 0x4e, 0x7a, 0xf1, 0x22, 0xdd, 0xe4, 0xb6, 0x14, 0x8f,
	0x21, 0x24, 0x4d, 0x15, 0x4a, 0x3f, 0xf6, 0xd2, 0x4d, 0xee, 0x84, 0x10, 0x10, 0x10, 0xaa, 0x5a,
	0x2e, 0xf8, 0x90, 0x6b, 0x21, 0x61, 0xa5, 0x2e, 0xf4, 0xc3, 0xb4, 0x9d, 0x0c, 0xb8, 0x7f, 0x94,
	0xe2, 0x00, 0x4b, 0x32, 0x85, 0x22, 0x94, 0x61, 0xec, 0xa5, 0xd1, 0x61, 0x97, 0xb9, 0x44, 0xd9,
	0x98, 0x28, 0x3b, 0x8e, 0x89, 0xf2, 0xc1, 0x29, 0xde, 0xc0, 0xfa, 0xdc, 0x9a, 0x9a, 0xbb, 0x96,
	0xb3, 0x5d, 0x77, 0x5e, 0x4e, 0xa6, 0xca, 0x4e, 0xae, 0x38, 0x02, 0xd7, 0xc9, 0x2b, 0x58, 0x7f,
	0x35, 0x84, 0x9f, 0x75, 0x47, 0xe2, 0x39, 0x38, 0x2a, 0x3c, 0x63, 0x74, 0x88, 0x32, 0x07, 0xcc,
	0xde, 0xe7, 0xee, 0x26, 0xd9, 0xc3, 0xe2, 0xa8, 0x4a, 0x37, 0xe3, 0x2d, 0x49, 0x6f, 0x9c, 0xf1,
	0x96, 0x2c, 0x8d, 0x6b, 0x53, 0x99, 0x76, 0xa4, 0xc1, 0x22, 0xf9, 0xe3, 0x43, 0x60, 0x1f, 0x10,
	0x0f, 0xc0, 0xd7, 0xc5, 0xd0, 0xe0, 0xeb, 0xe2, 0x3f, 0xf0, 0x9e, 0xc0, 0xd2, 0x91, 0x19, 0xf0,
	0x0d, 0xea, 0x0e, 0x6a, 0x30, 0x81, 0x2a, 0x20, 0x38, 0x99, 0xa2, 0x67, 0x70, 0x9b, 0x9c, 0x6b,
	0xf6, 0xf5, 0x8d, 0xc3, 0x62, 0x7d, 0x7d, 0x83, 0xe2, 0x2d, 0x44, 0x76, 0xfc, 0xf7, 0x2d, 0x2a,
	0xc2, 0x42, 0xae, 0x66, 0x89, 0x4d, 0xed, 0xe2, 0x1d, 0x6c, 0xad, 0xfc, 0x62, 0x0a, 0x7d, 0xd6,
	0x58, 0xc8, 0xf5, 0x6c, 0xfb, 0x3f, 0x7e, 0x9b, 0xa8, 0xd2, 0x37, 0x3f, 0xe5, 0xc6, 0x25, 0xb2,
	0xb5, 0x9d, 0x5d, 0xd7, 0xaa, 0x44, 0x09, 0xb1, 0x97, 0x6e, 0x73, 0x27, 0xc4, 0xd3, 0x61, 0x3d,
	0x11, 0xd3, 0x87, 0x81, 0xfe, 0x51, 0x95, 0x6e, 0x55, 0xb6, 0xab, 0xbb, 0x36, 0x2d, 0xca, 0x6d,
	0xec, 0xa5, 0x7e, 0xee, 0x44, 0xb2, 0x82, 0xf0, 0x43, 0xdd, 0x50, 0x7f, 0xf8, 0xe5, 0x41, 0x68,
	0x49, 0x77, 0xe2, 0x05, 0x80, 0x4b, 0xcf, 0xe0, 0xa7, 0x6b, 0xdc, 0x4d, 0x85, 0xf5, 0x5d, 0x61,
	0x85, 0xf7, 0xf1, 0x7d, 0x6f, 0x8a, 0xf9, 0xf7, 0x5e, 0xc2, 0xfa, 0x23, 0x92, 0xcb, 0xb0, 0x1d,
	0x2e, 0xf8, 0xdf, 0xd9, 0x3d, 0x9c, 0xd8, 0xec, 0xa7, 0x76, 0x5a, 0x32, 0xb7, 0xd7, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x46, 0xf6, 0x34, 0x2d, 0x94, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotesClient is the client API for Notes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotesClient interface {
	CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error)
	DeleteNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error)
	UpdateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error)
	GetNotes(ctx context.Context, in *Query, opts ...grpc.CallOption) (*NoteList, error)
}

type notesClient struct {
	cc *grpc.ClientConn
}

func NewNotesClient(cc *grpc.ClientConn) NotesClient {
	return &notesClient{cc}
}

func (c *notesClient) CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) DeleteNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) UpdateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) GetNotes(ctx context.Context, in *Query, opts ...grpc.CallOption) (*NoteList, error) {
	out := new(NoteList)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServer is the server API for Notes service.
type NotesServer interface {
	CreateNote(context.Context, *Note) (*Note, error)
	DeleteNote(context.Context, *Note) (*Note, error)
	UpdateNote(context.Context, *Note) (*Note, error)
	GetNotes(context.Context, *Query) (*NoteList, error)
}

func RegisterNotesServer(s *grpc.Server, srv NotesServer) {
	s.RegisterService(&_Notes_serviceDesc, srv)
}

func _Notes_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).CreateNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).DeleteNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).UpdateNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_GetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetNotes(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notes.Notes",
	HandlerType: (*NotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _Notes_CreateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _Notes_DeleteNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _Notes_UpdateNote_Handler,
		},
		{
			MethodName: "GetNotes",
			Handler:    _Notes_GetNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes.proto",
}
