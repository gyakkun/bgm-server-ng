// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package ctrl

import (
	"context"

	"github.com/bangumi/server/internal/episode"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/errgo"
	"github.com/bangumi/server/internal/pkg/null"
)

func (ctl Ctrl) ListEpisode(
	ctx context.Context,
	subjectID model.SubjectID,
	epType *episode.Type,
	limit, offset int,
) ([]episode.Episode, int64, error) {
	count, err := ctl.CountEpisode(ctx, subjectID, epType)
	if err != nil {
		return nil, 0, err
	}

	if count == 0 {
		return []episode.Episode{}, 0, nil
	}

	if int64(offset) > count {
		return []episode.Episode{}, count, ErrOffsetTooBig
	}

	var episodes []episode.Episode
	episodes, err = ctl.episode.List(ctx, subjectID, episode.Filter{Type: null.NewFromPtr(epType)}, limit, offset)
	if err != nil {
		return nil, 0, errgo.Wrap(err, "episode.List")
	}
	return episodes, count, nil
}