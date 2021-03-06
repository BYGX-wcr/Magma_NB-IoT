"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""


import unittest

import s1ap_types
import s1ap_wrapper


class TestSMCTimer3460Expiry(unittest.TestCase):
    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_smc_timer_3460_expiry(self):
        """ T3460 expiry for Security Mode Complete"""
        # Ground work.
        self._s1ap_wrapper.configUEDevice(1)
        req = self._s1ap_wrapper.ue_req
        maxNasMsgRetransmission = 5

        # Trigger Attach Request
        attach_req = s1ap_types.ueAttachRequest_t()
        sec_ctxt = s1ap_types.TFW_CREATE_NEW_SECURITY_CONTEXT
        id_type = s1ap_types.TFW_MID_TYPE_IMSI
        eps_type = s1ap_types.TFW_EPS_ATTACH_TYPE_EPS_ATTACH
        attach_req.ue_Id = req.ue_id
        attach_req.mIdType = id_type
        attach_req.epsAttachType = eps_type
        attach_req.useOldSecCtxt = sec_ctxt

        print("*** Triggering Attach Request ***")
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_ATTACH_REQUEST, attach_req,
        )

        # Waiting for Authentication Request
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_AUTH_REQ_IND.value,
        )
        print("*** Authentication Request Message Received ***")

        # Trigger Authentication Response
        auth_res = s1ap_types.ueAuthResp_t()
        auth_res.ue_Id = req.ue_id
        sqnRecvd = s1ap_types.ueSqnRcvd_t()
        sqnRecvd.pres = 0
        auth_res.sqnRcvd = sqnRecvd
        print("*** Sending Authentication Response Message ***")
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_AUTH_RESP, auth_res,
        )

        # Waiting for Security mode command
        # Wait for last Timer T3460 Expiry
        for i in range(maxNasMsgRetransmission):
            print(
                "*** Waiting for Security Mode Command Message (",
                str(i + 1),
                ") ***",
            )
            response = self._s1ap_wrapper.s1_util.get_response()
            self.assertEqual(
                response.msg_type, s1ap_types.tfwCmd.UE_SEC_MOD_CMD_IND.value,
            )
            print(
                "*** Security Mode Command Message Received (",
                str(i + 1),
                ") ***",
            )

        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_CTX_REL_IND.value,
        )
        print("*** Received UE Context Release Indication ***")


if __name__ == "__main__":
    unittest.main()
