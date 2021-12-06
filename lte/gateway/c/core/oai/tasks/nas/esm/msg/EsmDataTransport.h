/* Added by Yifei */

#ifndef ESM_DATA_TRANSPORT_H_
#define ESM_DATA_TRANSPORT_H_

#include <stdint.h>

#include "lte/gateway/c/core/oai/tasks/nas/ies/MessageType.h"
#include "lte/gateway/c/core/oai/tasks/nas/emm/emm_data.h"
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
    
int encode_esm_data_transport(
    esm_data_transport_msg* esm_data_transport, uint8_t* buffer,
    uint32_t len);

status_code_e esm_proc_data_transport(
    const bool is_standalone, emm_context_t* const emm_context_p,
    const ebi_t ebi, STOLEN_REF bstring* msg, const bool ue_triggered);

#endif
