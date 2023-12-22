//go:build release

package completers

import (
	acpi "github.com/rsteube/carapace-bin/completers_release/acpi_completer/cmd"
	acpid "github.com/rsteube/carapace-bin/completers_release/acpid_completer/cmd"
	adb "github.com/rsteube/carapace-bin/completers_release/adb_completer/cmd"
	age "github.com/rsteube/carapace-bin/completers_release/age_completer/cmd"
	agg "github.com/rsteube/carapace-bin/completers_release/agg_completer/cmd"
	alsamixer "github.com/rsteube/carapace-bin/completers_release/alsamixer_completer/cmd"
	ant "github.com/rsteube/carapace-bin/completers_release/ant_completer/cmd"
	aplay "github.com/rsteube/carapace-bin/completers_release/aplay_completer/cmd"
	apropos "github.com/rsteube/carapace-bin/completers_release/apropos_completer/cmd"
	apt_cache "github.com/rsteube/carapace-bin/completers_release/apt-cache_completer/cmd"
	apt_get "github.com/rsteube/carapace-bin/completers_release/apt-get_completer/cmd"
	ar "github.com/rsteube/carapace-bin/completers_release/ar_completer/cmd"
	arecord "github.com/rsteube/carapace-bin/completers_release/arecord_completer/cmd"
	asciinema "github.com/rsteube/carapace-bin/completers_release/asciinema_completer/cmd"
	autoconf "github.com/rsteube/carapace-bin/completers_release/autoconf_completer/cmd"
	avdmanager "github.com/rsteube/carapace-bin/completers_release/avdmanager_completer/cmd"
	awk "github.com/rsteube/carapace-bin/completers_release/awk_completer/cmd"
	aws "github.com/rsteube/carapace-bin/completers_release/aws_completer/cmd"
	az "github.com/rsteube/carapace-bin/completers_release/az_completer/cmd"
	baobab "github.com/rsteube/carapace-bin/completers_release/baobab_completer/cmd"
	basename "github.com/rsteube/carapace-bin/completers_release/basename_completer/cmd"
	bash_language_server "github.com/rsteube/carapace-bin/completers_release/bash-language-server_completer/cmd"
	bash "github.com/rsteube/carapace-bin/completers_release/bash_completer/cmd"
	bat "github.com/rsteube/carapace-bin/completers_release/bat_completer/cmd"
	batdiff "github.com/rsteube/carapace-bin/completers_release/batdiff_completer/cmd"
	batgrep "github.com/rsteube/carapace-bin/completers_release/batgrep_completer/cmd"
	batman "github.com/rsteube/carapace-bin/completers_release/batman_completer/cmd"
	bats "github.com/rsteube/carapace-bin/completers_release/bats_completer/cmd"
	bc "github.com/rsteube/carapace-bin/completers_release/bc_completer/cmd"
	benthos "github.com/rsteube/carapace-bin/completers_release/benthos_completer/cmd"
	black "github.com/rsteube/carapace-bin/completers_release/black_completer/cmd"
	boundary "github.com/rsteube/carapace-bin/completers_release/boundary_completer/cmd"
	brew "github.com/rsteube/carapace-bin/completers_release/brew_completer/cmd"
	brotli "github.com/rsteube/carapace-bin/completers_release/brotli_completer/cmd"
	bun "github.com/rsteube/carapace-bin/completers_release/bun_completer/cmd"
	bunx "github.com/rsteube/carapace-bin/completers_release/bunx_completer/cmd"
	calibre "github.com/rsteube/carapace-bin/completers_release/calibre_completer/cmd"
	capslock "github.com/rsteube/carapace-bin/completers_release/capslock_completer/cmd"
	carapace "github.com/rsteube/carapace-bin/completers_release/carapace_completer/cmd"
	cargo_clippy "github.com/rsteube/carapace-bin/completers_release/cargo-clippy_completer/cmd"
	cargo_fmt "github.com/rsteube/carapace-bin/completers_release/cargo-fmt_completer/cmd"
	cargo_metadata "github.com/rsteube/carapace-bin/completers_release/cargo-metadata_completer/cmd"
	cargo_rm "github.com/rsteube/carapace-bin/completers_release/cargo-rm_completer/cmd"
	cargo_set_version "github.com/rsteube/carapace-bin/completers_release/cargo-set-version_completer/cmd"
	cargo_upgrade "github.com/rsteube/carapace-bin/completers_release/cargo-upgrade_completer/cmd"
	cargo_watch "github.com/rsteube/carapace-bin/completers_release/cargo-watch_completer/cmd"
	cargo "github.com/rsteube/carapace-bin/completers_release/cargo_completer/cmd"
	cat "github.com/rsteube/carapace-bin/completers_release/cat_completer/cmd"
	cfdisk "github.com/rsteube/carapace-bin/completers_release/cfdisk_completer/cmd"
	charm "github.com/rsteube/carapace-bin/completers_release/charm_completer/cmd"
	cheese "github.com/rsteube/carapace-bin/completers_release/cheese_completer/cmd"
	chgrp "github.com/rsteube/carapace-bin/completers_release/chgrp_completer/cmd"
	chmod "github.com/rsteube/carapace-bin/completers_release/chmod_completer/cmd"
	chown "github.com/rsteube/carapace-bin/completers_release/chown_completer/cmd"
	chpasswd "github.com/rsteube/carapace-bin/completers_release/chpasswd_completer/cmd"
	chroma "github.com/rsteube/carapace-bin/completers_release/chroma_completer/cmd"
	chromium "github.com/rsteube/carapace-bin/completers_release/chromium_completer/cmd"
	chroot "github.com/rsteube/carapace-bin/completers_release/chroot_completer/cmd"
	chsh "github.com/rsteube/carapace-bin/completers_release/chsh_completer/cmd"
	circleci "github.com/rsteube/carapace-bin/completers_release/circleci_completer/cmd"
	cksum "github.com/rsteube/carapace-bin/completers_release/cksum_completer/cmd"
	clamav_config "github.com/rsteube/carapace-bin/completers_release/clamav-config_completer/cmd"
	clamav_milter "github.com/rsteube/carapace-bin/completers_release/clamav-milter_completer/cmd"
	clambc "github.com/rsteube/carapace-bin/completers_release/clambc_completer/cmd"
	clamconf "github.com/rsteube/carapace-bin/completers_release/clamconf_completer/cmd"
	clamd "github.com/rsteube/carapace-bin/completers_release/clamd_completer/cmd"
	clamdscan "github.com/rsteube/carapace-bin/completers_release/clamdscan_completer/cmd"
	clamdtop "github.com/rsteube/carapace-bin/completers_release/clamdtop_completer/cmd"
	clamonacc "github.com/rsteube/carapace-bin/completers_release/clamonacc_completer/cmd"
	clamscan "github.com/rsteube/carapace-bin/completers_release/clamscan_completer/cmd"
	clamsubmit "github.com/rsteube/carapace-bin/completers_release/clamsubmit_completer/cmd"
	cmus "github.com/rsteube/carapace-bin/completers_release/cmus_completer/cmd"
	code_insiders "github.com/rsteube/carapace-bin/completers_release/code-insiders_completer/cmd"
	code "github.com/rsteube/carapace-bin/completers_release/code_completer/cmd"
	codecov "github.com/rsteube/carapace-bin/completers_release/codecov_completer/cmd"
	comm "github.com/rsteube/carapace-bin/completers_release/comm_completer/cmd"
	conda_content_trust "github.com/rsteube/carapace-bin/completers_release/conda-content-trust_completer/cmd"
	conda_env "github.com/rsteube/carapace-bin/completers_release/conda-env_completer/cmd"
	conda "github.com/rsteube/carapace-bin/completers_release/conda_completer/cmd"
	conky "github.com/rsteube/carapace-bin/completers_release/conky_completer/cmd"
	consul "github.com/rsteube/carapace-bin/completers_release/consul_completer/cmd"
	coredumpctl "github.com/rsteube/carapace-bin/completers_release/coredumpctl_completer/cmd"
	cp "github.com/rsteube/carapace-bin/completers_release/cp_completer/cmd"
	csplit "github.com/rsteube/carapace-bin/completers_release/csplit_completer/cmd"
	csview "github.com/rsteube/carapace-bin/completers_release/csview_completer/cmd"
	cura "github.com/rsteube/carapace-bin/completers_release/cura_completer/cmd"
	curl "github.com/rsteube/carapace-bin/completers_release/curl_completer/cmd"
	cut "github.com/rsteube/carapace-bin/completers_release/cut_completer/cmd"
	darktable_cli "github.com/rsteube/carapace-bin/completers_release/darktable-cli_completer/cmd"
	darktable "github.com/rsteube/carapace-bin/completers_release/darktable_completer/cmd"
	dart "github.com/rsteube/carapace-bin/completers_release/dart_completer/cmd"
	date "github.com/rsteube/carapace-bin/completers_release/date_completer/cmd"
	dbt "github.com/rsteube/carapace-bin/completers_release/dbt_completer/cmd"
	dc "github.com/rsteube/carapace-bin/completers_release/dc_completer/cmd"
	dd "github.com/rsteube/carapace-bin/completers_release/dd_completer/cmd"
	deadcode "github.com/rsteube/carapace-bin/completers_release/deadcode_completer/cmd"
	delta "github.com/rsteube/carapace-bin/completers_release/delta_completer/cmd"
	deno "github.com/rsteube/carapace-bin/completers_release/deno_completer/cmd"
	devbox "github.com/rsteube/carapace-bin/completers_release/devbox_completer/cmd"
	df "github.com/rsteube/carapace-bin/completers_release/df_completer/cmd"
	dfc "github.com/rsteube/carapace-bin/completers_release/dfc_completer/cmd"
	dict "github.com/rsteube/carapace-bin/completers_release/dict_completer/cmd"
	diff3 "github.com/rsteube/carapace-bin/completers_release/diff3_completer/cmd"
	diff "github.com/rsteube/carapace-bin/completers_release/diff_completer/cmd"
	dig "github.com/rsteube/carapace-bin/completers_release/dig_completer/cmd"
	dir "github.com/rsteube/carapace-bin/completers_release/dir_completer/cmd"
	dircolors "github.com/rsteube/carapace-bin/completers_release/dircolors_completer/cmd"
	direnv "github.com/rsteube/carapace-bin/completers_release/direnv_completer/cmd"
	dirname "github.com/rsteube/carapace-bin/completers_release/dirname_completer/cmd"
	dive "github.com/rsteube/carapace-bin/completers_release/dive_completer/cmd"
	dlv "github.com/rsteube/carapace-bin/completers_release/dlv_completer/cmd"
	dmenu "github.com/rsteube/carapace-bin/completers_release/dmenu_completer/cmd"
	dmesg "github.com/rsteube/carapace-bin/completers_release/dmesg_completer/cmd"
	dms "github.com/rsteube/carapace-bin/completers_release/dms_completer/cmd"
	dnsmasq "github.com/rsteube/carapace-bin/completers_release/dnsmasq_completer/cmd"
	docker_buildx "github.com/rsteube/carapace-bin/completers_release/docker-buildx_completer/cmd"
	docker_compose "github.com/rsteube/carapace-bin/completers_release/docker-compose_completer/cmd"
	docker_scan "github.com/rsteube/carapace-bin/completers_release/docker-scan_completer/cmd"
	docker "github.com/rsteube/carapace-bin/completers_release/docker_completer/cmd"
	dockerd "github.com/rsteube/carapace-bin/completers_release/dockerd_completer/cmd"
	doctl "github.com/rsteube/carapace-bin/completers_release/doctl_completer/cmd"
	dos2unix "github.com/rsteube/carapace-bin/completers_release/dos2unix_completer/cmd"
	downgrade "github.com/rsteube/carapace-bin/completers_release/downgrade_completer/cmd"
	dpkg "github.com/rsteube/carapace-bin/completers_release/dpkg_completer/cmd"
	du "github.com/rsteube/carapace-bin/completers_release/du_completer/cmd"
	ebook_convert "github.com/rsteube/carapace-bin/completers_release/ebook-convert_completer/cmd"
	egrep "github.com/rsteube/carapace-bin/completers_release/egrep_completer/cmd"
	electron "github.com/rsteube/carapace-bin/completers_release/electron_completer/cmd"
	elvish "github.com/rsteube/carapace-bin/completers_release/elvish_completer/cmd"
	env "github.com/rsteube/carapace-bin/completers_release/env_completer/cmd"
	envsubst "github.com/rsteube/carapace-bin/completers_release/envsubst_completer/cmd"
	exa "github.com/rsteube/carapace-bin/completers_release/exa_completer/cmd"
	expand "github.com/rsteube/carapace-bin/completers_release/expand_completer/cmd"
	expr "github.com/rsteube/carapace-bin/completers_release/expr_completer/cmd"
	eza "github.com/rsteube/carapace-bin/completers_release/eza_completer/cmd"
	faas_cli "github.com/rsteube/carapace-bin/completers_release/faas-cli_completer/cmd"
	factor "github.com/rsteube/carapace-bin/completers_release/factor_completer/cmd"
	fakechroot "github.com/rsteube/carapace-bin/completers_release/fakechroot_completer/cmd"
	fakeroot "github.com/rsteube/carapace-bin/completers_release/fakeroot_completer/cmd"
	fastfetch "github.com/rsteube/carapace-bin/completers_release/fastfetch_completer/cmd"
	fc_cache "github.com/rsteube/carapace-bin/completers_release/fc-cache_completer/cmd"
	fc_cat "github.com/rsteube/carapace-bin/completers_release/fc-cat_completer/cmd"
	fc_conflist "github.com/rsteube/carapace-bin/completers_release/fc-conflist_completer/cmd"
	fc_list "github.com/rsteube/carapace-bin/completers_release/fc-list_completer/cmd"
	fd "github.com/rsteube/carapace-bin/completers_release/fd_completer/cmd"
	fdisk "github.com/rsteube/carapace-bin/completers_release/fdisk_completer/cmd"
	ffmpeg "github.com/rsteube/carapace-bin/completers_release/ffmpeg_completer/cmd"
	fgrep "github.com/rsteube/carapace-bin/completers_release/fgrep_completer/cmd"
	file "github.com/rsteube/carapace-bin/completers_release/file_completer/cmd"
	find "github.com/rsteube/carapace-bin/completers_release/find_completer/cmd"
	firefox "github.com/rsteube/carapace-bin/completers_release/firefox_completer/cmd"
	fish "github.com/rsteube/carapace-bin/completers_release/fish_completer/cmd"
	flatpak "github.com/rsteube/carapace-bin/completers_release/flatpak_completer/cmd"
	flutter "github.com/rsteube/carapace-bin/completers_release/flutter_completer/cmd"
	fmt "github.com/rsteube/carapace-bin/completers_release/fmt_completer/cmd"
	fold "github.com/rsteube/carapace-bin/completers_release/fold_completer/cmd"
	foot "github.com/rsteube/carapace-bin/completers_release/foot_completer/cmd"
	ftp "github.com/rsteube/carapace-bin/completers_release/ftp_completer/cmd"
	ftpd "github.com/rsteube/carapace-bin/completers_release/ftpd_completer/cmd"
	fzf "github.com/rsteube/carapace-bin/completers_release/fzf_completer/cmd"
	gatsby "github.com/rsteube/carapace-bin/completers_release/gatsby_completer/cmd"
	gcloud "github.com/rsteube/carapace-bin/completers_release/gcloud_completer/cmd"
	gdb "github.com/rsteube/carapace-bin/completers_release/gdb_completer/cmd"
	gdu "github.com/rsteube/carapace-bin/completers_release/gdu_completer/cmd"
	get_env "github.com/rsteube/carapace-bin/completers_release/get-env_completer/cmd"
	gftp "github.com/rsteube/carapace-bin/completers_release/gftp_completer/cmd"
	gh_dash "github.com/rsteube/carapace-bin/completers_release/gh-dash_completer/cmd"
	gh "github.com/rsteube/carapace-bin/completers_release/gh_completer/cmd"
	gimp "github.com/rsteube/carapace-bin/completers_release/gimp_completer/cmd"
	git_abort "github.com/rsteube/carapace-bin/completers_release/git-abort_completer/cmd"
	git_alias "github.com/rsteube/carapace-bin/completers_release/git-alias_completer/cmd"
	git_archive_file "github.com/rsteube/carapace-bin/completers_release/git-archive-file_completer/cmd"
	git_authors "github.com/rsteube/carapace-bin/completers_release/git-authors_completer/cmd"
	git_clang_format "github.com/rsteube/carapace-bin/completers_release/git-clang-format_completer/cmd"
	git_extras "github.com/rsteube/carapace-bin/completers_release/git-extras_completer/cmd"
	git_info "github.com/rsteube/carapace-bin/completers_release/git-info_completer/cmd"
	git_standup "github.com/rsteube/carapace-bin/completers_release/git-standup_completer/cmd"
	git "github.com/rsteube/carapace-bin/completers_release/git_completer/cmd"
	gitk "github.com/rsteube/carapace-bin/completers_release/gitk_completer/cmd"
	gitui "github.com/rsteube/carapace-bin/completers_release/gitui_completer/cmd"
	glab "github.com/rsteube/carapace-bin/completers_release/glab_completer/cmd"
	glow "github.com/rsteube/carapace-bin/completers_release/glow_completer/cmd"
	gnome_keyring_daemon "github.com/rsteube/carapace-bin/completers_release/gnome-keyring-daemon_completer/cmd"
	gnome_keyring "github.com/rsteube/carapace-bin/completers_release/gnome-keyring_completer/cmd"
	gnome_maps "github.com/rsteube/carapace-bin/completers_release/gnome-maps_completer/cmd"
	gnome_terminal "github.com/rsteube/carapace-bin/completers_release/gnome-terminal_completer/cmd"
	go_carpet "github.com/rsteube/carapace-bin/completers_release/go-carpet_completer/cmd"
	go_tool_asm "github.com/rsteube/carapace-bin/completers_release/go-tool-asm_completer/cmd"
	go_tool_buildid "github.com/rsteube/carapace-bin/completers_release/go-tool-buildid_completer/cmd"
	go_tool_cgo "github.com/rsteube/carapace-bin/completers_release/go-tool-cgo_completer/cmd"
	go_tool_compile "github.com/rsteube/carapace-bin/completers_release/go-tool-compile_completer/cmd"
	go_tool_covdata "github.com/rsteube/carapace-bin/completers_release/go-tool-covdata_completer/cmd"
	go_tool_cover "github.com/rsteube/carapace-bin/completers_release/go-tool-cover_completer/cmd"
	go_tool_dist "github.com/rsteube/carapace-bin/completers_release/go-tool-dist_completer/cmd"
	go_tool_doc "github.com/rsteube/carapace-bin/completers_release/go-tool-doc_completer/cmd"
	go_tool_fix "github.com/rsteube/carapace-bin/completers_release/go-tool-fix_completer/cmd"
	go_tool_link "github.com/rsteube/carapace-bin/completers_release/go-tool-link_completer/cmd"
	go_tool_nm "github.com/rsteube/carapace-bin/completers_release/go-tool-nm_completer/cmd"
	go_tool_objdump "github.com/rsteube/carapace-bin/completers_release/go-tool-objdump_completer/cmd"
	go_tool_pack "github.com/rsteube/carapace-bin/completers_release/go-tool-pack_completer/cmd"
	_go "github.com/rsteube/carapace-bin/completers_release/go_completer/cmd"
	gocyclo "github.com/rsteube/carapace-bin/completers_release/gocyclo_completer/cmd"
	gofmt "github.com/rsteube/carapace-bin/completers_release/gofmt_completer/cmd"
	goimports "github.com/rsteube/carapace-bin/completers_release/goimports_completer/cmd"
	golangci_lint "github.com/rsteube/carapace-bin/completers_release/golangci-lint_completer/cmd"
	gonew "github.com/rsteube/carapace-bin/completers_release/gonew_completer/cmd"
	google_chrome "github.com/rsteube/carapace-bin/completers_release/google-chrome_completer/cmd"
	gopls "github.com/rsteube/carapace-bin/completers_release/gopls_completer/cmd"
	goreleaser "github.com/rsteube/carapace-bin/completers_release/goreleaser_completer/cmd"
	goweight "github.com/rsteube/carapace-bin/completers_release/goweight_completer/cmd"
	gparted "github.com/rsteube/carapace-bin/completers_release/gparted_completer/cmd"
	gpasswd "github.com/rsteube/carapace-bin/completers_release/gpasswd_completer/cmd"
	gpg_agent "github.com/rsteube/carapace-bin/completers_release/gpg-agent_completer/cmd"
	gpg "github.com/rsteube/carapace-bin/completers_release/gpg_completer/cmd"
	gradle "github.com/rsteube/carapace-bin/completers_release/gradle_completer/cmd"
	grep "github.com/rsteube/carapace-bin/completers_release/grep_completer/cmd"
	groupadd "github.com/rsteube/carapace-bin/completers_release/groupadd_completer/cmd"
	groupdel "github.com/rsteube/carapace-bin/completers_release/groupdel_completer/cmd"
	groupmems "github.com/rsteube/carapace-bin/completers_release/groupmems_completer/cmd"
	groupmod "github.com/rsteube/carapace-bin/completers_release/groupmod_completer/cmd"
	groups "github.com/rsteube/carapace-bin/completers_release/groups_completer/cmd"
	grype "github.com/rsteube/carapace-bin/completers_release/grype_completer/cmd"
	gulp "github.com/rsteube/carapace-bin/completers_release/gulp_completer/cmd"
	gum "github.com/rsteube/carapace-bin/completers_release/gum_completer/cmd"
	gunzip "github.com/rsteube/carapace-bin/completers_release/gunzip_completer/cmd"
	gzip "github.com/rsteube/carapace-bin/completers_release/gzip_completer/cmd"
	halt "github.com/rsteube/carapace-bin/completers_release/halt_completer/cmd"
	head "github.com/rsteube/carapace-bin/completers_release/head_completer/cmd"
	helix "github.com/rsteube/carapace-bin/completers_release/helix_completer/cmd"
	helm "github.com/rsteube/carapace-bin/completers_release/helm_completer/cmd"
	helmsman "github.com/rsteube/carapace-bin/completers_release/helmsman_completer/cmd"
	hexchat "github.com/rsteube/carapace-bin/completers_release/hexchat_completer/cmd"
	hexdump "github.com/rsteube/carapace-bin/completers_release/hexdump_completer/cmd"
	hostid "github.com/rsteube/carapace-bin/completers_release/hostid_completer/cmd"
	hostname "github.com/rsteube/carapace-bin/completers_release/hostname_completer/cmd"
	htop "github.com/rsteube/carapace-bin/completers_release/htop_completer/cmd"
	http "github.com/rsteube/carapace-bin/completers_release/http_completer/cmd"
	https "github.com/rsteube/carapace-bin/completers_release/https_completer/cmd"
	hugo "github.com/rsteube/carapace-bin/completers_release/hugo_completer/cmd"
	hwinfo "github.com/rsteube/carapace-bin/completers_release/hwinfo_completer/cmd"
	hx "github.com/rsteube/carapace-bin/completers_release/hx_completer/cmd"
	i3_scrot "github.com/rsteube/carapace-bin/completers_release/i3-scrot_completer/cmd"
	i3 "github.com/rsteube/carapace-bin/completers_release/i3_completer/cmd"
	i3exit "github.com/rsteube/carapace-bin/completers_release/i3exit_completer/cmd"
	i3lock "github.com/rsteube/carapace-bin/completers_release/i3lock_completer/cmd"
	i3status_rs "github.com/rsteube/carapace-bin/completers_release/i3status-rs_completer/cmd"
	i3status "github.com/rsteube/carapace-bin/completers_release/i3status_completer/cmd"
	id "github.com/rsteube/carapace-bin/completers_release/id_completer/cmd"
	imv "github.com/rsteube/carapace-bin/completers_release/imv_completer/cmd"
	inkscape "github.com/rsteube/carapace-bin/completers_release/inkscape_completer/cmd"
	inshellisense "github.com/rsteube/carapace-bin/completers_release/inshellisense_completer/cmd"
	install "github.com/rsteube/carapace-bin/completers_release/install_completer/cmd"
	ion "github.com/rsteube/carapace-bin/completers_release/ion_completer/cmd"
	jar "github.com/rsteube/carapace-bin/completers_release/jar_completer/cmd"
	java "github.com/rsteube/carapace-bin/completers_release/java_completer/cmd"
	javac "github.com/rsteube/carapace-bin/completers_release/javac_completer/cmd"
	jj "github.com/rsteube/carapace-bin/completers_release/jj_completer/cmd"
	join "github.com/rsteube/carapace-bin/completers_release/join_completer/cmd"
	journalctl "github.com/rsteube/carapace-bin/completers_release/journalctl_completer/cmd"
	jq "github.com/rsteube/carapace-bin/completers_release/jq_completer/cmd"
	julia "github.com/rsteube/carapace-bin/completers_release/julia_completer/cmd"
	just "github.com/rsteube/carapace-bin/completers_release/just_completer/cmd"
	kak_lsp "github.com/rsteube/carapace-bin/completers_release/kak-lsp_completer/cmd"
	kak "github.com/rsteube/carapace-bin/completers_release/kak_completer/cmd"
	kill "github.com/rsteube/carapace-bin/completers_release/kill_completer/cmd"
	killall "github.com/rsteube/carapace-bin/completers_release/killall_completer/cmd"
	kmonad "github.com/rsteube/carapace-bin/completers_release/kmonad_completer/cmd"
	kompose "github.com/rsteube/carapace-bin/completers_release/kompose_completer/cmd"
	kotlin "github.com/rsteube/carapace-bin/completers_release/kotlin_completer/cmd"
	kotlinc "github.com/rsteube/carapace-bin/completers_release/kotlinc_completer/cmd"
	ktlint "github.com/rsteube/carapace-bin/completers_release/ktlint_completer/cmd"
	kubeadm "github.com/rsteube/carapace-bin/completers_release/kubeadm_completer/cmd"
	kubectl "github.com/rsteube/carapace-bin/completers_release/kubectl_completer/cmd"
	last "github.com/rsteube/carapace-bin/completers_release/last_completer/cmd"
	lastb "github.com/rsteube/carapace-bin/completers_release/lastb_completer/cmd"
	lastlog "github.com/rsteube/carapace-bin/completers_release/lastlog_completer/cmd"
	lazygit "github.com/rsteube/carapace-bin/completers_release/lazygit_completer/cmd"
	lf "github.com/rsteube/carapace-bin/completers_release/lf_completer/cmd"
	light "github.com/rsteube/carapace-bin/completers_release/light_completer/cmd"
	lightdm "github.com/rsteube/carapace-bin/completers_release/lightdm_completer/cmd"
	link "github.com/rsteube/carapace-bin/completers_release/link_completer/cmd"
	ln "github.com/rsteube/carapace-bin/completers_release/ln_completer/cmd"
	lnav "github.com/rsteube/carapace-bin/completers_release/lnav_completer/cmd"
	lncrawl "github.com/rsteube/carapace-bin/completers_release/lncrawl_completer/cmd"
	locale "github.com/rsteube/carapace-bin/completers_release/locale_completer/cmd"
	localectl "github.com/rsteube/carapace-bin/completers_release/localectl_completer/cmd"
	logname "github.com/rsteube/carapace-bin/completers_release/logname_completer/cmd"
	ls "github.com/rsteube/carapace-bin/completers_release/ls_completer/cmd"
	lsb_release "github.com/rsteube/carapace-bin/completers_release/lsb_release_completer/cmd"
	lsblk "github.com/rsteube/carapace-bin/completers_release/lsblk_completer/cmd"
	lscpu "github.com/rsteube/carapace-bin/completers_release/lscpu_completer/cmd"
	lslocks "github.com/rsteube/carapace-bin/completers_release/lslocks_completer/cmd"
	lslogins "github.com/rsteube/carapace-bin/completers_release/lslogins_completer/cmd"
	lsmem "github.com/rsteube/carapace-bin/completers_release/lsmem_completer/cmd"
	lsns "github.com/rsteube/carapace-bin/completers_release/lsns_completer/cmd"
	lsusb "github.com/rsteube/carapace-bin/completers_release/lsusb_completer/cmd"
	lua "github.com/rsteube/carapace-bin/completers_release/lua_completer/cmd"
	lzcat "github.com/rsteube/carapace-bin/completers_release/lzcat_completer/cmd"
	lzma "github.com/rsteube/carapace-bin/completers_release/lzma_completer/cmd"
	make "github.com/rsteube/carapace-bin/completers_release/make_completer/cmd"
	makepkg "github.com/rsteube/carapace-bin/completers_release/makepkg_completer/cmd"
	man "github.com/rsteube/carapace-bin/completers_release/man_completer/cmd"
	mcomix "github.com/rsteube/carapace-bin/completers_release/mcomix_completer/cmd"
	md5sum "github.com/rsteube/carapace-bin/completers_release/md5sum_completer/cmd"
	mdbook "github.com/rsteube/carapace-bin/completers_release/mdbook_completer/cmd"
	meld "github.com/rsteube/carapace-bin/completers_release/meld_completer/cmd"
	melt "github.com/rsteube/carapace-bin/completers_release/melt_completer/cmd"
	micro "github.com/rsteube/carapace-bin/completers_release/micro_completer/cmd"
	minikube "github.com/rsteube/carapace-bin/completers_release/minikube_completer/cmd"
	mitmproxy "github.com/rsteube/carapace-bin/completers_release/mitmproxy_completer/cmd"
	mkcert "github.com/rsteube/carapace-bin/completers_release/mkcert_completer/cmd"
	mkdir "github.com/rsteube/carapace-bin/completers_release/mkdir_completer/cmd"
	mkfifo "github.com/rsteube/carapace-bin/completers_release/mkfifo_completer/cmd"
	mkfs "github.com/rsteube/carapace-bin/completers_release/mkfs_completer/cmd"
	mknod "github.com/rsteube/carapace-bin/completers_release/mknod_completer/cmd"
	mkswap "github.com/rsteube/carapace-bin/completers_release/mkswap_completer/cmd"
	mktemp "github.com/rsteube/carapace-bin/completers_release/mktemp_completer/cmd"
	modinfo "github.com/rsteube/carapace-bin/completers_release/modinfo_completer/cmd"
	modprobe "github.com/rsteube/carapace-bin/completers_release/modprobe_completer/cmd"
	more "github.com/rsteube/carapace-bin/completers_release/more_completer/cmd"
	mosh "github.com/rsteube/carapace-bin/completers_release/mosh_completer/cmd"
	mount "github.com/rsteube/carapace-bin/completers_release/mount_completer/cmd"
	mousepad "github.com/rsteube/carapace-bin/completers_release/mousepad_completer/cmd"
	mpv "github.com/rsteube/carapace-bin/completers_release/mpv_completer/cmd"
	mv "github.com/rsteube/carapace-bin/completers_release/mv_completer/cmd"
	mvn "github.com/rsteube/carapace-bin/completers_release/mvn_completer/cmd"
	nano "github.com/rsteube/carapace-bin/completers_release/nano_completer/cmd"
	nc "github.com/rsteube/carapace-bin/completers_release/nc_completer/cmd"
	ncdu "github.com/rsteube/carapace-bin/completers_release/ncdu_completer/cmd"
	neomutt "github.com/rsteube/carapace-bin/completers_release/neomutt_completer/cmd"
	netcat "github.com/rsteube/carapace-bin/completers_release/netcat_completer/cmd"
	newman "github.com/rsteube/carapace-bin/completers_release/newman_completer/cmd"
	newrelic "github.com/rsteube/carapace-bin/completers_release/newrelic_completer/cmd"
	nfpm "github.com/rsteube/carapace-bin/completers_release/nfpm_completer/cmd"
	ng "github.com/rsteube/carapace-bin/completers_release/ng_completer/cmd"
	nice "github.com/rsteube/carapace-bin/completers_release/nice_completer/cmd"
	nilaway "github.com/rsteube/carapace-bin/completers_release/nilaway_completer/cmd"
	nix_build "github.com/rsteube/carapace-bin/completers_release/nix-build_completer/cmd"
	nix_channel "github.com/rsteube/carapace-bin/completers_release/nix-channel_completer/cmd"
	nix_instantiate "github.com/rsteube/carapace-bin/completers_release/nix-instantiate_completer/cmd"
	nix_shell "github.com/rsteube/carapace-bin/completers_release/nix-shell_completer/cmd"
	nix "github.com/rsteube/carapace-bin/completers_release/nix_completer/cmd"
	nl "github.com/rsteube/carapace-bin/completers_release/nl_completer/cmd"
	nmcli "github.com/rsteube/carapace-bin/completers_release/nmcli_completer/cmd"
	node "github.com/rsteube/carapace-bin/completers_release/node_completer/cmd"
	nohup "github.com/rsteube/carapace-bin/completers_release/nohup_completer/cmd"
	nomad "github.com/rsteube/carapace-bin/completers_release/nomad_completer/cmd"
	npm "github.com/rsteube/carapace-bin/completers_release/npm_completer/cmd"
	ntpd "github.com/rsteube/carapace-bin/completers_release/ntpd_completer/cmd"
	nu "github.com/rsteube/carapace-bin/completers_release/nu_completer/cmd"
	nvim "github.com/rsteube/carapace-bin/completers_release/nvim_completer/cmd"
	od "github.com/rsteube/carapace-bin/completers_release/od_completer/cmd"
	openscad "github.com/rsteube/carapace-bin/completers_release/openscad_completer/cmd"
	optipng "github.com/rsteube/carapace-bin/completers_release/optipng_completer/cmd"
	packer "github.com/rsteube/carapace-bin/completers_release/packer_completer/cmd"
	pacman_conf "github.com/rsteube/carapace-bin/completers_release/pacman-conf_completer/cmd"
	pacman_db_upgrade "github.com/rsteube/carapace-bin/completers_release/pacman-db-upgrade_completer/cmd"
	pacman_key "github.com/rsteube/carapace-bin/completers_release/pacman-key_completer/cmd"
	pacman_mirrors "github.com/rsteube/carapace-bin/completers_release/pacman-mirrors_completer/cmd"
	pacman "github.com/rsteube/carapace-bin/completers_release/pacman_completer/cmd"
	palemoon "github.com/rsteube/carapace-bin/completers_release/palemoon_completer/cmd"
	pamac "github.com/rsteube/carapace-bin/completers_release/pamac_completer/cmd"
	pandoc "github.com/rsteube/carapace-bin/completers_release/pandoc_completer/cmd"
	paru "github.com/rsteube/carapace-bin/completers_release/paru_completer/cmd"
	pass "github.com/rsteube/carapace-bin/completers_release/pass_completer/cmd"
	passwd "github.com/rsteube/carapace-bin/completers_release/passwd_completer/cmd"
	paste "github.com/rsteube/carapace-bin/completers_release/paste_completer/cmd"
	patch "github.com/rsteube/carapace-bin/completers_release/patch_completer/cmd"
	pathchk "github.com/rsteube/carapace-bin/completers_release/pathchk_completer/cmd"
	pcmanfm "github.com/rsteube/carapace-bin/completers_release/pcmanfm_completer/cmd"
	pgrep "github.com/rsteube/carapace-bin/completers_release/pgrep_completer/cmd"
	picard "github.com/rsteube/carapace-bin/completers_release/picard_completer/cmd"
	ping "github.com/rsteube/carapace-bin/completers_release/ping_completer/cmd"
	pinky "github.com/rsteube/carapace-bin/completers_release/pinky_completer/cmd"
	pip "github.com/rsteube/carapace-bin/completers_release/pip_completer/cmd"
	pkg "github.com/rsteube/carapace-bin/completers_release/pkg_completer/cmd"
	pkgsite "github.com/rsteube/carapace-bin/completers_release/pkgsite_completer/cmd"
	pkill "github.com/rsteube/carapace-bin/completers_release/pkill_completer/cmd"
	pnpm "github.com/rsteube/carapace-bin/completers_release/pnpm_completer/cmd"
	podman "github.com/rsteube/carapace-bin/completers_release/podman_completer/cmd"
	poweroff "github.com/rsteube/carapace-bin/completers_release/poweroff_completer/cmd"
	powertop "github.com/rsteube/carapace-bin/completers_release/powertop_completer/cmd"
	pprof "github.com/rsteube/carapace-bin/completers_release/pprof_completer/cmd"
	pr "github.com/rsteube/carapace-bin/completers_release/pr_completer/cmd"
	present "github.com/rsteube/carapace-bin/completers_release/present_completer/cmd"
	prettybat "github.com/rsteube/carapace-bin/completers_release/prettybat_completer/cmd"
	prettyping "github.com/rsteube/carapace-bin/completers_release/prettyping_completer/cmd"
	printenv "github.com/rsteube/carapace-bin/completers_release/printenv_completer/cmd"
	ps "github.com/rsteube/carapace-bin/completers_release/ps_completer/cmd"
	ptx "github.com/rsteube/carapace-bin/completers_release/ptx_completer/cmd"
	pulumi "github.com/rsteube/carapace-bin/completers_release/pulumi_completer/cmd"
	pwait "github.com/rsteube/carapace-bin/completers_release/pwait_completer/cmd"
	pwd "github.com/rsteube/carapace-bin/completers_release/pwd_completer/cmd"
	python "github.com/rsteube/carapace-bin/completers_release/python_completer/cmd"
	qmk "github.com/rsteube/carapace-bin/completers_release/qmk_completer/cmd"
	qrencode "github.com/rsteube/carapace-bin/completers_release/qrencode_completer/cmd"
	qutebrowser "github.com/rsteube/carapace-bin/completers_release/qutebrowser_completer/cmd"
	ranger "github.com/rsteube/carapace-bin/completers_release/ranger_completer/cmd"
	readlink "github.com/rsteube/carapace-bin/completers_release/readlink_completer/cmd"
	reboot "github.com/rsteube/carapace-bin/completers_release/reboot_completer/cmd"
	redis_cli "github.com/rsteube/carapace-bin/completers_release/redis-cli_completer/cmd"
	restic "github.com/rsteube/carapace-bin/completers_release/restic_completer/cmd"
	resume_cli "github.com/rsteube/carapace-bin/completers_release/resume-cli_completer/cmd"
	rg "github.com/rsteube/carapace-bin/completers_release/rg_completer/cmd"
	rifle "github.com/rsteube/carapace-bin/completers_release/rifle_completer/cmd"
	rm "github.com/rsteube/carapace-bin/completers_release/rm_completer/cmd"
	rmdir "github.com/rsteube/carapace-bin/completers_release/rmdir_completer/cmd"
	rmmod "github.com/rsteube/carapace-bin/completers_release/rmmod_completer/cmd"
	rsync "github.com/rsteube/carapace-bin/completers_release/rsync_completer/cmd"
	rust_analyzer "github.com/rsteube/carapace-bin/completers_release/rust-analyzer_completer/cmd"
	rustc "github.com/rsteube/carapace-bin/completers_release/rustc_completer/cmd"
	rustdoc "github.com/rsteube/carapace-bin/completers_release/rustdoc_completer/cmd"
	rustup "github.com/rsteube/carapace-bin/completers_release/rustup_completer/cmd"
	scp "github.com/rsteube/carapace-bin/completers_release/scp_completer/cmd"
	scrot "github.com/rsteube/carapace-bin/completers_release/scrot_completer/cmd"
	sdkmanager "github.com/rsteube/carapace-bin/completers_release/sdkmanager_completer/cmd"
	sed "github.com/rsteube/carapace-bin/completers_release/sed_completer/cmd"
	semver "github.com/rsteube/carapace-bin/completers_release/semver_completer/cmd"
	seq "github.com/rsteube/carapace-bin/completers_release/seq_completer/cmd"
	set_env "github.com/rsteube/carapace-bin/completers_release/set-env_completer/cmd"
	sftp "github.com/rsteube/carapace-bin/completers_release/sftp_completer/cmd"
	sha1sum "github.com/rsteube/carapace-bin/completers_release/sha1sum_completer/cmd"
	sha256sum "github.com/rsteube/carapace-bin/completers_release/sha256sum_completer/cmd"
	showkey "github.com/rsteube/carapace-bin/completers_release/showkey_completer/cmd"
	shred "github.com/rsteube/carapace-bin/completers_release/shred_completer/cmd"
	shutdown "github.com/rsteube/carapace-bin/completers_release/shutdown_completer/cmd"
	sleep "github.com/rsteube/carapace-bin/completers_release/sleep_completer/cmd"
	slides "github.com/rsteube/carapace-bin/completers_release/slides_completer/cmd"
	soft "github.com/rsteube/carapace-bin/completers_release/soft_completer/cmd"
	sort "github.com/rsteube/carapace-bin/completers_release/sort_completer/cmd"
	speedtest_cli "github.com/rsteube/carapace-bin/completers_release/speedtest-cli_completer/cmd"
	split "github.com/rsteube/carapace-bin/completers_release/split_completer/cmd"
	ssh_agent "github.com/rsteube/carapace-bin/completers_release/ssh-agent_completer/cmd"
	ssh_copy_id "github.com/rsteube/carapace-bin/completers_release/ssh-copy-id_completer/cmd"
	ssh_keygen "github.com/rsteube/carapace-bin/completers_release/ssh-keygen_completer/cmd"
	ssh "github.com/rsteube/carapace-bin/completers_release/ssh_completer/cmd"
	st "github.com/rsteube/carapace-bin/completers_release/st_completer/cmd"
	starship "github.com/rsteube/carapace-bin/completers_release/starship_completer/cmd"
	stat "github.com/rsteube/carapace-bin/completers_release/stat_completer/cmd"
	staticcheck "github.com/rsteube/carapace-bin/completers_release/staticcheck_completer/cmd"
	strings "github.com/rsteube/carapace-bin/completers_release/strings_completer/cmd"
	stty "github.com/rsteube/carapace-bin/completers_release/stty_completer/cmd"
	su "github.com/rsteube/carapace-bin/completers_release/su_completer/cmd"
	sudo "github.com/rsteube/carapace-bin/completers_release/sudo_completer/cmd"
	sudoedit "github.com/rsteube/carapace-bin/completers_release/sudoedit_completer/cmd"
	sudoreplay "github.com/rsteube/carapace-bin/completers_release/sudoreplay_completer/cmd"
	sulogin "github.com/rsteube/carapace-bin/completers_release/sulogin_completer/cmd"
	sum "github.com/rsteube/carapace-bin/completers_release/sum_completer/cmd"
	supervisorctl "github.com/rsteube/carapace-bin/completers_release/supervisorctl_completer/cmd"
	supervisord "github.com/rsteube/carapace-bin/completers_release/supervisord_completer/cmd"
	svg_term "github.com/rsteube/carapace-bin/completers_release/svg-term_completer/cmd"
	svgcleaner "github.com/rsteube/carapace-bin/completers_release/svgcleaner_completer/cmd"
	sway "github.com/rsteube/carapace-bin/completers_release/sway_completer/cmd"
	swaybar "github.com/rsteube/carapace-bin/completers_release/swaybar_completer/cmd"
	swaybg "github.com/rsteube/carapace-bin/completers_release/swaybg_completer/cmd"
	swayidle "github.com/rsteube/carapace-bin/completers_release/swayidle_completer/cmd"
	swaylock "github.com/rsteube/carapace-bin/completers_release/swaylock_completer/cmd"
	swaymsg "github.com/rsteube/carapace-bin/completers_release/swaymsg_completer/cmd"
	swaynag "github.com/rsteube/carapace-bin/completers_release/swaynag_completer/cmd"
	syft "github.com/rsteube/carapace-bin/completers_release/syft_completer/cmd"
	sync "github.com/rsteube/carapace-bin/completers_release/sync_completer/cmd"
	sysctl "github.com/rsteube/carapace-bin/completers_release/sysctl_completer/cmd"
	systemctl "github.com/rsteube/carapace-bin/completers_release/systemctl_completer/cmd"
	tac "github.com/rsteube/carapace-bin/completers_release/tac_completer/cmd"
	tail "github.com/rsteube/carapace-bin/completers_release/tail_completer/cmd"
	tar "github.com/rsteube/carapace-bin/completers_release/tar_completer/cmd"
	task "github.com/rsteube/carapace-bin/completers_release/task_completer/cmd"
	tea "github.com/rsteube/carapace-bin/completers_release/tea_completer/cmd"
	tee "github.com/rsteube/carapace-bin/completers_release/tee_completer/cmd"
	telnet "github.com/rsteube/carapace-bin/completers_release/telnet_completer/cmd"
	templ "github.com/rsteube/carapace-bin/completers_release/templ_completer/cmd"
	termux_apt_repo "github.com/rsteube/carapace-bin/completers_release/termux-apt-repo_completer/cmd"
	terraform_ls "github.com/rsteube/carapace-bin/completers_release/terraform-ls_completer/cmd"
	terraform "github.com/rsteube/carapace-bin/completers_release/terraform_completer/cmd"
	terragrunt "github.com/rsteube/carapace-bin/completers_release/terragrunt_completer/cmd"
	terramate "github.com/rsteube/carapace-bin/completers_release/terramate_completer/cmd"
	tesseract "github.com/rsteube/carapace-bin/completers_release/tesseract_completer/cmd"
	tig "github.com/rsteube/carapace-bin/completers_release/tig_completer/cmd"
	tinygo "github.com/rsteube/carapace-bin/completers_release/tinygo_completer/cmd"
	tldr "github.com/rsteube/carapace-bin/completers_release/tldr_completer/cmd"
	tmate "github.com/rsteube/carapace-bin/completers_release/tmate_completer/cmd"
	tmux "github.com/rsteube/carapace-bin/completers_release/tmux_completer/cmd"
	tofu "github.com/rsteube/carapace-bin/completers_release/tofu_completer/cmd"
	toit_lsp "github.com/rsteube/carapace-bin/completers_release/toit.lsp_completer/cmd"
	toit_pkg "github.com/rsteube/carapace-bin/completers_release/toit.pkg_completer/cmd"
	top "github.com/rsteube/carapace-bin/completers_release/top_completer/cmd"
	tor_browser "github.com/rsteube/carapace-bin/completers_release/tor-browser_completer/cmd"
	tor_gencert "github.com/rsteube/carapace-bin/completers_release/tor-gencert_completer/cmd"
	tor_print_ed_signing_cert "github.com/rsteube/carapace-bin/completers_release/tor-print-ed-signing-cert_completer/cmd"
	tor_resolve "github.com/rsteube/carapace-bin/completers_release/tor-resolve_completer/cmd"
	torsocks "github.com/rsteube/carapace-bin/completers_release/torsocks_completer/cmd"
	touch "github.com/rsteube/carapace-bin/completers_release/touch_completer/cmd"
	tr "github.com/rsteube/carapace-bin/completers_release/tr_completer/cmd"
	traefik "github.com/rsteube/carapace-bin/completers_release/traefik_completer/cmd"
	tree "github.com/rsteube/carapace-bin/completers_release/tree_completer/cmd"
	truncate "github.com/rsteube/carapace-bin/completers_release/truncate_completer/cmd"
	ts "github.com/rsteube/carapace-bin/completers_release/ts_completer/cmd"
	tsc "github.com/rsteube/carapace-bin/completers_release/tsc_completer/cmd"
	tsh "github.com/rsteube/carapace-bin/completers_release/tsh_completer/cmd"
	tshark "github.com/rsteube/carapace-bin/completers_release/tshark_completer/cmd"
	tsort "github.com/rsteube/carapace-bin/completers_release/tsort_completer/cmd"
	tty "github.com/rsteube/carapace-bin/completers_release/tty_completer/cmd"
	ttyd "github.com/rsteube/carapace-bin/completers_release/ttyd_completer/cmd"
	turbo "github.com/rsteube/carapace-bin/completers_release/turbo_completer/cmd"
	umount "github.com/rsteube/carapace-bin/completers_release/umount_completer/cmd"
	uname "github.com/rsteube/carapace-bin/completers_release/uname_completer/cmd"
	unbrotli "github.com/rsteube/carapace-bin/completers_release/unbrotli_completer/cmd"
	unexpand "github.com/rsteube/carapace-bin/completers_release/unexpand_completer/cmd"
	uniq "github.com/rsteube/carapace-bin/completers_release/uniq_completer/cmd"
	unlink "github.com/rsteube/carapace-bin/completers_release/unlink_completer/cmd"
	unlzma "github.com/rsteube/carapace-bin/completers_release/unlzma_completer/cmd"
	unset_env "github.com/rsteube/carapace-bin/completers_release/unset-env_completer/cmd"
	unxz "github.com/rsteube/carapace-bin/completers_release/unxz_completer/cmd"
	unzip "github.com/rsteube/carapace-bin/completers_release/unzip_completer/cmd"
	upower "github.com/rsteube/carapace-bin/completers_release/upower_completer/cmd"
	uptime "github.com/rsteube/carapace-bin/completers_release/uptime_completer/cmd"
	upx "github.com/rsteube/carapace-bin/completers_release/upx_completer/cmd"
	useradd "github.com/rsteube/carapace-bin/completers_release/useradd_completer/cmd"
	userdel "github.com/rsteube/carapace-bin/completers_release/userdel_completer/cmd"
	usermod "github.com/rsteube/carapace-bin/completers_release/usermod_completer/cmd"
	users "github.com/rsteube/carapace-bin/completers_release/users_completer/cmd"
	vagrant "github.com/rsteube/carapace-bin/completers_release/vagrant_completer/cmd"
	vault "github.com/rsteube/carapace-bin/completers_release/vault_completer/cmd"
	vdir "github.com/rsteube/carapace-bin/completers_release/vdir_completer/cmd"
	vercel "github.com/rsteube/carapace-bin/completers_release/vercel_completer/cmd"
	vhs "github.com/rsteube/carapace-bin/completers_release/vhs_completer/cmd"
	vi "github.com/rsteube/carapace-bin/completers_release/vi_completer/cmd"
	viewnior "github.com/rsteube/carapace-bin/completers_release/viewnior_completer/cmd"
	visudo "github.com/rsteube/carapace-bin/completers_release/visudo_completer/cmd"
	viu "github.com/rsteube/carapace-bin/completers_release/viu_completer/cmd"
	vivid "github.com/rsteube/carapace-bin/completers_release/vivid_completer/cmd"
	vlc "github.com/rsteube/carapace-bin/completers_release/vlc_completer/cmd"
	volta "github.com/rsteube/carapace-bin/completers_release/volta_completer/cmd"
	w "github.com/rsteube/carapace-bin/completers_release/w_completer/cmd"
	watch "github.com/rsteube/carapace-bin/completers_release/watch_completer/cmd"
	watchexec "github.com/rsteube/carapace-bin/completers_release/watchexec_completer/cmd"
	watchgnupg "github.com/rsteube/carapace-bin/completers_release/watchgnupg_completer/cmd"
	waypoint "github.com/rsteube/carapace-bin/completers_release/waypoint_completer/cmd"
	wc "github.com/rsteube/carapace-bin/completers_release/wc_completer/cmd"
	wezterm "github.com/rsteube/carapace-bin/completers_release/wezterm_completer/cmd"
	wget "github.com/rsteube/carapace-bin/completers_release/wget_completer/cmd"
	whereis "github.com/rsteube/carapace-bin/completers_release/whereis_completer/cmd"
	which "github.com/rsteube/carapace-bin/completers_release/which_completer/cmd"
	who "github.com/rsteube/carapace-bin/completers_release/who_completer/cmd"
	whoami "github.com/rsteube/carapace-bin/completers_release/whoami_completer/cmd"
	wine "github.com/rsteube/carapace-bin/completers_release/wine_completer/cmd"
	wineboot "github.com/rsteube/carapace-bin/completers_release/wineboot_completer/cmd"
	winepath "github.com/rsteube/carapace-bin/completers_release/winepath_completer/cmd"
	wineserver "github.com/rsteube/carapace-bin/completers_release/wineserver_completer/cmd"
	winetricks "github.com/rsteube/carapace-bin/completers_release/winetricks_completer/cmd"
	wire "github.com/rsteube/carapace-bin/completers_release/wire_completer/cmd"
	wireshark "github.com/rsteube/carapace-bin/completers_release/wireshark_completer/cmd"
	wishlist "github.com/rsteube/carapace-bin/completers_release/wishlist_completer/cmd"
	woeusb "github.com/rsteube/carapace-bin/completers_release/woeusb_completer/cmd"
	xargs "github.com/rsteube/carapace-bin/completers_release/xargs_completer/cmd"
	xbacklight "github.com/rsteube/carapace-bin/completers_release/xbacklight_completer/cmd"
	xdotool "github.com/rsteube/carapace-bin/completers_release/xdotool_completer/cmd"
	xonsh "github.com/rsteube/carapace-bin/completers_release/xonsh_completer/cmd"
	xz "github.com/rsteube/carapace-bin/completers_release/xz_completer/cmd"
	xzcat "github.com/rsteube/carapace-bin/completers_release/xzcat_completer/cmd"
	yarn "github.com/rsteube/carapace-bin/completers_release/yarn_completer/cmd"
	yay "github.com/rsteube/carapace-bin/completers_release/yay_completer/cmd"
	yes "github.com/rsteube/carapace-bin/completers_release/yes_completer/cmd"
	yj "github.com/rsteube/carapace-bin/completers_release/yj_completer/cmd"
	youtube_dl "github.com/rsteube/carapace-bin/completers_release/youtube-dl_completer/cmd"
	yt_dlp "github.com/rsteube/carapace-bin/completers_release/yt-dlp_completer/cmd"
	zathura "github.com/rsteube/carapace-bin/completers_release/zathura_completer/cmd"
	zcat "github.com/rsteube/carapace-bin/completers_release/zcat_completer/cmd"
	zip "github.com/rsteube/carapace-bin/completers_release/zip_completer/cmd"
	zoxide "github.com/rsteube/carapace-bin/completers_release/zoxide_completer/cmd"
)

