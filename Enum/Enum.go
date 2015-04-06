/*
  +----------------------------------------------------------------------+
  | HostsSitter                                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 2.0 of the Apache license,    |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.apache.org/licenses/LICENSE-2.0.html                      |
  | If you did not receive a copy of the Apache2.0 license and are unable|
  | to obtain it through the world-wide-web, please send a note to       |
  | neeke@php.net so we can mail you a copy immediately.                 |
  +----------------------------------------------------------------------+
  | Author: Neeke.Gao  <neeke@php.net>                                   |
  +----------------------------------------------------------------------+
*/
package Enum

const HOST_PATH_LINUX = "/etc/hosts"
const HOST_PATH_WINDOWS = "C:\\Windows\\system32\\drivers\\etc\\hosts"

const HOST_PRE_STR = "#Got new hosts from Neeke"
const HOST_POST_STR = "#Got new hosts from Neeke End"

const FLAGS_COMMAND_UPDATE = "update"
const FLAGS_COMMAND_BACKUP = "backup"
const FLAGS_COMMAND_RECOVERY = "recovery"

const FLAGS_SOURCE_ENUM_360kb = 1
