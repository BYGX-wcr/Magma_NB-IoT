/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package analytics

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"magma/orc8r/cloud/go/services/analytics/calculations"
)

func TestRawMetricsCalculation(t *testing.T) {
	analytics := &calculations.AnalyticsConfig{
		Metrics: map[string]calculations.MetricConfig{
			"test_metric_1": {
				Expr: "test_metric_expr",
			},
			"test_metric_2": {
				Register: false,
			},
			"test_metric_3": {
				Expr: "test_metric_3_expr",
			},
		},
	}

	calcs := GetRawMetricsCalculations(analytics)
	assert.Equal(t, len(calcs), 2)
	rawMetricNames := []string{"test_metric_1", "test_metric_3"}
	calcMetricNames := []string{}
	for _, calc := range calcs {
		rawMetricCalc := calc.(*calculations.RawMetricsCalculation)
		calcMetricNames = append(calcMetricNames, rawMetricCalc.CalculationParams.Name)
	}
	sort.Strings(calcMetricNames)
	assert.Equal(t, reflect.DeepEqual(rawMetricNames, calcMetricNames), true)
}