func ExecuteCompleter(completer string) {
	switch completer {
	case "acpi":
		acpi.Execute()
	case "acpid":
		acpid.Execute()
	case "adb":
		adb.Execute()
	case "age":
		age.Execute()
	case "agg":
		agg.Execute()
	case "alsamixer":
		alsamixer.Execute()
	case "ant":
		ant.Execute()
	case "aplay":
		aplay.Execute()
	case "apropos":
		apropos.Execute()
	case "apt-cache":
		apt_cache.Execute()
	case "apt-get":
		apt_get.Execute()
	case "ar":
		ar.Execute()
	case "arecord":
		arecord.Execute()
	case "asciinema":
		asciinema.Execute()
	case "autoconf":
		autoconf.Execute()
	case "avdmanager":
		avdmanager.Execute()
	case "awk":
		awk.Execute()
	case "aws":
		aws.Execute()
	case "az":
		az.Execute()
	case "baobab":
		baobab.Execute()
	case "basename":
		basename.Execute()
	case "bash-language-server":
		bash_language_server.Execute()
	case "bash":
		bash.Execute()
	case "bat":
		bat.Execute()
	case "batdiff":
		batdiff.Execute()
	case "batgrep":
		batgrep.Execute()
	case "batman":
		batman.Execute()
	case "bats":
		bats.Execute()
	case "bc":
		bc.Execute()
	case "benthos":
		benthos.Execute()
	case "black":
		black.Execute()
	case "boundary":
		boundary.Execute()
	case "brew":
		brew.Execute()
	case "brotli":
		brotli.Execute()
	case "bun":
		bun.Execute()
	case "bunx":
		bunx.Execute()
	case "calibre":
		calibre.Execute()
	case "capslock":
		capslock.Execute()
	case "carapace":
		carapace.Execute()
	case "cargo-clippy":
		cargo_clippy.Execute()
	case "cargo-fmt":
		cargo_fmt.Execute()
	case "cargo-metadata":
		cargo_metadata.Execute()
	case "cargo-rm":
		cargo_rm.Execute()
	case "cargo-set-version":
		cargo_set_version.Execute()
	case "cargo-upgrade":
		cargo_upgrade.Execute()
	case "cargo-watch":
		cargo_watch.Execute()
	case "cargo":
		cargo.Execute()
	case "cat":
		cat.Execute()
	case "cfdisk":
		cfdisk.Execute()
	case "charm":
		charm.Execute()
	case "cheese":
		cheese.Execute()
	case "chgrp":
		chgrp.Execute()
	case "chmod":
		chmod.Execute()
	case "chown":
		chown.Execute()
	case "chpasswd":
		chpasswd.Execute()
	case "chroma":
		chroma.Execute()
	case "chromium":
		chromium.Execute()
	case "chroot":
		chroot.Execute()
	case "chsh":
		chsh.Execute()
	case "circleci":
		circleci.Execute()
	case "cksum":
		cksum.Execute()
	case "clamav-config":
		clamav_config.Execute()
	case "clamav-milter":
		clamav_milter.Execute()
	case "clambc":
		clambc.Execute()
	case "clamconf":
		clamconf.Execute()
	case "clamd":
		clamd.Execute()
	case "clamdscan":
		clamdscan.Execute()
	case "clamdtop":
		clamdtop.Execute()
	case "clamonacc":
		clamonacc.Execute()
	case "clamscan":
		clamscan.Execute()
	case "clamsubmit":
		clamsubmit.Execute()
	case "cmus":
		cmus.Execute()
	case "code-insiders":
		code_insiders.Execute()
	case "code":
		code.Execute()
	case "codecov":
		codecov.Execute()
	case "comm":
		comm.Execute()
	case "conda-content-trust":
		conda_content_trust.Execute()
	case "conda-env":
		conda_env.Execute()
	case "conda":
		conda.Execute()
	case "conky":
		conky.Execute()
	case "consul":
		consul.Execute()
	case "coredumpctl":
		coredumpctl.Execute()
	case "cp":
		cp.Execute()
	case "csplit":
		csplit.Execute()
	case "csview":
		csview.Execute()
	case "cura":
		cura.Execute()
	case "curl":
		curl.Execute()
	case "cut":
		cut.Execute()
	case "darktable-cli":
		darktable_cli.Execute()
	case "darktable":
		darktable.Execute()
	case "dart":
		dart.Execute()
	case "date":
		date.Execute()
	case "dbt":
		dbt.Execute()
	case "dc":
		dc.Execute()
	case "dd":
		dd.Execute()
	case "deadcode":
		deadcode.Execute()
	case "delta":
		delta.Execute()
	case "deno":
		deno.Execute()
	case "devbox":
		devbox.Execute()
	case "df":
		df.Execute()
	case "dfc":
		dfc.Execute()
	case "dict":
		dict.Execute()
	case "diff3":
		diff3.Execute()
	case "diff":
		diff.Execute()
	case "dig":
		dig.Execute()
	case "dir":
		dir.Execute()
	case "dircolors":
		dircolors.Execute()
	case "direnv":
		direnv.Execute()
	case "dirname":
		dirname.Execute()
	case "dive":
		dive.Execute()
	case "dlv":
		dlv.Execute()
	case "dmenu":
		dmenu.Execute()
	case "dmesg":
		dmesg.Execute()
	case "dms":
		dms.Execute()
	case "dnsmasq":
		dnsmasq.Execute()
	case "docker-buildx":
		docker_buildx.Execute()
	case "docker-compose":
		docker_compose.Execute()
	case "docker-scan":
		docker_scan.Execute()
	case "docker":
		docker.Execute()
	case "dockerd":
		dockerd.Execute()
	case "doctl":
		doctl.Execute()
	case "dos2unix":
		dos2unix.Execute()
	case "downgrade":
		downgrade.Execute()
	case "dpkg":
		dpkg.Execute()
	case "du":
		du.Execute()
	case "ebook-convert":
		ebook_convert.Execute()
	case "egrep":
		egrep.Execute()
	case "electron":
		electron.Execute()
	case "elvish":
		elvish.Execute()
	case "env":
		env.Execute()
	case "envsubst":
		envsubst.Execute()
	case "exa":
		exa.Execute()
	case "expand":
		expand.Execute()
	case "expr":
		expr.Execute()
	case "eza":
		eza.Execute()
	case "faas-cli":
		faas_cli.Execute()
	case "factor":
		factor.Execute()
	case "fakechroot":
		fakechroot.Execute()
	case "fakeroot":
		fakeroot.Execute()
	case "fastfetch":
		fastfetch.Execute()
	case "fc-cache":
		fc_cache.Execute()
	case "fc-cat":
		fc_cat.Execute()
	case "fc-conflist":
		fc_conflist.Execute()
	case "fc-list":
		fc_list.Execute()
	case "fd":
		fd.Execute()
	case "fdisk":
		fdisk.Execute()
	case "ffmpeg":
		ffmpeg.Execute()
	case "fgrep":
		fgrep.Execute()
	case "file":
		file.Execute()
	case "find":
		find.Execute()
	case "firefox":
		firefox.Execute()
	case "fish":
		fish.Execute()
	case "flatpak":
		flatpak.Execute()
	case "flutter":
		flutter.Execute()
	case "fmt":
		fmt.Execute()
	case "fold":
		fold.Execute()
	case "foot":
		foot.Execute()
	case "ftp":
		ftp.Execute()
	case "ftpd":
		ftpd.Execute()
	case "fzf":
		fzf.Execute()
	case "gatsby":
		gatsby.Execute()
	case "gcloud":
		gcloud.Execute()
	case "gdb":
		gdb.Execute()
	case "gdu":
		gdu.Execute()
	case "get-env":
		get_env.Execute()
	case "gftp":
		gftp.Execute()
	case "gh-dash":
		gh_dash.Execute()
	case "gh":
		gh.Execute()
	case "gimp":
		gimp.Execute()
	case "git-abort":
		git_abort.Execute()
	case "git-alias":
		git_alias.Execute()
	case "git-archive-file":
		git_archive_file.Execute()
	case "git-authors":
		git_authors.Execute()
	case "git-clang-format":
		git_clang_format.Execute()
	case "git-extras":
		git_extras.Execute()
	case "git-info":
		git_info.Execute()
	case "git-standup":
		git_standup.Execute()
	case "git":
		git.Execute()
	case "gitk":
		gitk.Execute()
	case "gitui":
		gitui.Execute()
	case "glab":
		glab.Execute()
	case "glow":
		glow.Execute()
	case "gnome-keyring-daemon":
		gnome_keyring_daemon.Execute()
	case "gnome-keyring":
		gnome_keyring.Execute()
	case "gnome-maps":
		gnome_maps.Execute()
	case "gnome-terminal":
		gnome_terminal.Execute()
	case "go-carpet":
		go_carpet.Execute()
	case "go-tool-asm":
		go_tool_asm.Execute()
	case "go-tool-buildid":
		go_tool_buildid.Execute()
	case "go-tool-cgo":
		go_tool_cgo.Execute()
	case "go-tool-compile":
		go_tool_compile.Execute()
	case "go-tool-covdata":
		go_tool_covdata.Execute()
	case "go-tool-cover":
		go_tool_cover.Execute()
	case "go-tool-dist":
		go_tool_dist.Execute()
	case "go-tool-doc":
		go_tool_doc.Execute()
	case "go-tool-fix":
		go_tool_fix.Execute()
	case "go-tool-link":
		go_tool_link.Execute()
	case "go-tool-nm":
		go_tool_nm.Execute()
	case "go-tool-objdump":
		go_tool_objdump.Execute()
	case "go-tool-pack":
		go_tool_pack.Execute()
	case "go":
		_go.Execute()
	case "gocyclo":
		gocyclo.Execute()
	case "gofmt":
		gofmt.Execute()
	case "goimports":
		goimports.Execute()
	case "golangci-lint":
		golangci_lint.Execute()
	case "gonew":
		gonew.Execute()
	case "google-chrome":
		google_chrome.Execute()
	case "gopls":
		gopls.Execute()
	case "goreleaser":
		goreleaser.Execute()
	case "goweight":
		goweight.Execute()
	case "gparted":
		gparted.Execute()
	case "gpasswd":
		gpasswd.Execute()
	case "gpg-agent":
		gpg_agent.Execute()
	case "gpg":
		gpg.Execute()
	case "gradle":
		gradle.Execute()
	case "grep":
		grep.Execute()
	case "groupadd":
		groupadd.Execute()
	case "groupdel":
		groupdel.Execute()
	case "groupmems":
		groupmems.Execute()
	case "groupmod":
		groupmod.Execute()
	case "groups":
		groups.Execute()
	case "grype":
		grype.Execute()
	case "gulp":
		gulp.Execute()
	case "gum":
		gum.Execute()
	case "gunzip":
		gunzip.Execute()
	case "gzip":
		gzip.Execute()
	case "halt":
		halt.Execute()
	case "head":
		head.Execute()
	case "helix":
		helix.Execute()
	case "helm":
		helm.Execute()
	case "helmsman":
		helmsman.Execute()
	case "hexchat":
		hexchat.Execute()
	case "hexdump":
		hexdump.Execute()
	case "hostid":
		hostid.Execute()
	case "hostname":
		hostname.Execute()
	case "htop":
		htop.Execute()
	case "http":
		http.Execute()
	case "https":
		https.Execute()
	case "hugo":
		hugo.Execute()
	case "hwinfo":
		hwinfo.Execute()
	case "hx":
		hx.Execute()
	case "i3-scrot":
		i3_scrot.Execute()
	case "i3":
		i3.Execute()
	case "i3exit":
		i3exit.Execute()
	case "i3lock":
		i3lock.Execute()
	case "i3status-rs":
		i3status_rs.Execute()
	case "i3status":
		i3status.Execute()
	case "id":
		id.Execute()
	case "imv":
		imv.Execute()
	case "inkscape":
		inkscape.Execute()
	case "inshellisense":
		inshellisense.Execute()
	case "install":
		install.Execute()
	case "ion":
		ion.Execute()
	case "jar":
		jar.Execute()
	case "java":
		java.Execute()
	case "javac":
		javac.Execute()
	case "jj":
		jj.Execute()
	case "join":
		join.Execute()
	case "journalctl":
		journalctl.Execute()
	case "jq":
		jq.Execute()
	case "julia":
		julia.Execute()
	case "just":
		just.Execute()
	case "kak-lsp":
		kak_lsp.Execute()
	case "kak":
		kak.Execute()
	case "kill":
		kill.Execute()
	case "killall":
		killall.Execute()
	case "kmonad":
		kmonad.Execute()
	case "kompose":
		kompose.Execute()
	case "kotlin":
		kotlin.Execute()
	case "kotlinc":
		kotlinc.Execute()
	case "ktlint":
		ktlint.Execute()
	case "kubeadm":
		kubeadm.Execute()
	case "kubectl":
		kubectl.Execute()
	case "last":
		last.Execute()
	case "lastb":
		lastb.Execute()
	case "lastlog":
		lastlog.Execute()
	case "lazygit":
		lazygit.Execute()
	case "lf":
		lf.Execute()
	case "light":
		light.Execute()
	case "lightdm":
		lightdm.Execute()
	case "link":
		link.Execute()
	case "ln":
		ln.Execute()
	case "lnav":
		lnav.Execute()
	case "lncrawl":
		lncrawl.Execute()
	case "locale":
		locale.Execute()
	case "localectl":
		localectl.Execute()
	case "logname":
		logname.Execute()
	case "ls":
		ls.Execute()
	case "lsb_release":
		lsb_release.Execute()
	case "lsblk":
		lsblk.Execute()
	case "lscpu":
		lscpu.Execute()
	case "lslocks":
		lslocks.Execute()
	case "lslogins":
		lslogins.Execute()
	case "lsmem":
		lsmem.Execute()
	case "lsns":
		lsns.Execute()
	case "lsusb":
		lsusb.Execute()
	case "lua":
		lua.Execute()
	case "lzcat":
		lzcat.Execute()
	case "lzma":
		lzma.Execute()
	case "make":
		make.Execute()
	case "makepkg":
		makepkg.Execute()
	case "man":
		man.Execute()
	case "mcomix":
		mcomix.Execute()
	case "md5sum":
		md5sum.Execute()
	case "mdbook":
		mdbook.Execute()
	case "meld":
		meld.Execute()
	case "melt":
		melt.Execute()
	case "micro":
		micro.Execute()
	case "minikube":
		minikube.Execute()
	case "mitmproxy":
		mitmproxy.Execute()
	case "mkcert":
		mkcert.Execute()
	case "mkdir":
		mkdir.Execute()
	case "mkfifo":
		mkfifo.Execute()
	case "mkfs":
		mkfs.Execute()
	case "mknod":
		mknod.Execute()
	case "mkswap":
		mkswap.Execute()
	case "mktemp":
		mktemp.Execute()
	case "modinfo":
		modinfo.Execute()
	case "modprobe":
		modprobe.Execute()
	case "more":
		more.Execute()
	case "mosh":
		mosh.Execute()
	case "mount":
		mount.Execute()
	case "mousepad":
		mousepad.Execute()
	case "mpv":
		mpv.Execute()
	case "mv":
		mv.Execute()
	case "mvn":
		mvn.Execute()
	case "nano":
		nano.Execute()
	case "nc":
		nc.Execute()
	case "ncdu":
		ncdu.Execute()
	case "neomutt":
		neomutt.Execute()
	case "netcat":
		netcat.Execute()
	case "newman":
		newman.Execute()
	case "newrelic":
		newrelic.Execute()
	case "nfpm":
		nfpm.Execute()
	case "ng":
		ng.Execute()
	case "nice":
		nice.Execute()
	case "nilaway":
		nilaway.Execute()
	case "nix-build":
		nix_build.Execute()
	case "nix-channel":
		nix_channel.Execute()
	case "nix-instantiate":
		nix_instantiate.Execute()
	case "nix-shell":
		nix_shell.Execute()
	case "nix":
		nix.Execute()
	case "nl":
		nl.Execute()
	case "nmcli":
		nmcli.Execute()
	case "node":
		node.Execute()
	case "nohup":
		nohup.Execute()
	case "nomad":
		nomad.Execute()
	case "npm":
		npm.Execute()
	case "ntpd":
		ntpd.Execute()
	case "nu":
		nu.Execute()
	case "nvim":
		nvim.Execute()
	case "od":
		od.Execute()
	case "openscad":
		openscad.Execute()
	case "optipng":
		optipng.Execute()
	case "packer":
		packer.Execute()
	case "pacman-conf":
		pacman_conf.Execute()
	case "pacman-db-upgrade":
		pacman_db_upgrade.Execute()
	case "pacman-key":
		pacman_key.Execute()
	case "pacman-mirrors":
		pacman_mirrors.Execute()
	case "pacman":
		pacman.Execute()
	case "palemoon":
		palemoon.Execute()
	case "pamac":
		pamac.Execute()
	case "pandoc":
		pandoc.Execute()
	case "paru":
		paru.Execute()
	case "pass":
		pass.Execute()
	case "passwd":
		passwd.Execute()
	case "paste":
		paste.Execute()
	case "patch":
		patch.Execute()
	case "pathchk":
		pathchk.Execute()
	case "pcmanfm":
		pcmanfm.Execute()
	case "pgrep":
		pgrep.Execute()
	case "picard":
		picard.Execute()
	case "ping":
		ping.Execute()
	case "pinky":
		pinky.Execute()
	case "pip":
		pip.Execute()
	case "pkg":
		pkg.Execute()
	case "pkgsite":
		pkgsite.Execute()
	case "pkill":
		pkill.Execute()
	case "pnpm":
		pnpm.Execute()
	case "podman":
		podman.Execute()
	case "poweroff":
		poweroff.Execute()
	case "powertop":
		powertop.Execute()
	case "pprof":
		pprof.Execute()
	case "pr":
		pr.Execute()
	case "present":
		present.Execute()
	case "prettybat":
		prettybat.Execute()
	case "prettyping":
		prettyping.Execute()
	case "printenv":
		printenv.Execute()
	case "ps":
		ps.Execute()
	case "ptx":
		ptx.Execute()
	case "pulumi":
		pulumi.Execute()
	case "pwait":
		pwait.Execute()
	case "pwd":
		pwd.Execute()
	case "python":
		python.Execute()
	case "qmk":
		qmk.Execute()
	case "qrencode":
		qrencode.Execute()
	case "qutebrowser":
		qutebrowser.Execute()
	case "ranger":
		ranger.Execute()
	case "readlink":
		readlink.Execute()
	case "reboot":
		reboot.Execute()
	case "redis-cli":
		redis_cli.Execute()
	case "restic":
		restic.Execute()
	case "resume-cli":
		resume_cli.Execute()
	case "rg":
		rg.Execute()
	case "rifle":
		rifle.Execute()
	case "rm":
		rm.Execute()
	case "rmdir":
		rmdir.Execute()
	case "rmmod":
		rmmod.Execute()
	case "rsync":
		rsync.Execute()
	case "rust-analyzer":
		rust_analyzer.Execute()
	case "rustc":
		rustc.Execute()
	case "rustdoc":
		rustdoc.Execute()
	case "rustup":
		rustup.Execute()
	case "scp":
		scp.Execute()
	case "scrot":
		scrot.Execute()
	case "sdkmanager":
		sdkmanager.Execute()
	case "sed":
		sed.Execute()
	case "semver":
		semver.Execute()
	case "seq":
		seq.Execute()
	case "set-env":
		set_env.Execute()
	case "sftp":
		sftp.Execute()
	case "sha1sum":
		sha1sum.Execute()
	case "sha256sum":
		sha256sum.Execute()
	case "showkey":
		showkey.Execute()
	case "shred":
		shred.Execute()
	case "shutdown":
		shutdown.Execute()
	case "sleep":
		sleep.Execute()
	case "slides":
		slides.Execute()
	case "soft":
		soft.Execute()
	case "sort":
		sort.Execute()
	case "speedtest-cli":
		speedtest_cli.Execute()
	case "split":
		split.Execute()
	case "ssh-agent":
		ssh_agent.Execute()
	case "ssh-copy-id":
		ssh_copy_id.Execute()
	case "ssh-keygen":
		ssh_keygen.Execute()
	case "ssh":
		ssh.Execute()
	case "st":
		st.Execute()
	case "starship":
		starship.Execute()
	case "stat":
		stat.Execute()
	case "staticcheck":
		staticcheck.Execute()
	case "strings":
		strings.Execute()
	case "stty":
		stty.Execute()
	case "su":
		su.Execute()
	case "sudo":
		sudo.Execute()
	case "sudoedit":
		sudoedit.Execute()
	case "sudoreplay":
		sudoreplay.Execute()
	case "sulogin":
		sulogin.Execute()
	case "sum":
		sum.Execute()
	case "supervisorctl":
		supervisorctl.Execute()
	case "supervisord":
		supervisord.Execute()
	case "svg-term":
		svg_term.Execute()
	case "svgcleaner":
		svgcleaner.Execute()
	case "sway":
		sway.Execute()
	case "swaybar":
		swaybar.Execute()
	case "swaybg":
		swaybg.Execute()
	case "swayidle":
		swayidle.Execute()
	case "swaylock":
		swaylock.Execute()
	case "swaymsg":
		swaymsg.Execute()
	case "swaynag":
		swaynag.Execute()
	case "syft":
		syft.Execute()
	case "sync":
		sync.Execute()
	case "sysctl":
		sysctl.Execute()
	case "systemctl":
		systemctl.Execute()
	case "tac":
		tac.Execute()
	case "tail":
		tail.Execute()
	case "tar":
		tar.Execute()
	case "task":
		task.Execute()
	case "tea":
		tea.Execute()
	case "tee":
		tee.Execute()
	case "telnet":
		telnet.Execute()
	case "templ":
		templ.Execute()
	case "termux-apt-repo":
		termux_apt_repo.Execute()
	case "terraform-ls":
		terraform_ls.Execute()
	case "terraform":
		terraform.Execute()
	case "terragrunt":
		terragrunt.Execute()
	case "terramate":
		terramate.Execute()
	case "tesseract":
		tesseract.Execute()
	case "tig":
		tig.Execute()
	case "tinygo":
		tinygo.Execute()
	case "tldr":
		tldr.Execute()
	case "tmate":
		tmate.Execute()
	case "tmux":
		tmux.Execute()
	case "tofu":
		tofu.Execute()
	case "toit.lsp":
		toit_lsp.Execute()
	case "toit.pkg":
		toit_pkg.Execute()
	case "top":
		top.Execute()
	case "tor-browser":
		tor_browser.Execute()
	case "tor-gencert":
		tor_gencert.Execute()
	case "tor-print-ed-signing-cert":
		tor_print_ed_signing_cert.Execute()
	case "tor-resolve":
		tor_resolve.Execute()
	case "torsocks":
		torsocks.Execute()
	case "touch":
		touch.Execute()
	case "tr":
		tr.Execute()
	case "traefik":
		traefik.Execute()
	case "tree":
		tree.Execute()
	case "truncate":
		truncate.Execute()
	case "ts":
		ts.Execute()
	case "tsc":
		tsc.Execute()
	case "tsh":
		tsh.Execute()
	case "tshark":
		tshark.Execute()
	case "tsort":
		tsort.Execute()
	case "tty":
		tty.Execute()
	case "ttyd":
		ttyd.Execute()
	case "turbo":
		turbo.Execute()
	case "umount":
		umount.Execute()
	case "uname":
		uname.Execute()
	case "unbrotli":
		unbrotli.Execute()
	case "unexpand":
		unexpand.Execute()
	case "uniq":
		uniq.Execute()
	case "unlink":
		unlink.Execute()
	case "unlzma":
		unlzma.Execute()
	case "unset-env":
		unset_env.Execute()
	case "unxz":
		unxz.Execute()
	case "unzip":
		unzip.Execute()
	case "upower":
		upower.Execute()
	case "uptime":
		uptime.Execute()
	case "upx":
		upx.Execute()
	case "useradd":
		useradd.Execute()
	case "userdel":
		userdel.Execute()
	case "usermod":
		usermod.Execute()
	case "users":
		users.Execute()
	case "vagrant":
		vagrant.Execute()
	case "vault":
		vault.Execute()
	case "vdir":
		vdir.Execute()
	case "vercel":
		vercel.Execute()
	case "vhs":
		vhs.Execute()
	case "vi":
		vi.Execute()
	case "viewnior":
		viewnior.Execute()
	case "visudo":
		visudo.Execute()
	case "viu":
		viu.Execute()
	case "vivid":
		vivid.Execute()
	case "vlc":
		vlc.Execute()
	case "volta":
		volta.Execute()
	case "w":
		w.Execute()
	case "watch":
		watch.Execute()
	case "watchexec":
		watchexec.Execute()
	case "watchgnupg":
		watchgnupg.Execute()
	case "waypoint":
		waypoint.Execute()
	case "wc":
		wc.Execute()
	case "wezterm":
		wezterm.Execute()
	case "wget":
		wget.Execute()
	case "whereis":
		whereis.Execute()
	case "which":
		which.Execute()
	case "who":
		who.Execute()
	case "whoami":
		whoami.Execute()
	case "wine":
		wine.Execute()
	case "wineboot":
		wineboot.Execute()
	case "winepath":
		winepath.Execute()
	case "wineserver":
		wineserver.Execute()
	case "winetricks":
		winetricks.Execute()
	case "wire":
		wire.Execute()
	case "wireshark":
		wireshark.Execute()
	case "wishlist":
		wishlist.Execute()
	case "woeusb":
		woeusb.Execute()
	case "xargs":
		xargs.Execute()
	case "xbacklight":
		xbacklight.Execute()
	case "xdotool":
		xdotool.Execute()
	case "xonsh":
		xonsh.Execute()
	case "xz":
		xz.Execute()
	case "xzcat":
		xzcat.Execute()
	case "yarn":
		yarn.Execute()
	case "yay":
		yay.Execute()
	case "yes":
		yes.Execute()
	case "yj":
		yj.Execute()
	case "youtube-dl":
		youtube_dl.Execute()
	case "yt-dlp":
		yt_dlp.Execute()
	case "zathura":
		zathura.Execute()
	case "zcat":
		zcat.Execute()
	case "zip":
		zip.Execute()
	case "zoxide":
		zoxide.Execute()
	}
}
