/* Added by Yifei */
#ifndef FILE_CONTROL_PLANE_SERVICE_REQUEST_SEEN
#define FILE_CONTROL_PLANE_SERVICE_REQUEST_SEEN

#include <stdint.h>

#include "lte/gateway/c/core/oai/tasks/nas/ies/SecurityHeaderType.h"
#include "lte/gateway/c/core/oai/tasks/nas/ies/MessageType.h"
#include "lte/gateway/c/core/oai/tasks/nas/ies/EsmMessageContainer.h"
#include "lte/gateway/c/core/oai/tasks/nas/ies/EpsBearerContextStatus.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_23.003.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_24.007.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_24.008.h"

typedef enum control_plane_service_request_iei_tag {
  CONTROL_PLANE_SERVICE_REQUEST_EPS_BEARER_CONTEXT_STATUS_IEI = 0x57, /* 0x57 = 87 */
  CONTROL_PLANE_SERVICE_REQUEST_ESM_MESSAGE_CONTAINER_IEI = 0x78, /* 0x78 = 120 */
} control_plane_service_request_iei;

typedef struct control_plane_service_request_msg_tag {
  /* Mandatory fields */
  eps_protocol_discriminator_t protocoldiscriminator : 4;
  security_header_type_t securityheadertype : 4;
  message_type_t messagetype;
  uint8_t headeroctet;
  EsmMessageContainer esmmessagecontainer;
  eps_bearer_context_status_t epsbearercontextstatus;
} control_plane_service_request_msg;

int decode_control_plane_service_request(
    control_plane_service_request_msg* controlplaneservicerequest, uint8_t* buffer, uint32_t len);

#endif