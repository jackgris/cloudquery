package monitor

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createActivityLogAlerts(router *mux.Router) error {
	var item armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyString := ""
	item.NextLink = &emptyString

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/activityLogAlerts", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestActivityLogAlerts(t *testing.T) {
	client.MockTestHelper(t, ActivityLogAlerts(), createActivityLogAlerts)
}
