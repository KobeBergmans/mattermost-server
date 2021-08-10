// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/v6/store/searchtest"
	"github.com/mattermost/mattermost-server/v6/store/storetest"
)

func TestUserStore(t *testing.T) {
	StoreTestWithSQLStore(t, storetest.TestUserStore)
}

func TestSearchUserStore(t *testing.T) {
	StoreTestWithSearchTestEngine(t, searchtest.TestSearchUserStore)
}
