# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: empty_service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='empty_service.proto',
  package='twirk.internal.twirktest.emptyservice',
  syntax='proto3',
  serialized_pb=_b('\n\x13\x65mpty_service.proto\x12%twirk.internal.twirktest.emptyservice2\x07\n\x05\x45mptyB\x0fZ\rempty_serviceb\x06proto3')
)



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\rempty_service'))

_EMPTY = _descriptor.ServiceDescriptor(
  name='Empty',
  full_name='twirk.internal.twirktest.emptyservice.Empty',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=62,
  serialized_end=69,
  methods=[
])
_sym_db.RegisterServiceDescriptor(_EMPTY)

DESCRIPTOR.services_by_name['Empty'] = _EMPTY

# @@protoc_insertion_point(module_scope)
