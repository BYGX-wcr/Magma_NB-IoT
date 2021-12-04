/* Added by Yifei */

#ifndef ESM_DATA_TRANSPORT_H_
#define ESM_DATA_TRANSPORT_H_

#include <stdint.h>

#include "lte/gateway/c/core/oai/tasks/nas/ies/MessageType.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_23.003.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_24.007.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_24.008.h"

typedef struct esm_data_transport_tag {
  eps_protocol_discriminator_t protocoldiscriminator : 4;
  ebi_t epsbeareridentity : 4;
  pti_t proceduretransactionidentity;
  message_type_t messagetype;
  uint16_t userdatalen;
  bstring userdata;
} esm_data_transport_msg;

int decode_esm_data_transport(
    esm_data_transport_msg* esmdatatransport, uint8_t* buffer,
    uint32_t len);

#endif
