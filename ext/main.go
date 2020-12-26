package ext

// #cgo LDFLAGS: -lnftables
/*
#include <nftables/libnftables.h>
#include <stdlib.h>

struct nft_ctx* initialize_nft_handle() {
	struct nft_ctx* handle =  nft_ctx_new(NFT_CTX_DEFAULT);

	nft_ctx_output_set_flags(handle,
		NFT_CTX_OUTPUT_HANDLE |
		NFT_CTX_OUTPUT_JSON |
		NFT_CTX_OUTPUT_NUMERIC_PRIO);

	return handle;
}
*/
import "C"
