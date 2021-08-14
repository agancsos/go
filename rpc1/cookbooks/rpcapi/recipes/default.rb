# Cookbook:: rpcapi
# Recipe:: default
#
# Copyright:: 2021, Abel Gancsos, All Rights Reserved.

base_path = "#{ENV["HOME"]}/stuff/go/rpc1"
if node["platform_family"] == "windows"
	base_path = "#{ENV["USERPROFILE"]}/stuff/go/rpc1"
end

## Copy source directories to our base directory
remote_directory "#{base_path}/src" do
  source "src"
  files_owner "#{ENV["USER"]}"
  files_mode "0777"   ## 777 to avoid access denied issues
  action :create
  recursive true
  overwrite true
  path "#{base_path}/src"
end

## Compile source and run Unit Tests
execute "Compile and run Unit Tests" do
    command "python #{base_path}/src/compile.py"  ## Python3 would provide expected results, but Python should also be acceptable
end

## Cleanup directories
directory "#{base_path}" do
	recursive true
	action :delete
end

## Footer messages
## Traditionally, these would be placed in the header, but due to the logging of the execution, it may be difficult to debug.
log "".ljust(80, "#") 
log "# User              : #{ENV["USER"]}".ljust(79, " ") + "#"
log "# Platform          : #{node["platform"]}".ljust(79, " ") + "#"
log "# BasePath          : #{base_path}".ljust(79, " ") + "#"
log "".ljust(80, "#")
log ""

