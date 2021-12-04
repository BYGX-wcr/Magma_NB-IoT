/* Added by Yifei */

#include <stdint.h>

#include "lte/gateway/c/core/oai/common/log.h"
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#include "lte/gateway/c/core/oai/common/common_defs.h"
#include "lte/gateway/c/core/oai/tasks/nas/esm/msg/EsmDataTransport.h"

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
