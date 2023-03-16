# frozen_string_literal: true

require 'json'
require "uri"
require "net/http"

module Kitchen
  module Driver
    class RestSample
      URL = "https://management.azure.com/subscriptions/80b824de-ec53-4116-9868-3deeab10b0cd/resourcegroups/%s?api-version=2020-09-01"

      class << self
        def create_resource_group(location, name)
          puts format('Creating a new resource group %s on location: %s', name, location)

          url = URI(format(URL, name))
          https = Net::HTTP.new(url.host, url.port)
          https.use_ssl = true

          request = Net::HTTP::Put.new(url)
          request["Authorization"] = "Bearer #{token}"

          request["Content-Type"] = "application/json"
          request.body = JSON.dump(
            "location": location
          )

          response = https.request(request)
          puts response.read_body

        end

        private

        def token
          @token ||= JSON.parse(`az account get-access-token -o json --resource https://management.core.windows.net/`)["accessToken"]
        end
      end
    end
  end
end
