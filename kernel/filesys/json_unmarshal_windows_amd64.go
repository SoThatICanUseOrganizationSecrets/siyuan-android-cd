// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package filesys

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/siyuan-note/logging"
)

func unmarshalJSON(data []byte, v interface{}) error {
	now := time.Now()
	defer func() {
		elapsed := time.Since(now)
		logging.LogInfof("[sonic] unmarshalJSON took %s", elapsed)
	}()

	return sonic.Unmarshal(data, v)
}
