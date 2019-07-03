# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: helloworld.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='helloworld.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x10helloworld.proto\x12\x02pb\"\x1d\n\x02Hi\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0b\n\x03msg\x18\x02 \x01(\t\"\x1e\n\x03\x42ye\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0b\n\x03msg\x18\x02 \x01(\t2+\n\x0cHelloService\x12\x1b\n\x08SayHello\x12\x06.pb.Hi\x1a\x07.pb.Byeb\x06proto3')
)




_HI = _descriptor.Descriptor(
  name='Hi',
  full_name='pb.Hi',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.Hi.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg', full_name='pb.Hi.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=53,
)


_BYE = _descriptor.Descriptor(
  name='Bye',
  full_name='pb.Bye',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.Bye.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='msg', full_name='pb.Bye.msg', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=55,
  serialized_end=85,
)

DESCRIPTOR.message_types_by_name['Hi'] = _HI
DESCRIPTOR.message_types_by_name['Bye'] = _BYE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Hi = _reflection.GeneratedProtocolMessageType('Hi', (_message.Message,), {
  'DESCRIPTOR' : _HI,
  '__module__' : 'helloworld_pb2'
  # @@protoc_insertion_point(class_scope:pb.Hi)
  })
_sym_db.RegisterMessage(Hi)

Bye = _reflection.GeneratedProtocolMessageType('Bye', (_message.Message,), {
  'DESCRIPTOR' : _BYE,
  '__module__' : 'helloworld_pb2'
  # @@protoc_insertion_point(class_scope:pb.Bye)
  })
_sym_db.RegisterMessage(Bye)



_HELLOSERVICE = _descriptor.ServiceDescriptor(
  name='HelloService',
  full_name='pb.HelloService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=87,
  serialized_end=130,
  methods=[
  _descriptor.MethodDescriptor(
    name='SayHello',
    full_name='pb.HelloService.SayHello',
    index=0,
    containing_service=None,
    input_type=_HI,
    output_type=_BYE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_HELLOSERVICE)

DESCRIPTOR.services_by_name['HelloService'] = _HELLOSERVICE

# @@protoc_insertion_point(module_scope)