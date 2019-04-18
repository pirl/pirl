Pod::Spec.new do |spec|
  spec.name         = 'pirl'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://git.pirl.io/community/pirl'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Ethereum Client'
  spec.source       = { :git => 'https://git.pirl.io/community/pirl.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/pirl.framework'

	spec.prepare_command = <<-CMD
    curl https://pirlstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/pirl.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
