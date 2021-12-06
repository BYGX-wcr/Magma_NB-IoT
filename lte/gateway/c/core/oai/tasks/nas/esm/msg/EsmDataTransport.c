/* Added by Yifei */

#include <stdint.h>

#include "lte/gateway/c/core/oai/common/log.h"
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#include "lte/gateway/c/core/oai/common/common_defs.h"
#include "lte/gateway/c/core/oai/common/common_types.h"
#include "lte/gateway/c/core/oai/lib/bstr/bstrlib.h"
#include "lte/gateway/c/core/oai/tasks/mme_app/mme_app_defs.h"
#include "lte/gateway/c/core/oai/tasks/nas/esm/msg/EsmDataTransport.h"
#include "lte/gateway/c/core/oai/tasks/nas/api/network/nas_message.h"

int decode_esm_data_transport(
    esm_data_transport_msg* esm_data_transport, uint8_t* buffer,
    uint32_t len) {
  OAILOG_FUNC_IN(LOG_NAS_ESM);

  int decoded = 0;
  int decode_result;
  uint16_t ielen;

  DECODE_LENGTH_U16(buffer, ielen, decoded);
  CHECK_LENGTH_DECODER(len - decoded, ielen);
  esm_data_transport->userdatalen = ielen;

  if ((decode_result = decode_bstring(
           &esm_data_transport->userdata, ielen, buffer + decoded, len - decoded)) < 0) {
    OAILOG_FUNC_RETURN(LOG_NAS_ESM, decode_result);
  } else {
    decoded += decode_result;
  }

  OAILOG_FUNC_RETURN(LOG_NAS_ESM, decoded);
}

int encode_esm_data_transport(
    esm_data_transport_msg* esm_data_transport, uint8_t* buffer,
    uint32_t len) {
  OAILOG_FUNC_IN(LOG_NAS_ESM);
  OAILOG_DEBUG(LOG_NAS_ESM,
          "Encoding esm_data_transport: buffer limit-%d, userdata length:%d", len, (uint32_t)esm_data_transport->userdatalen);
  int encoded       = 0;
  int encode_result = 0;

  // reserve the space for length field
  uint16_t* lenPtr  = (uint16_t*)buffer;
  encoded += 2;

  memcpy((void*)(buffer + encoded), (void*)(esm_data_transport->userdata->data), (uint32_t)esm_data_transport->userdatalen);
  encoded += esm_data_transport->userdatalen;
  OAILOG_INFO(LOG_NAS_ESM, "encoded bytes: %d\n", encoded);

  *lenPtr = htons (esm_data_transport->userdatalen);
  OAILOG_FUNC_RETURN(LOG_NAS_ESM, encoded);
}

status_code_e esm_proc_data_transport(
    const bool is_standalone, emm_context_t* const emm_context_p,
    const ebi_t ebi, bstring* msg, const bool ue_triggered) {
  OAILOG_FUNC_IN(LOG_NAS_ESM);
  int rc = 0;
  mme_ue_s1ap_id_t ue_id =
      PARENT_STRUCT(emm_context_p, struct ue_mm_context_s, emm_context)
          ->mme_ue_s1ap_id;

  OAILOG_INFO(
      LOG_NAS_ESM,
      "ESM-PROC - Process ESM data transport message"
      "(ue_id=" MME_UE_S1AP_ID_FMT ", ebi=%d)\n",
      ue_id, ebi);

  if (!ue_triggered) {
    OAILOG_INFO(
      LOG_NAS_ESM,
      "ESM - "
      "Sending ESM data transport reply to (ue_id=" MME_UE_S1AP_ID_FMT ", ebi=%d)\n",
      ue_id, ebi);
    nas_error_code_t status_code = AS_SUCCESS;
    rc = mme_app_handle_nas_dl_req(ue_id, *msg, status_code);
    
  }
  
  OAILOG_FUNC_RETURN(LOG_NAS_ESM, rc);
}