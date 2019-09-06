# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto

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
  name='service.proto',
  package='twirk.internal.twirktest',
  syntax='proto3',
  serialized_pb=_b('\n\rservice.proto\x12\x18twirk.internal.twirktest\"0\n\x03Hat\x12\x0c\n\x04size\x18\x01 \x01(\x05\x12\r\n\x05\x63olor\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"\x16\n\x04Size\x12\x0e\n\x06inches\x18\x01 \x01(\x05\x32W\n\x0bHaberdasher\x12H\n\x07MakeHat\x12\x1e.twirk.internal.twirktest.Size\x1a\x1d.twirk.internal.twirktest.HatB\x0bZ\ttwirktestb\x06proto3')
)




_HAT = _descriptor.Descriptor(
  name='Hat',
  full_name='twirk.internal.twirktest.Hat',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='size', full_name='twirk.internal.twirktest.Hat.size', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='color', full_name='twirk.internal.twirktest.Hat.color', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='twirk.internal.twirktest.Hat.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=43,
  serialized_end=91,
)


_SIZE = _descriptor.Descriptor(
  name='Size',
  full_name='twirk.internal.twirktest.Size',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='inches', full_name='twirk.internal.twirktest.Size.inches', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=93,
  serialized_end=115,
)

DESCRIPTOR.message_types_by_name['Hat'] = _HAT
DESCRIPTOR.message_types_by_name['Size'] = _SIZE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Hat = _reflection.GeneratedProtocolMessageType('Hat', (_message.Message,), dict(
  DESCRIPTOR = _HAT,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:twirk.internal.twirktest.Hat)
  ))
_sym_db.RegisterMessage(Hat)

Size = _reflection.GeneratedProtocolMessageType('Size', (_message.Message,), dict(
  DESCRIPTOR = _SIZE,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:twirk.internal.twirktest.Size)
  ))
_sym_db.RegisterMessage(Size)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\ttwirktest'))

_HABERDASHER = _descriptor.ServiceDescriptor(
  name='Haberdasher',
  full_name='twirk.internal.twirktest.Haberdasher',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=117,
  serialized_end=204,
  methods=[
  _descriptor.MethodDescriptor(
    name='MakeHat',
    full_name='twirk.internal.twirktest.Haberdasher.MakeHat',
    index=0,
    containing_service=None,
    input_type=_SIZE,
    output_type=_HAT,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_HABERDASHER)

DESCRIPTOR.services_by_name['Haberdasher'] = _HABERDASHER

# @@protoc_insertion_point(module_scope)
