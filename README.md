# Ansible Mixin for Porter

This is a mixin for Porter that provides the Ansible CLIs.

## Mixin Configuration

You can simply declaire the plugin and install the latest version of ansible

**Use the vanilla ansible command**
```yaml
mixins:
- ansible
```

You also have the option to add some custom installations setup : 

**Install additional extensions**

```yaml
mixins:
- ansible:
   ## If you want to install a specific Ansible version.
   ## You need to define the version specifier for pip. See https://www.python.org/dev/peps/pep-0440/#version-specifiers.
   clientVersion: "==2.10" 
   otherPipDependencies: ## Python modules to install with Ansible
      - jmespath 
   requirementsFiles: ## Possibly your python dependencies can be in some requirements files
      - requirements.txt
   constraintsFiles: ## Possibly you have some constraints to add during your pip installation
      - constraints.txt
```

## Mixin Syntax

This mixin has several subcommands **adhoc**, **playbook** and **galaxy**.

All the following commands are example, but there is no prerequisites on the number or name of **arguments** or **flags** so, any ansible command is possible with this plugin.

### Adhoc

This command runs an ansible adhoc command.

```yaml
install:
  - ansible:
      description: "Run our ansible adhoc command with all fields"
      adhoc: 
        arguments:
          - "host_pattern*"
        flags:
          inventory: "/etc/ansible/hosts"
          module-name: "debug"
          args: "var=variable_name"
```

Would resolve in the following command : 

```console
ansible --inventory /etc/ansible/hosts --module-name debug --args var=variable_name host_pattern* 
```

### Playbook

This command runs an ansible playbook that you would have put in your bundle.

```yaml
install:
  - ansible:
      description: "Run our ansible playbook command with arbitrary fields"
      playbook: 
        arguments:
          - "playbook.yml"
        flags:
          inventory: "my_inventory"
```

Would resolve in the following command : 

```console
ansible-playbook --inventory my_inventory playbook.yml
```

### Galaxy

This command runs an ansible galaxy command to install your roles and/or collections. 

```yaml
install:
  - ansible:
      description: "Run our ansible galaxy command with arbitrary fields"
      galaxy: 
        arguments:
          - "collection"
          - "install"
          - "community.general"
        flags:
          collection-path: "collections"
          force: ""
```

Would resolve in the following command : 

```console
ansible-galaxy collection install community.general --collections-path collections --force
```