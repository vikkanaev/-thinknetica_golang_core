this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'messages_pb'
require 'messages_services_pb'

def messages
  stub = Protobuf::Signalman::Stub.new('localhost:12345', :this_channel_is_insecure)
  memo = []
  stub.messages(Protobuf::Empty.new) { |r| memo << r }

  memo.each { |e| puts "Id: #{e.id} text: #{e.text}" }
end

messages

# из директории Lesson_14-RPC/ruby
# > ruby ./client.rb
# Id: 1 text: The Go Programming Language
# Id: 2 text: 1984