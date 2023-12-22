//go:build !release

package completers

import (
	acpi "github.com/rsteube/carapace-bin/completers/acpi_completer/cmd"
	acpid "github.com/rsteube/carapace-bin/completers/acpid_completer/cmd"
	adb "github.com/rsteube/carapace-bin/completers/adb_completer/cmd"
	age "github.com/rsteube/carapace-bin/completers/age_completer/cmd"
	agg "github.com/rsteube/carapace-bin/completers/agg_completer/cmd"
	alsamixer "github.com/rsteube/carapace-bin/completers/alsamixer_completer/cmd"
	ant "github.com/rsteube/carapace-bin/completers/ant_completer/cmd"
	aplay "github.com/rsteube/carapace-bin/completers/aplay_completer/cmd"
	apropos "github.com/rsteube/carapace-bin/completers/apropos_completer/cmd"
	apt_cache "github.com/rsteube/carapace-bin/completers/apt-cache_completer/cmd"
	apt_get "github.com/rsteube/carapace-bin/completers/apt-get_completer/cmd"
	ar "github.com/rsteube/carapace-bin/completers/ar_completer/cmd"
	arecord "github.com/rsteube/carapace-bin/completers/arecord_completer/cmd"
	asciinema "github.com/rsteube/carapace-bin/completers/asciinema_completer/cmd"
	autoconf "github.com/rsteube/carapace-bin/completers/autoconf_completer/cmd"
	avdmanager "github.com/rsteube/carapace-bin/completers/avdmanager_completer/cmd"
	awk "github.com/rsteube/carapace-bin/completers/awk_completer/cmd"
	aws "github.com/rsteube/carapace-bin/completers/aws_completer/cmd"
	az "github.com/rsteube/carapace-bin/completers/az_completer/cmd"
	baobab "github.com/rsteube/carapace-bin/completers/baobab_completer/cmd"
	basename "github.com/rsteube/carapace-bin/completers/basename_completer/cmd"
	bash_language_server "github.com/rsteube/carapace-bin/completers/bash-language-server_completer/cmd"
	bash "github.com/rsteube/carapace-bin/completers/bash_completer/cmd"
	bat "github.com/rsteube/carapace-bin/completers/bat_completer/cmd"
	batdiff "github.com/rsteube/carapace-bin/completers/batdiff_completer/cmd"
	batgrep "github.com/rsteube/carapace-bin/completers/batgrep_completer/cmd"
	batman "github.com/rsteube/carapace-bin/completers/batman_completer/cmd"
	bats "github.com/rsteube/carapace-bin/completers/bats_completer/cmd"
	bc "github.com/rsteube/carapace-bin/completers/bc_completer/cmd"
	benthos "github.com/rsteube/carapace-bin/completers/benthos_completer/cmd"
	black "github.com/rsteube/carapace-bin/completers/black_completer/cmd"
	boundary "github.com/rsteube/carapace-bin/completers/boundary_completer/cmd"
	brew "github.com/rsteube/carapace-bin/completers/brew_completer/cmd"
	brotli "github.com/rsteube/carapace-bin/completers/brotli_completer/cmd"
	bun "github.com/rsteube/carapace-bin/completers/bun_completer/cmd"
	bunx "github.com/rsteube/carapace-bin/completers/bunx_completer/cmd"
	calibre "github.com/rsteube/carapace-bin/completers/calibre_completer/cmd"
	capslock "github.com/rsteube/carapace-bin/completers/capslock_completer/cmd"
	carapace "github.com/rsteube/carapace-bin/completers/carapace_completer/cmd"
	cargo_clippy "github.com/rsteube/carapace-bin/completers/cargo-clippy_completer/cmd"
	cargo_fmt "github.com/rsteube/carapace-bin/completers/cargo-fmt_completer/cmd"
	cargo_metadata "github.com/rsteube/carapace-bin/completers/cargo-metadata_completer/cmd"
	cargo_rm "github.com/rsteube/carapace-bin/completers/cargo-rm_completer/cmd"
	cargo_set_version "github.com/rsteube/carapace-bin/completers/cargo-set-version_completer/cmd"
	cargo_upgrade "github.com/rsteube/carapace-bin/completers/cargo-upgrade_completer/cmd"
	cargo_watch "github.com/rsteube/carapace-bin/completers/cargo-watch_completer/cmd"
	cargo "github.com/rsteube/carapace-bin/completers/cargo_completer/cmd"
	cat "github.com/rsteube/carapace-bin/completers/cat_completer/cmd"
	cfdisk "github.com/rsteube/carapace-bin/completers/cfdisk_completer/cmd"
	charm "github.com/rsteube/carapace-bin/completers/charm_completer/cmd"
	cheese "github.com/rsteube/carapace-bin/completers/cheese_completer/cmd"
	chgrp "github.com/rsteube/carapace-bin/completers/chgrp_completer/cmd"
	chmod "github.com/rsteube/carapace-bin/completers/chmod_completer/cmd"
	chown "github.com/rsteube/carapace-bin/completers/chown_completer/cmd"
	chpasswd "github.com/rsteube/carapace-bin/completers/chpasswd_completer/cmd"
	chroma "github.com/rsteube/carapace-bin/completers/chroma_completer/cmd"
	chromium "github.com/rsteube/carapace-bin/completers/chromium_completer/cmd"
	chroot "github.com/rsteube/carapace-bin/completers/chroot_completer/cmd"
	chsh "github.com/rsteube/carapace-bin/completers/chsh_completer/cmd"
	circleci "github.com/rsteube/carapace-bin/completers/circleci_completer/cmd"
	cksum "github.com/rsteube/carapace-bin/completers/cksum_completer/cmd"
	clamav_config "github.com/rsteube/carapace-bin/completers/clamav-config_completer/cmd"
	clamav_milter "github.com/rsteube/carapace-bin/completers/clamav-milter_completer/cmd"
	clambc "github.com/rsteube/carapace-bin/completers/clambc_completer/cmd"
	clamconf "github.com/rsteube/carapace-bin/completers/clamconf_completer/cmd"
	clamd "github.com/rsteube/carapace-bin/completers/clamd_completer/cmd"
	clamdscan "github.com/rsteube/carapace-bin/completers/clamdscan_completer/cmd"
	clamdtop "github.com/rsteube/carapace-bin/completers/clamdtop_completer/cmd"
	clamonacc "github.com/rsteube/carapace-bin/completers/clamonacc_completer/cmd"
	clamscan "github.com/rsteube/carapace-bin/completers/clamscan_completer/cmd"
	clamsubmit "github.com/rsteube/carapace-bin/completers/clamsubmit_completer/cmd"
	cmus "github.com/rsteube/carapace-bin/completers/cmus_completer/cmd"
	code_insiders "github.com/rsteube/carapace-bin/completers/code-insiders_completer/cmd"
	code "github.com/rsteube/carapace-bin/completers/code_completer/cmd"
	codecov "github.com/rsteube/carapace-bin/completers/codecov_completer/cmd"
	comm "github.com/rsteube/carapace-bin/completers/comm_completer/cmd"
	conda_content_trust "github.com/rsteube/carapace-bin/completers/conda-content-trust_completer/cmd"
	conda_env "github.com/rsteube/carapace-bin/completers/conda-env_completer/cmd"
	conda "github.com/rsteube/carapace-bin/completers/conda_completer/cmd"
	conky "github.com/rsteube/carapace-bin/completers/conky_completer/cmd"
	consul "github.com/rsteube/carapace-bin/completers/consul_completer/cmd"
	coredumpctl "github.com/rsteube/carapace-bin/completers/coredumpctl_completer/cmd"
	cp "github.com/rsteube/carapace-bin/completers/cp_completer/cmd"
	csplit "github.com/rsteube/carapace-bin/completers/csplit_completer/cmd"
	csview "github.com/rsteube/carapace-bin/completers/csview_completer/cmd"
	cura "github.com/rsteube/carapace-bin/completers/cura_completer/cmd"
	curl "github.com/rsteube/carapace-bin/completers/curl_completer/cmd"
	cut "github.com/rsteube/carapace-bin/completers/cut_completer/cmd"
	darktable_cli "github.com/rsteube/carapace-bin/completers/darktable-cli_completer/cmd"
	darktable "github.com/rsteube/carapace-bin/completers/darktable_completer/cmd"
	dart "github.com/rsteube/carapace-bin/completers/dart_completer/cmd"
	date "github.com/rsteube/carapace-bin/completers/date_completer/cmd"
	dbt "github.com/rsteube/carapace-bin/completers/dbt_completer/cmd"
	dc "github.com/rsteube/carapace-bin/completers/dc_completer/cmd"
	dd "github.com/rsteube/carapace-bin/completers/dd_completer/cmd"
	deadcode "github.com/rsteube/carapace-bin/completers/deadcode_completer/cmd"
	delta "github.com/rsteube/carapace-bin/completers/delta_completer/cmd"
	deno "github.com/rsteube/carapace-bin/completers/deno_completer/cmd"
	devbox "github.com/rsteube/carapace-bin/completers/devbox_completer/cmd"
	df "github.com/rsteube/carapace-bin/completers/df_completer/cmd"
	dfc "github.com/rsteube/carapace-bin/completers/dfc_completer/cmd"
	dict "github.com/rsteube/carapace-bin/completers/dict_completer/cmd"
	diff3 "github.com/rsteube/carapace-bin/completers/diff3_completer/cmd"
	diff "github.com/rsteube/carapace-bin/completers/diff_completer/cmd"
	dig "github.com/rsteube/carapace-bin/completers/dig_completer/cmd"
	dir "github.com/rsteube/carapace-bin/completers/dir_completer/cmd"
	dircolors "github.com/rsteube/carapace-bin/completers/dircolors_completer/cmd"
	direnv "github.com/rsteube/carapace-bin/completers/direnv_completer/cmd"
	dirname "github.com/rsteube/carapace-bin/completers/dirname_completer/cmd"
	dive "github.com/rsteube/carapace-bin/completers/dive_completer/cmd"
	dlv "github.com/rsteube/carapace-bin/completers/dlv_completer/cmd"
	dmenu "github.com/rsteube/carapace-bin/completers/dmenu_completer/cmd"
	dmesg "github.com/rsteube/carapace-bin/completers/dmesg_completer/cmd"
	dms "github.com/rsteube/carapace-bin/completers/dms_completer/cmd"
	dnsmasq "github.com/rsteube/carapace-bin/completers/dnsmasq_completer/cmd"
	docker_buildx "github.com/rsteube/carapace-bin/completers/docker-buildx_completer/cmd"
	docker_compose "github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd"
	docker_scan "github.com/rsteube/carapace-bin/completers/docker-scan_completer/cmd"
	docker "github.com/rsteube/carapace-bin/completers/docker_completer/cmd"
	dockerd "github.com/rsteube/carapace-bin/completers/dockerd_completer/cmd"
	doctl "github.com/rsteube/carapace-bin/completers/doctl_completer/cmd"
	dos2unix "github.com/rsteube/carapace-bin/completers/dos2unix_completer/cmd"
	downgrade "github.com/rsteube/carapace-bin/completers/downgrade_completer/cmd"
	dpkg "github.com/rsteube/carapace-bin/completers/dpkg_completer/cmd"
	du "github.com/rsteube/carapace-bin/completers/du_completer/cmd"
	ebook_convert "github.com/rsteube/carapace-bin/completers/ebook-convert_completer/cmd"
	egrep "github.com/rsteube/carapace-bin/completers/egrep_completer/cmd"
	electron "github.com/rsteube/carapace-bin/completers/electron_completer/cmd"
	elvish "github.com/rsteube/carapace-bin/completers/elvish_completer/cmd"
	env "github.com/rsteube/carapace-bin/completers/env_completer/cmd"
	envsubst "github.com/rsteube/carapace-bin/completers/envsubst_completer/cmd"
	exa "github.com/rsteube/carapace-bin/completers/exa_completer/cmd"
	expand "github.com/rsteube/carapace-bin/completers/expand_completer/cmd"
	expr "github.com/rsteube/carapace-bin/completers/expr_completer/cmd"
	eza "github.com/rsteube/carapace-bin/completers/eza_completer/cmd"
	faas_cli "github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd"
	factor "github.com/rsteube/carapace-bin/completers/factor_completer/cmd"
	fakechroot "github.com/rsteube/carapace-bin/completers/fakechroot_completer/cmd"
	fakeroot "github.com/rsteube/carapace-bin/completers/fakeroot_completer/cmd"
	fastfetch "github.com/rsteube/carapace-bin/completers/fastfetch_completer/cmd"
	fc_cache "github.com/rsteube/carapace-bin/completers/fc-cache_completer/cmd"
	fc_cat "github.com/rsteube/carapace-bin/completers/fc-cat_completer/cmd"
	fc_conflist "github.com/rsteube/carapace-bin/completers/fc-conflist_completer/cmd"
	fc_list "github.com/rsteube/carapace-bin/completers/fc-list_completer/cmd"
	fd "github.com/rsteube/carapace-bin/completers/fd_completer/cmd"
	fdisk "github.com/rsteube/carapace-bin/completers/fdisk_completer/cmd"
	ffmpeg "github.com/rsteube/carapace-bin/completers/ffmpeg_completer/cmd"
	fgrep "github.com/rsteube/carapace-bin/completers/fgrep_completer/cmd"
	file "github.com/rsteube/carapace-bin/completers/file_completer/cmd"
	find "github.com/rsteube/carapace-bin/completers/find_completer/cmd"
	firefox "github.com/rsteube/carapace-bin/completers/firefox_completer/cmd"
	fish "github.com/rsteube/carapace-bin/completers/fish_completer/cmd"
	flatpak "github.com/rsteube/carapace-bin/completers/flatpak_completer/cmd"
	flutter "github.com/rsteube/carapace-bin/completers/flutter_completer/cmd"
	fmt "github.com/rsteube/carapace-bin/completers/fmt_completer/cmd"
	fold "github.com/rsteube/carapace-bin/completers/fold_completer/cmd"
	foot "github.com/rsteube/carapace-bin/completers/foot_completer/cmd"
	ftp "github.com/rsteube/carapace-bin/completers/ftp_completer/cmd"
	ftpd "github.com/rsteube/carapace-bin/completers/ftpd_completer/cmd"
	fzf "github.com/rsteube/carapace-bin/completers/fzf_completer/cmd"
	gatsby "github.com/rsteube/carapace-bin/completers/gatsby_completer/cmd"
	gcloud "github.com/rsteube/carapace-bin/completers/gcloud_completer/cmd"
	gdb "github.com/rsteube/carapace-bin/completers/gdb_completer/cmd"
	gdu "github.com/rsteube/carapace-bin/completers/gdu_completer/cmd"
	get_env "github.com/rsteube/carapace-bin/completers/get-env_completer/cmd"
	gftp "github.com/rsteube/carapace-bin/completers/gftp_completer/cmd"
	gh_dash "github.com/rsteube/carapace-bin/completers/gh-dash_completer/cmd"
	gh "github.com/rsteube/carapace-bin/completers/gh_completer/cmd"
	gimp "github.com/rsteube/carapace-bin/completers/gimp_completer/cmd"
	git_abort "github.com/rsteube/carapace-bin/completers/git-abort_completer/cmd"
	git_alias "github.com/rsteube/carapace-bin/completers/git-alias_completer/cmd"
	git_archive_file "github.com/rsteube/carapace-bin/completers/git-archive-file_completer/cmd"
	git_authors "github.com/rsteube/carapace-bin/completers/git-authors_completer/cmd"
	git_clang_format "github.com/rsteube/carapace-bin/completers/git-clang-format_completer/cmd"
	git_extras "github.com/rsteube/carapace-bin/completers/git-extras_completer/cmd"
	git_info "github.com/rsteube/carapace-bin/completers/git-info_completer/cmd"
	git_standup "github.com/rsteube/carapace-bin/completers/git-standup_completer/cmd"
	git "github.com/rsteube/carapace-bin/completers/git_completer/cmd"
	gitk "github.com/rsteube/carapace-bin/completers/gitk_completer/cmd"
	gitui "github.com/rsteube/carapace-bin/completers/gitui_completer/cmd"
	glab "github.com/rsteube/carapace-bin/completers/glab_completer/cmd"
	glow "github.com/rsteube/carapace-bin/completers/glow_completer/cmd"
	gnome_keyring_daemon "github.com/rsteube/carapace-bin/completers/gnome-keyring-daemon_completer/cmd"
	gnome_keyring "github.com/rsteube/carapace-bin/completers/gnome-keyring_completer/cmd"
	gnome_maps "github.com/rsteube/carapace-bin/completers/gnome-maps_completer/cmd"
	gnome_terminal "github.com/rsteube/carapace-bin/completers/gnome-terminal_completer/cmd"
	go_carpet "github.com/rsteube/carapace-bin/completers/go-carpet_completer/cmd"
	go_tool_asm "github.com/rsteube/carapace-bin/completers/go-tool-asm_completer/cmd"
	go_tool_buildid "github.com/rsteube/carapace-bin/completers/go-tool-buildid_completer/cmd"
	go_tool_cgo "github.com/rsteube/carapace-bin/completers/go-tool-cgo_completer/cmd"
	go_tool_compile "github.com/rsteube/carapace-bin/completers/go-tool-compile_completer/cmd"
	go_tool_covdata "github.com/rsteube/carapace-bin/completers/go-tool-covdata_completer/cmd"
	go_tool_cover "github.com/rsteube/carapace-bin/completers/go-tool-cover_completer/cmd"
	go_tool_dist "github.com/rsteube/carapace-bin/completers/go-tool-dist_completer/cmd"
	go_tool_doc "github.com/rsteube/carapace-bin/completers/go-tool-doc_completer/cmd"
	go_tool_fix "github.com/rsteube/carapace-bin/completers/go-tool-fix_completer/cmd"
	go_tool_link "github.com/rsteube/carapace-bin/completers/go-tool-link_completer/cmd"
	go_tool_nm "github.com/rsteube/carapace-bin/completers/go-tool-nm_completer/cmd"
	go_tool_objdump "github.com/rsteube/carapace-bin/completers/go-tool-objdump_completer/cmd"
	go_tool_pack "github.com/rsteube/carapace-bin/completers/go-tool-pack_completer/cmd"
	_go "github.com/rsteube/carapace-bin/completers/go_completer/cmd"
	gocyclo "github.com/rsteube/carapace-bin/completers/gocyclo_completer/cmd"
	gofmt "github.com/rsteube/carapace-bin/completers/gofmt_completer/cmd"
	goimports "github.com/rsteube/carapace-bin/completers/goimports_completer/cmd"
	golangci_lint "github.com/rsteube/carapace-bin/completers/golangci-lint_completer/cmd"
	gonew "github.com/rsteube/carapace-bin/completers/gonew_completer/cmd"
	google_chrome "github.com/rsteube/carapace-bin/completers/google-chrome_completer/cmd"
	gopls "github.com/rsteube/carapace-bin/completers/gopls_completer/cmd"
	goreleaser "github.com/rsteube/carapace-bin/completers/goreleaser_completer/cmd"
	goweight "github.com/rsteube/carapace-bin/completers/goweight_completer/cmd"
	gparted "github.com/rsteube/carapace-bin/completers/gparted_completer/cmd"
	gpasswd "github.com/rsteube/carapace-bin/completers/gpasswd_completer/cmd"
	gpg_agent "github.com/rsteube/carapace-bin/completers/gpg-agent_completer/cmd"
	gpg "github.com/rsteube/carapace-bin/completers/gpg_completer/cmd"
	gradle "github.com/rsteube/carapace-bin/completers/gradle_completer/cmd"
	grep "github.com/rsteube/carapace-bin/completers/grep_completer/cmd"
	groupadd "github.com/rsteube/carapace-bin/completers/groupadd_completer/cmd"
	groupdel "github.com/rsteube/carapace-bin/completers/groupdel_completer/cmd"
	groupmems "github.com/rsteube/carapace-bin/completers/groupmems_completer/cmd"
	groupmod "github.com/rsteube/carapace-bin/completers/groupmod_completer/cmd"
	groups "github.com/rsteube/carapace-bin/completers/groups_completer/cmd"
	grype "github.com/rsteube/carapace-bin/completers/grype_completer/cmd"
	gulp "github.com/rsteube/carapace-bin/completers/gulp_completer/cmd"
	gum "github.com/rsteube/carapace-bin/completers/gum_completer/cmd"
	gunzip "github.com/rsteube/carapace-bin/completers/gunzip_completer/cmd"
	gzip "github.com/rsteube/carapace-bin/completers/gzip_completer/cmd"
	halt "github.com/rsteube/carapace-bin/completers/halt_completer/cmd"
	head "github.com/rsteube/carapace-bin/completers/head_completer/cmd"
	helix "github.com/rsteube/carapace-bin/completers/helix_completer/cmd"
	helm "github.com/rsteube/carapace-bin/completers/helm_completer/cmd"
	helmsman "github.com/rsteube/carapace-bin/completers/helmsman_completer/cmd"
	hexchat "github.com/rsteube/carapace-bin/completers/hexchat_completer/cmd"
	hexdump "github.com/rsteube/carapace-bin/completers/hexdump_completer/cmd"
	hostid "github.com/rsteube/carapace-bin/completers/hostid_completer/cmd"
	hostname "github.com/rsteube/carapace-bin/completers/hostname_completer/cmd"
	htop "github.com/rsteube/carapace-bin/completers/htop_completer/cmd"
	http "github.com/rsteube/carapace-bin/completers/http_completer/cmd"
	https "github.com/rsteube/carapace-bin/completers/https_completer/cmd"
	hugo "github.com/rsteube/carapace-bin/completers/hugo_completer/cmd"
	hwinfo "github.com/rsteube/carapace-bin/completers/hwinfo_completer/cmd"
	hx "github.com/rsteube/carapace-bin/completers/hx_completer/cmd"
	i3_scrot "github.com/rsteube/carapace-bin/completers/i3-scrot_completer/cmd"
	i3 "github.com/rsteube/carapace-bin/completers/i3_completer/cmd"
	i3exit "github.com/rsteube/carapace-bin/completers/i3exit_completer/cmd"
	i3lock "github.com/rsteube/carapace-bin/completers/i3lock_completer/cmd"
	i3status_rs "github.com/rsteube/carapace-bin/completers/i3status-rs_completer/cmd"
	i3status "github.com/rsteube/carapace-bin/completers/i3status_completer/cmd"
	id "github.com/rsteube/carapace-bin/completers/id_completer/cmd"
	imv "github.com/rsteube/carapace-bin/completers/imv_completer/cmd"
	inkscape "github.com/rsteube/carapace-bin/completers/inkscape_completer/cmd"
	inshellisense "github.com/rsteube/carapace-bin/completers/inshellisense_completer/cmd"
	install "github.com/rsteube/carapace-bin/completers/install_completer/cmd"
	ion "github.com/rsteube/carapace-bin/completers/ion_completer/cmd"
	jar "github.com/rsteube/carapace-bin/completers/jar_completer/cmd"
	java "github.com/rsteube/carapace-bin/completers/java_completer/cmd"
	javac "github.com/rsteube/carapace-bin/completers/javac_completer/cmd"
	jj "github.com/rsteube/carapace-bin/completers/jj_completer/cmd"
	join "github.com/rsteube/carapace-bin/completers/join_completer/cmd"
	journalctl "github.com/rsteube/carapace-bin/completers/journalctl_completer/cmd"
	jq "github.com/rsteube/carapace-bin/completers/jq_completer/cmd"
	julia "github.com/rsteube/carapace-bin/completers/julia_completer/cmd"
	just "github.com/rsteube/carapace-bin/completers/just_completer/cmd"
	kak_lsp "github.com/rsteube/carapace-bin/completers/kak-lsp_completer/cmd"
	kak "github.com/rsteube/carapace-bin/completers/kak_completer/cmd"
	kill "github.com/rsteube/carapace-bin/completers/kill_completer/cmd"
	killall "github.com/rsteube/carapace-bin/completers/killall_completer/cmd"
	kmonad "github.com/rsteube/carapace-bin/completers/kmonad_completer/cmd"
	kompose "github.com/rsteube/carapace-bin/completers/kompose_completer/cmd"
	kotlin "github.com/rsteube/carapace-bin/completers/kotlin_completer/cmd"
	kotlinc "github.com/rsteube/carapace-bin/completers/kotlinc_completer/cmd"
	ktlint "github.com/rsteube/carapace-bin/completers/ktlint_completer/cmd"
	kubeadm "github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd"
	kubectl "github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd"
	last "github.com/rsteube/carapace-bin/completers/last_completer/cmd"
	lastb "github.com/rsteube/carapace-bin/completers/lastb_completer/cmd"
	lastlog "github.com/rsteube/carapace-bin/completers/lastlog_completer/cmd"
	lazygit "github.com/rsteube/carapace-bin/completers/lazygit_completer/cmd"
	lf "github.com/rsteube/carapace-bin/completers/lf_completer/cmd"
	light "github.com/rsteube/carapace-bin/completers/light_completer/cmd"
	lightdm "github.com/rsteube/carapace-bin/completers/lightdm_completer/cmd"
	link "github.com/rsteube/carapace-bin/completers/link_completer/cmd"
	ln "github.com/rsteube/carapace-bin/completers/ln_completer/cmd"
	lnav "github.com/rsteube/carapace-bin/completers/lnav_completer/cmd"
	lncrawl "github.com/rsteube/carapace-bin/completers/lncrawl_completer/cmd"
	locale "github.com/rsteube/carapace-bin/completers/locale_completer/cmd"
	localectl "github.com/rsteube/carapace-bin/completers/localectl_completer/cmd"
	logname "github.com/rsteube/carapace-bin/completers/logname_completer/cmd"
	ls "github.com/rsteube/carapace-bin/completers/ls_completer/cmd"
	lsb_release "github.com/rsteube/carapace-bin/completers/lsb_release_completer/cmd"
	lsblk "github.com/rsteube/carapace-bin/completers/lsblk_completer/cmd"
	lscpu "github.com/rsteube/carapace-bin/completers/lscpu_completer/cmd"
	lslocks "github.com/rsteube/carapace-bin/completers/lslocks_completer/cmd"
	lslogins "github.com/rsteube/carapace-bin/completers/lslogins_completer/cmd"
	lsmem "github.com/rsteube/carapace-bin/completers/lsmem_completer/cmd"
	lsns "github.com/rsteube/carapace-bin/completers/lsns_completer/cmd"
	lsusb "github.com/rsteube/carapace-bin/completers/lsusb_completer/cmd"
	lua "github.com/rsteube/carapace-bin/completers/lua_completer/cmd"
	lzcat "github.com/rsteube/carapace-bin/completers/lzcat_completer/cmd"
	lzma "github.com/rsteube/carapace-bin/completers/lzma_completer/cmd"
	make "github.com/rsteube/carapace-bin/completers/make_completer/cmd"
	makepkg "github.com/rsteube/carapace-bin/completers/makepkg_completer/cmd"
	man "github.com/rsteube/carapace-bin/completers/man_completer/cmd"
	mcomix "github.com/rsteube/carapace-bin/completers/mcomix_completer/cmd"
	md5sum "github.com/rsteube/carapace-bin/completers/md5sum_completer/cmd"
	mdbook "github.com/rsteube/carapace-bin/completers/mdbook_completer/cmd"
	meld "github.com/rsteube/carapace-bin/completers/meld_completer/cmd"
	melt "github.com/rsteube/carapace-bin/completers/melt_completer/cmd"
	micro "github.com/rsteube/carapace-bin/completers/micro_completer/cmd"
	minikube "github.com/rsteube/carapace-bin/completers/minikube_completer/cmd"
	mitmproxy "github.com/rsteube/carapace-bin/completers/mitmproxy_completer/cmd"
	mkcert "github.com/rsteube/carapace-bin/completers/mkcert_completer/cmd"
	mkdir "github.com/rsteube/carapace-bin/completers/mkdir_completer/cmd"
	mkfifo "github.com/rsteube/carapace-bin/completers/mkfifo_completer/cmd"
	mkfs "github.com/rsteube/carapace-bin/completers/mkfs_completer/cmd"
	mknod "github.com/rsteube/carapace-bin/completers/mknod_completer/cmd"
	mkswap "github.com/rsteube/carapace-bin/completers/mkswap_completer/cmd"
	mktemp "github.com/rsteube/carapace-bin/completers/mktemp_completer/cmd"
	modinfo "github.com/rsteube/carapace-bin/completers/modinfo_completer/cmd"
	modprobe "github.com/rsteube/carapace-bin/completers/modprobe_completer/cmd"
	more "github.com/rsteube/carapace-bin/completers/more_completer/cmd"
	mosh "github.com/rsteube/carapace-bin/completers/mosh_completer/cmd"
	mount "github.com/rsteube/carapace-bin/completers/mount_completer/cmd"
	mousepad "github.com/rsteube/carapace-bin/completers/mousepad_completer/cmd"
	mpv "github.com/rsteube/carapace-bin/completers/mpv_completer/cmd"
	mv "github.com/rsteube/carapace-bin/completers/mv_completer/cmd"
	mvn "github.com/rsteube/carapace-bin/completers/mvn_completer/cmd"
	nano "github.com/rsteube/carapace-bin/completers/nano_completer/cmd"
	nc "github.com/rsteube/carapace-bin/completers/nc_completer/cmd"
	ncdu "github.com/rsteube/carapace-bin/completers/ncdu_completer/cmd"
	neomutt "github.com/rsteube/carapace-bin/completers/neomutt_completer/cmd"
	netcat "github.com/rsteube/carapace-bin/completers/netcat_completer/cmd"
	newman "github.com/rsteube/carapace-bin/completers/newman_completer/cmd"
	newrelic "github.com/rsteube/carapace-bin/completers/newrelic_completer/cmd"
	nfpm "github.com/rsteube/carapace-bin/completers/nfpm_completer/cmd"
	ng "github.com/rsteube/carapace-bin/completers/ng_completer/cmd"
	nice "github.com/rsteube/carapace-bin/completers/nice_completer/cmd"
	nilaway "github.com/rsteube/carapace-bin/completers/nilaway_completer/cmd"
	nix_build "github.com/rsteube/carapace-bin/completers/nix-build_completer/cmd"
	nix_channel "github.com/rsteube/carapace-bin/completers/nix-channel_completer/cmd"
	nix_instantiate "github.com/rsteube/carapace-bin/completers/nix-instantiate_completer/cmd"
	nix_shell "github.com/rsteube/carapace-bin/completers/nix-shell_completer/cmd"
	nix "github.com/rsteube/carapace-bin/completers/nix_completer/cmd"
	nl "github.com/rsteube/carapace-bin/completers/nl_completer/cmd"
	nmcli "github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd"
	node "github.com/rsteube/carapace-bin/completers/node_completer/cmd"
	nohup "github.com/rsteube/carapace-bin/completers/nohup_completer/cmd"
	nomad "github.com/rsteube/carapace-bin/completers/nomad_completer/cmd"
	npm "github.com/rsteube/carapace-bin/completers/npm_completer/cmd"
	ntpd "github.com/rsteube/carapace-bin/completers/ntpd_completer/cmd"
	nu "github.com/rsteube/carapace-bin/completers/nu_completer/cmd"
	nvim "github.com/rsteube/carapace-bin/completers/nvim_completer/cmd"
	od "github.com/rsteube/carapace-bin/completers/od_completer/cmd"
	openscad "github.com/rsteube/carapace-bin/completers/openscad_completer/cmd"
	optipng "github.com/rsteube/carapace-bin/completers/optipng_completer/cmd"
	packer "github.com/rsteube/carapace-bin/completers/packer_completer/cmd"
	pacman_conf "github.com/rsteube/carapace-bin/completers/pacman-conf_completer/cmd"
	pacman_db_upgrade "github.com/rsteube/carapace-bin/completers/pacman-db-upgrade_completer/cmd"
	pacman_key "github.com/rsteube/carapace-bin/completers/pacman-key_completer/cmd"
	pacman_mirrors "github.com/rsteube/carapace-bin/completers/pacman-mirrors_completer/cmd"
	pacman "github.com/rsteube/carapace-bin/completers/pacman_completer/cmd"
	palemoon "github.com/rsteube/carapace-bin/completers/palemoon_completer/cmd"
	pamac "github.com/rsteube/carapace-bin/completers/pamac_completer/cmd"
	pandoc "github.com/rsteube/carapace-bin/completers/pandoc_completer/cmd"
	paru "github.com/rsteube/carapace-bin/completers/paru_completer/cmd"
	pass "github.com/rsteube/carapace-bin/completers/pass_completer/cmd"
	passwd "github.com/rsteube/carapace-bin/completers/passwd_completer/cmd"
	paste "github.com/rsteube/carapace-bin/completers/paste_completer/cmd"
	patch "github.com/rsteube/carapace-bin/completers/patch_completer/cmd"
	pathchk "github.com/rsteube/carapace-bin/completers/pathchk_completer/cmd"
	pcmanfm "github.com/rsteube/carapace-bin/completers/pcmanfm_completer/cmd"
	pgrep "github.com/rsteube/carapace-bin/completers/pgrep_completer/cmd"
	picard "github.com/rsteube/carapace-bin/completers/picard_completer/cmd"
	ping "github.com/rsteube/carapace-bin/completers/ping_completer/cmd"
	pinky "github.com/rsteube/carapace-bin/completers/pinky_completer/cmd"
	pip "github.com/rsteube/carapace-bin/completers/pip_completer/cmd"
	pkg "github.com/rsteube/carapace-bin/completers/pkg_completer/cmd"
	pkgsite "github.com/rsteube/carapace-bin/completers/pkgsite_completer/cmd"
	pkill "github.com/rsteube/carapace-bin/completers/pkill_completer/cmd"
	pnpm "github.com/rsteube/carapace-bin/completers/pnpm_completer/cmd"
	podman "github.com/rsteube/carapace-bin/completers/podman_completer/cmd"
	poweroff "github.com/rsteube/carapace-bin/completers/poweroff_completer/cmd"
	powertop "github.com/rsteube/carapace-bin/completers/powertop_completer/cmd"
	pprof "github.com/rsteube/carapace-bin/completers/pprof_completer/cmd"
	pr "github.com/rsteube/carapace-bin/completers/pr_completer/cmd"
	present "github.com/rsteube/carapace-bin/completers/present_completer/cmd"
	prettybat "github.com/rsteube/carapace-bin/completers/prettybat_completer/cmd"
	prettyping "github.com/rsteube/carapace-bin/completers/prettyping_completer/cmd"
	printenv "github.com/rsteube/carapace-bin/completers/printenv_completer/cmd"
	ps "github.com/rsteube/carapace-bin/completers/ps_completer/cmd"
	ptx "github.com/rsteube/carapace-bin/completers/ptx_completer/cmd"
	pulumi "github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd"
	pwait "github.com/rsteube/carapace-bin/completers/pwait_completer/cmd"
	pwd "github.com/rsteube/carapace-bin/completers/pwd_completer/cmd"
	python "github.com/rsteube/carapace-bin/completers/python_completer/cmd"
	qmk "github.com/rsteube/carapace-bin/completers/qmk_completer/cmd"
	qrencode "github.com/rsteube/carapace-bin/completers/qrencode_completer/cmd"
	qutebrowser "github.com/rsteube/carapace-bin/completers/qutebrowser_completer/cmd"
	ranger "github.com/rsteube/carapace-bin/completers/ranger_completer/cmd"
	readlink "github.com/rsteube/carapace-bin/completers/readlink_completer/cmd"
	reboot "github.com/rsteube/carapace-bin/completers/reboot_completer/cmd"
	redis_cli "github.com/rsteube/carapace-bin/completers/redis-cli_completer/cmd"
	restic "github.com/rsteube/carapace-bin/completers/restic_completer/cmd"
	resume_cli "github.com/rsteube/carapace-bin/completers/resume-cli_completer/cmd"
	rg "github.com/rsteube/carapace-bin/completers/rg_completer/cmd"
	rifle "github.com/rsteube/carapace-bin/completers/rifle_completer/cmd"
	rm "github.com/rsteube/carapace-bin/completers/rm_completer/cmd"
	rmdir "github.com/rsteube/carapace-bin/completers/rmdir_completer/cmd"
	rmmod "github.com/rsteube/carapace-bin/completers/rmmod_completer/cmd"
	rsync "github.com/rsteube/carapace-bin/completers/rsync_completer/cmd"
	rust_analyzer "github.com/rsteube/carapace-bin/completers/rust-analyzer_completer/cmd"
	rustc "github.com/rsteube/carapace-bin/completers/rustc_completer/cmd"
	rustdoc "github.com/rsteube/carapace-bin/completers/rustdoc_completer/cmd"
	rustup "github.com/rsteube/carapace-bin/completers/rustup_completer/cmd"
	scp "github.com/rsteube/carapace-bin/completers/scp_completer/cmd"
	scrot "github.com/rsteube/carapace-bin/completers/scrot_completer/cmd"
	sdkmanager "github.com/rsteube/carapace-bin/completers/sdkmanager_completer/cmd"
	sed "github.com/rsteube/carapace-bin/completers/sed_completer/cmd"
	semver "github.com/rsteube/carapace-bin/completers/semver_completer/cmd"
	seq "github.com/rsteube/carapace-bin/completers/seq_completer/cmd"
	set_env "github.com/rsteube/carapace-bin/completers/set-env_completer/cmd"
	sftp "github.com/rsteube/carapace-bin/completers/sftp_completer/cmd"
	sha1sum "github.com/rsteube/carapace-bin/completers/sha1sum_completer/cmd"
	sha256sum "github.com/rsteube/carapace-bin/completers/sha256sum_completer/cmd"
	showkey "github.com/rsteube/carapace-bin/completers/showkey_completer/cmd"
	shred "github.com/rsteube/carapace-bin/completers/shred_completer/cmd"
	shutdown "github.com/rsteube/carapace-bin/completers/shutdown_completer/cmd"
	sleep "github.com/rsteube/carapace-bin/completers/sleep_completer/cmd"
	slides "github.com/rsteube/carapace-bin/completers/slides_completer/cmd"
	soft "github.com/rsteube/carapace-bin/completers/soft_completer/cmd"
	sort "github.com/rsteube/carapace-bin/completers/sort_completer/cmd"
	speedtest_cli "github.com/rsteube/carapace-bin/completers/speedtest-cli_completer/cmd"
	split "github.com/rsteube/carapace-bin/completers/split_completer/cmd"
	ssh_agent "github.com/rsteube/carapace-bin/completers/ssh-agent_completer/cmd"
	ssh_copy_id "github.com/rsteube/carapace-bin/completers/ssh-copy-id_completer/cmd"
	ssh_keygen "github.com/rsteube/carapace-bin/completers/ssh-keygen_completer/cmd"
	ssh "github.com/rsteube/carapace-bin/completers/ssh_completer/cmd"
	st "github.com/rsteube/carapace-bin/completers/st_completer/cmd"
	starship "github.com/rsteube/carapace-bin/completers/starship_completer/cmd"
	stat "github.com/rsteube/carapace-bin/completers/stat_completer/cmd"
	staticcheck "github.com/rsteube/carapace-bin/completers/staticcheck_completer/cmd"
	strings "github.com/rsteube/carapace-bin/completers/strings_completer/cmd"
	stty "github.com/rsteube/carapace-bin/completers/stty_completer/cmd"
	su "github.com/rsteube/carapace-bin/completers/su_completer/cmd"
	sudo "github.com/rsteube/carapace-bin/completers/sudo_completer/cmd"
	sudoedit "github.com/rsteube/carapace-bin/completers/sudoedit_completer/cmd"
	sudoreplay "github.com/rsteube/carapace-bin/completers/sudoreplay_completer/cmd"
	sulogin "github.com/rsteube/carapace-bin/completers/sulogin_completer/cmd"
	sum "github.com/rsteube/carapace-bin/completers/sum_completer/cmd"
	supervisorctl "github.com/rsteube/carapace-bin/completers/supervisorctl_completer/cmd"
	supervisord "github.com/rsteube/carapace-bin/completers/supervisord_completer/cmd"
	svg_term "github.com/rsteube/carapace-bin/completers/svg-term_completer/cmd"
	svgcleaner "github.com/rsteube/carapace-bin/completers/svgcleaner_completer/cmd"
	sway "github.com/rsteube/carapace-bin/completers/sway_completer/cmd"
	swaybar "github.com/rsteube/carapace-bin/completers/swaybar_completer/cmd"
	swaybg "github.com/rsteube/carapace-bin/completers/swaybg_completer/cmd"
	swayidle "github.com/rsteube/carapace-bin/completers/swayidle_completer/cmd"
	swaylock "github.com/rsteube/carapace-bin/completers/swaylock_completer/cmd"
	swaymsg "github.com/rsteube/carapace-bin/completers/swaymsg_completer/cmd"
	swaynag "github.com/rsteube/carapace-bin/completers/swaynag_completer/cmd"
	syft "github.com/rsteube/carapace-bin/completers/syft_completer/cmd"
	sync "github.com/rsteube/carapace-bin/completers/sync_completer/cmd"
	sysctl "github.com/rsteube/carapace-bin/completers/sysctl_completer/cmd"
	systemctl "github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd"
	tac "github.com/rsteube/carapace-bin/completers/tac_completer/cmd"
	tail "github.com/rsteube/carapace-bin/completers/tail_completer/cmd"
	tar "github.com/rsteube/carapace-bin/completers/tar_completer/cmd"
	task "github.com/rsteube/carapace-bin/completers/task_completer/cmd"
	tea "github.com/rsteube/carapace-bin/completers/tea_completer/cmd"
	tee "github.com/rsteube/carapace-bin/completers/tee_completer/cmd"
	telnet "github.com/rsteube/carapace-bin/completers/telnet_completer/cmd"
	templ "github.com/rsteube/carapace-bin/completers/templ_completer/cmd"
	termux_apt_repo "github.com/rsteube/carapace-bin/completers/termux-apt-repo_completer/cmd"
	terraform_ls "github.com/rsteube/carapace-bin/completers/terraform-ls_completer/cmd"
	terraform "github.com/rsteube/carapace-bin/completers/terraform_completer/cmd"
	terragrunt "github.com/rsteube/carapace-bin/completers/terragrunt_completer/cmd"
	terramate "github.com/rsteube/carapace-bin/completers/terramate_completer/cmd"
	tesseract "github.com/rsteube/carapace-bin/completers/tesseract_completer/cmd"
	tig "github.com/rsteube/carapace-bin/completers/tig_completer/cmd"
	tinygo "github.com/rsteube/carapace-bin/completers/tinygo_completer/cmd"
	tldr "github.com/rsteube/carapace-bin/completers/tldr_completer/cmd"
	tmate "github.com/rsteube/carapace-bin/completers/tmate_completer/cmd"
	tmux "github.com/rsteube/carapace-bin/completers/tmux_completer/cmd"
	tofu "github.com/rsteube/carapace-bin/completers/tofu_completer/cmd"
	toit_lsp "github.com/rsteube/carapace-bin/completers/toit.lsp_completer/cmd"
	toit_pkg "github.com/rsteube/carapace-bin/completers/toit.pkg_completer/cmd"
	top "github.com/rsteube/carapace-bin/completers/top_completer/cmd"
	tor_browser "github.com/rsteube/carapace-bin/completers/tor-browser_completer/cmd"
	tor_gencert "github.com/rsteube/carapace-bin/completers/tor-gencert_completer/cmd"
	tor_print_ed_signing_cert "github.com/rsteube/carapace-bin/completers/tor-print-ed-signing-cert_completer/cmd"
	tor_resolve "github.com/rsteube/carapace-bin/completers/tor-resolve_completer/cmd"
	torsocks "github.com/rsteube/carapace-bin/completers/torsocks_completer/cmd"
	touch "github.com/rsteube/carapace-bin/completers/touch_completer/cmd"
	tr "github.com/rsteube/carapace-bin/completers/tr_completer/cmd"
	traefik "github.com/rsteube/carapace-bin/completers/traefik_completer/cmd"
	tree "github.com/rsteube/carapace-bin/completers/tree_completer/cmd"
	truncate "github.com/rsteube/carapace-bin/completers/truncate_completer/cmd"
	ts "github.com/rsteube/carapace-bin/completers/ts_completer/cmd"
	tsc "github.com/rsteube/carapace-bin/completers/tsc_completer/cmd"
	tsh "github.com/rsteube/carapace-bin/completers/tsh_completer/cmd"
	tshark "github.com/rsteube/carapace-bin/completers/tshark_completer/cmd"
	tsort "github.com/rsteube/carapace-bin/completers/tsort_completer/cmd"
	tty "github.com/rsteube/carapace-bin/completers/tty_completer/cmd"
	ttyd "github.com/rsteube/carapace-bin/completers/ttyd_completer/cmd"
	turbo "github.com/rsteube/carapace-bin/completers/turbo_completer/cmd"
	umount "github.com/rsteube/carapace-bin/completers/umount_completer/cmd"
	uname "github.com/rsteube/carapace-bin/completers/uname_completer/cmd"
	unbrotli "github.com/rsteube/carapace-bin/completers/unbrotli_completer/cmd"
	unexpand "github.com/rsteube/carapace-bin/completers/unexpand_completer/cmd"
	uniq "github.com/rsteube/carapace-bin/completers/uniq_completer/cmd"
	unlink "github.com/rsteube/carapace-bin/completers/unlink_completer/cmd"
	unlzma "github.com/rsteube/carapace-bin/completers/unlzma_completer/cmd"
	unset_env "github.com/rsteube/carapace-bin/completers/unset-env_completer/cmd"
	unxz "github.com/rsteube/carapace-bin/completers/unxz_completer/cmd"
	unzip "github.com/rsteube/carapace-bin/completers/unzip_completer/cmd"
	upower "github.com/rsteube/carapace-bin/completers/upower_completer/cmd"
	uptime "github.com/rsteube/carapace-bin/completers/uptime_completer/cmd"
	upx "github.com/rsteube/carapace-bin/completers/upx_completer/cmd"
	useradd "github.com/rsteube/carapace-bin/completers/useradd_completer/cmd"
	userdel "github.com/rsteube/carapace-bin/completers/userdel_completer/cmd"
	usermod "github.com/rsteube/carapace-bin/completers/usermod_completer/cmd"
	users "github.com/rsteube/carapace-bin/completers/users_completer/cmd"
	vagrant "github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd"
	vault "github.com/rsteube/carapace-bin/completers/vault_completer/cmd"
	vdir "github.com/rsteube/carapace-bin/completers/vdir_completer/cmd"
	vercel "github.com/rsteube/carapace-bin/completers/vercel_completer/cmd"
	vhs "github.com/rsteube/carapace-bin/completers/vhs_completer/cmd"
	vi "github.com/rsteube/carapace-bin/completers/vi_completer/cmd"
	viewnior "github.com/rsteube/carapace-bin/completers/viewnior_completer/cmd"
	visudo "github.com/rsteube/carapace-bin/completers/visudo_completer/cmd"
	viu "github.com/rsteube/carapace-bin/completers/viu_completer/cmd"
	vivid "github.com/rsteube/carapace-bin/completers/vivid_completer/cmd"
	vlc "github.com/rsteube/carapace-bin/completers/vlc_completer/cmd"
	volta "github.com/rsteube/carapace-bin/completers/volta_completer/cmd"
	w "github.com/rsteube/carapace-bin/completers/w_completer/cmd"
	watch "github.com/rsteube/carapace-bin/completers/watch_completer/cmd"
	watchexec "github.com/rsteube/carapace-bin/completers/watchexec_completer/cmd"
	watchgnupg "github.com/rsteube/carapace-bin/completers/watchgnupg_completer/cmd"
	waypoint "github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd"
	wc "github.com/rsteube/carapace-bin/completers/wc_completer/cmd"
	wezterm "github.com/rsteube/carapace-bin/completers/wezterm_completer/cmd"
	wget "github.com/rsteube/carapace-bin/completers/wget_completer/cmd"
	whereis "github.com/rsteube/carapace-bin/completers/whereis_completer/cmd"
	which "github.com/rsteube/carapace-bin/completers/which_completer/cmd"
	who "github.com/rsteube/carapace-bin/completers/who_completer/cmd"
	whoami "github.com/rsteube/carapace-bin/completers/whoami_completer/cmd"
	wine "github.com/rsteube/carapace-bin/completers/wine_completer/cmd"
	wineboot "github.com/rsteube/carapace-bin/completers/wineboot_completer/cmd"
	winepath "github.com/rsteube/carapace-bin/completers/winepath_completer/cmd"
	wineserver "github.com/rsteube/carapace-bin/completers/wineserver_completer/cmd"
	winetricks "github.com/rsteube/carapace-bin/completers/winetricks_completer/cmd"
	wire "github.com/rsteube/carapace-bin/completers/wire_completer/cmd"
	wireshark "github.com/rsteube/carapace-bin/completers/wireshark_completer/cmd"
	wishlist "github.com/rsteube/carapace-bin/completers/wishlist_completer/cmd"
	woeusb "github.com/rsteube/carapace-bin/completers/woeusb_completer/cmd"
	xargs "github.com/rsteube/carapace-bin/completers/xargs_completer/cmd"
	xbacklight "github.com/rsteube/carapace-bin/completers/xbacklight_completer/cmd"
	xdotool "github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd"
	xonsh "github.com/rsteube/carapace-bin/completers/xonsh_completer/cmd"
	xz "github.com/rsteube/carapace-bin/completers/xz_completer/cmd"
	xzcat "github.com/rsteube/carapace-bin/completers/xzcat_completer/cmd"
	yarn "github.com/rsteube/carapace-bin/completers/yarn_completer/cmd"
	yay "github.com/rsteube/carapace-bin/completers/yay_completer/cmd"
	yes "github.com/rsteube/carapace-bin/completers/yes_completer/cmd"
	yj "github.com/rsteube/carapace-bin/completers/yj_completer/cmd"
	youtube_dl "github.com/rsteube/carapace-bin/completers/youtube-dl_completer/cmd"
	yt_dlp "github.com/rsteube/carapace-bin/completers/yt-dlp_completer/cmd"
	zathura "github.com/rsteube/carapace-bin/completers/zathura_completer/cmd"
	zcat "github.com/rsteube/carapace-bin/completers/zcat_completer/cmd"
	zip "github.com/rsteube/carapace-bin/completers/zip_completer/cmd"
	zoxide "github.com/rsteube/carapace-bin/completers/zoxide_completer/cmd"
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
