kitex_gen:
	kitex -I idl/service -module tiktok_micro idl/service/user.proto     # execute in the project root directory
	&& kitex -I idl/service -module tiktok_micro idl/service/user.proto
	&& kitex -I idl/service -module tiktok_micro idl/service/user.proto

