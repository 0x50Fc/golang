#include "duk_config.h"
#include "duktape.h"
#include "kk.h"

struct kk_ptr * kk_push_ptr(struct duk_hthread *ctx) {
	return (struct kk_ptr *) duk_push_fixed_buffer(ctx,sizeof(struct kk_ptr));
}

struct kk_ptr * kk_to_ptr(struct duk_hthread *ctx,duk_idx_t idx) {
	size_t n=0;
	return (struct kk_ptr *) duk_to_buffer(ctx,idx,&n);
}

static void kk_duk_create_heap_fatal(void *udata, const char *msg) {
	printf("[duktape] [fail] %s\n",msg);
}

struct duk_hthread * kk_duk_create_heap() {
	return duk_create_heap(NULL,NULL,NULL,NULL,kk_duk_create_heap_fatal);
}