# frozen_string_literal: true

require 'ffi'

module Kitchen
  module Driver
    module GoSample
      extend FFI::Library
      ffi_lib './lib/kitchen/driver/go_sdk/azure_go.so'

      attach_function :create_resource_group, %i[string], :int
    end
  end
end
