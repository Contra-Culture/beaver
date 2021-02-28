# beaver the builder

``beaver`` is a simple build tool that highly rely on KISS and CoC principles.

---

## Philosophy

Beaver is a golang build tool, that relies on the following principles:

### **KISS**
Keep it simple, stupid: ``beaver`` is only responsible for replacing abstract (interface) dependencies with concrete implementations before compilation.
### **CoC**
Conventions over Configurations: ``beaver`` uses source-files naming and nesting conventions instead of complex configuration files.

---

## Concepts

### **Source Tree**
Source tree is a directory/file tree with all the sources of the project you want to build. ``beaver`` reads source tree, copies sources into separate directory and places correct dependency implementations together with them.

### **Interface**

### **Implementations**
Implementation is a package that implements a dependency interface.


### **Dumb Implementation**


---

## Installation

---
## Usage

### Conventions
#### **Dumb implementations**
The dumb implementation is the only required implementation. It works as an example and a placeholder for more userful implementations. The dumb implementation should be in the ``<interface-name>`` folder on the same level as its client (higher-level) code.

#### **Named implementations**
All the named implementations should be placed within ``__<interface-name>`` directories, within the directories that represent their names: ``<interface-name>/<implementation-name>``.

#### **Default dumb implementation**
Because the dumb implementation is the only required implementation, it is a good default default (yes, 2 times!) value that will be used if ``build-defaults`` config has no other implementation defined and no implementation provided by ``-impl`` parameter (see: [Command Line Interface](#command-line-interface) section).

### Configuration

#### **build-defaults config file**

``build-defaults`` file defines build defaults in the same format as CLI attributes:

```
<interface-name-1> : <implementation-name>
<interface-name-2> : <implementation-name>
```

#### **\<name>-build-settings config file**
``<name>-build-settings`` file defines interface implementations for named builds, like: ``test``, ``production``, etc., using the folowing format:

```
<interface-name-1> : <implementation-name>
<interface-name-2> : <implementation-name>
```

### Command Line Interface

#### **beaver \<path>**
``beaver <path>`` builds from sources from ``<path>`` directory  with defaults from: ``build-defaults`` files. The build will be named ``default-<timestamp>``.

#### **beaver \<path> --**

Builds from sources from ``<path>`` directory with given implementations (see: [Configuration](#configuration) section) and names them by the following schema:

```<interface-name-1>(<implementation-name>)[_<interface-name-2>(<implementation-name>)[..._<interface-name-<n>>(<implementation-name>)]]-<timestamp>```

#### **cartesian builds**

You can make cartesian builds (multi-builds) with ``--`` parameter list. For this purpose you only need to provide a list of interface implementations instead of a single one. For example:

``beaver <path> -- publications/repository : mongodb | postgresql | inmemory -- transport : restful | soap``

... will make the following builds:

```
__builds/publications--repository(mongodb)__transport(restful)-<timestamp>
__builds/publications--repository(mongodb)__transport(soap)-<timestamp>
__builds/publications--repository(postgresql)__transport(restful)-<timestamp>
__builds/publications--repository(postgresql)__transport(soap)-<timestamp>
__builds/publications--repository(inmemory)__transport(restful)-<timestamp>
__builds/publications--repository(inmemory)__transport(soap)-<timestamp>
```

#### **beaver \<path> \<name>**

Builds from sources from ``<path>`` directory, using ``<name>-build-settings`` config file and names the build by the ``<name>-<timestamp>`` schema.

---

## Development

---
