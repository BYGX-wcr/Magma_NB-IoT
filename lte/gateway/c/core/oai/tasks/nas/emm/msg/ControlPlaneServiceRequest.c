/* Added by Yifei */
#include <stdint.h>

#include "lte/gateway/c/core/oai/common/log.h"
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#include "lte/gateway/c/core/oai/common/common_defs.h"
#include "lte/gateway/c/core/oai/tasks/nas/emm/msg/ControlPlaneServiceRequest.h"

int decode_control_plane_service_request(
    control_plane_service_request_msg* control_plane_service_request, uint8_t* buffer, uint32_t len) {
  uint32_t decoded   = 0;
  int decoded_result = 0;

  OAILOG_FUNC_IN(LOG_NAS_EMM);
  
  DECODE_U8(buffer, control_plane_service_request->headeroctet, decoded);
  if ((decoded_result = decode_esm_message_container(
    &control_plane_service_request->esmmessagecontainer, CONTROL_PLANE_SERVICE_REQUEST_ESM_MESSAGE_CONTAINER_IEI, buffer + decoded,
    len - decoded)) < 0)
    return decoded_result;
  else
    decoded += decoded_result;
  
  if ((decoded_result = decode_eps_bearer_context_status(
                 &control_plane_service_request->epsbearercontextstatus,
                 CONTROL_PLANE_SERVICE_REQUEST_EPS_BEARER_CONTEXT_STATUS_IEI,
                 buffer + decoded, len - decoded)) <= 0)
    return decoded_result;
  else
    decoded += decoded_result;

  OAILOG_FUNC_RETURN(LOG_NAS_EMM, decoded);
}