Name: {{.Name}}
Roles: {{range $i, $role := .Roles}}
	Role_{{$i}} : {{printf "%q" $role }}{{end}}
Created on: {{.CreationTime.Format "2006-01-02T15:04:05Z07:00"}}
Active: {{if .Active}}Y{{else}}N{{end}}
Active: {{ yesno .Active }}
