// Copyright 2017 clair authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpm

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coreos/clair/database"
	"github.com/coreos/clair/ext/featurefmt"
	"github.com/coreos/clair/ext/versionfmt/rpm"
)

var expectedBigCaseInfo = []database.Feature{
	{"publicsuffix-list-dafsa", "20180514-1.fc28", "publicsuffix-list", "20180514-1.fc28", rpm.ParserName},
	{"libreport-filesystem", "2.9.5-1.fc28", "libreport", "2.9.5-1.fc28", rpm.ParserName},
	{"fedora-gpg-keys", "28-5", "fedora-repos", "28-5", rpm.ParserName},
	{"fedora-release", "28-2", "fedora-release", "28-2", rpm.ParserName},
	{"filesystem", "3.8-2.fc28", "filesystem", "3.8-2.fc28", rpm.ParserName},
	{"tzdata", "2018e-1.fc28", "tzdata", "2018e-1.fc28", rpm.ParserName},
	{"pcre2", "10.31-10.fc28", "pcre2", "10.31-10.fc28", rpm.ParserName},
	{"glibc-minimal-langpack", "2.27-32.fc28", "glibc", "2.27-32.fc28", rpm.ParserName},
	{"glibc-common", "2.27-32.fc28", "glibc", "2.27-32.fc28", rpm.ParserName},
	{"bash", "4.4.23-1.fc28", "bash", "4.4.23-1.fc28", rpm.ParserName},
	{"zlib", "1.2.11-8.fc28", "zlib", "1.2.11-8.fc28", rpm.ParserName},
	{"bzip2-libs", "1.0.6-26.fc28", "bzip2", "1.0.6-26.fc28", rpm.ParserName},
	{"libcap", "2.25-9.fc28", "libcap", "2.25-9.fc28", rpm.ParserName},
	{"libgpg-error", "1.31-1.fc28", "libgpg-error", "1.31-1.fc28", rpm.ParserName},
	{"libzstd", "1.3.5-1.fc28", "zstd", "1.3.5-1.fc28", rpm.ParserName},
	{"expat", "2.2.5-3.fc28", "expat", "2.2.5-3.fc28", rpm.ParserName},
	{"nss-util", "3.38.0-1.0.fc28", "nss-util", "3.38.0-1.0.fc28", rpm.ParserName},
	{"libcom_err", "1.44.2-0.fc28", "e2fsprogs", "1.44.2-0.fc28", rpm.ParserName},
	{"libffi", "3.1-16.fc28", "libffi", "3.1-16.fc28", rpm.ParserName},
	{"libgcrypt", "1.8.3-1.fc28", "libgcrypt", "1.8.3-1.fc28", rpm.ParserName},
	{"libxml2", "2.9.8-4.fc28", "libxml2", "2.9.8-4.fc28", rpm.ParserName},
	{"libacl", "2.2.53-1.fc28", "acl", "2.2.53-1.fc28", rpm.ParserName},
	{"sed", "4.5-1.fc28", "sed", "4.5-1.fc28", rpm.ParserName},
	{"libmount", "2.32.1-1.fc28", "util-linux", "2.32.1-1.fc28", rpm.ParserName},
	{"p11-kit", "0.23.12-1.fc28", "p11-kit", "0.23.12-1.fc28", rpm.ParserName},
	{"libidn2", "2.0.5-1.fc28", "libidn2", "2.0.5-1.fc28", rpm.ParserName},
	{"libcap-ng", "0.7.9-4.fc28", "libcap-ng", "0.7.9-4.fc28", rpm.ParserName},
	{"lz4-libs", "1.8.1.2-4.fc28", "lz4", "1.8.1.2-4.fc28", rpm.ParserName},
	{"libassuan", "2.5.1-3.fc28", "libassuan", "2.5.1-3.fc28", rpm.ParserName},
	{"keyutils-libs", "1.5.10-6.fc28", "keyutils", "1.5.10-6.fc28", rpm.ParserName},
	{"glib2", "2.56.1-4.fc28", "glib2", "2.56.1-4.fc28", rpm.ParserName},
	{"systemd-libs", "238-9.git0e0aa59.fc28", "systemd", "238-9.git0e0aa59.fc28", rpm.ParserName},
	{"dbus-libs", "1:1.12.10-1.fc28", "dbus", "1.12.10-1.fc28", rpm.ParserName},
	{"libtasn1", "4.13-2.fc28", "libtasn1", "4.13-2.fc28", rpm.ParserName},
	{"ca-certificates", "2018.2.24-1.0.fc28", "ca-certificates", "2018.2.24-1.0.fc28", rpm.ParserName},
	{"libarchive", "3.3.1-4.fc28", "libarchive", "3.3.1-4.fc28", rpm.ParserName},
	{"openssl", "1:1.1.0h-3.fc28", "openssl", "1.1.0h-3.fc28", rpm.ParserName},
	{"libusbx", "1.0.22-1.fc28", "libusbx", "1.0.22-1.fc28", rpm.ParserName},
	{"libsemanage", "2.8-2.fc28", "libsemanage", "2.8-2.fc28", rpm.ParserName},
	{"libutempter", "1.1.6-14.fc28", "libutempter", "1.1.6-14.fc28", rpm.ParserName},
	{"mpfr", "3.1.6-1.fc28", "mpfr", "3.1.6-1.fc28", rpm.ParserName},
	{"gnutls", "3.6.3-4.fc28", "gnutls", "3.6.3-4.fc28", rpm.ParserName},
	{"gzip", "1.9-3.fc28", "gzip", "1.9-3.fc28", rpm.ParserName},
	{"acl", "2.2.53-1.fc28", "acl", "2.2.53-1.fc28", rpm.ParserName},
	{"nss-softokn-freebl", "3.38.0-1.0.fc28", "nss-softokn", "3.38.0-1.0.fc28", rpm.ParserName},
	{"nss", "3.38.0-1.0.fc28", "nss", "3.38.0-1.0.fc28", rpm.ParserName},
	{"libmetalink", "0.1.3-6.fc28", "libmetalink", "0.1.3-6.fc28", rpm.ParserName},
	{"libdb-utils", "5.3.28-30.fc28", "libdb", "5.3.28-30.fc28", rpm.ParserName},
	{"file-libs", "5.33-7.fc28", "file", "5.33-7.fc28", rpm.ParserName},
	{"libsss_idmap", "1.16.3-2.fc28", "sssd", "1.16.3-2.fc28", rpm.ParserName},
	{"libsigsegv", "2.11-5.fc28", "libsigsegv", "2.11-5.fc28", rpm.ParserName},
	{"krb5-libs", "1.16.1-13.fc28", "krb5", "1.16.1-13.fc28", rpm.ParserName},
	{"libnsl2", "1.2.0-2.20180605git4a062cf.fc28", "libnsl2", "1.2.0-2.20180605git4a062cf.fc28", rpm.ParserName},
	{"python3-pip", "9.0.3-2.fc28", "python-pip", "9.0.3-2.fc28", rpm.ParserName},
	{"python3", "3.6.6-1.fc28", "python3", "3.6.6-1.fc28", rpm.ParserName},
	{"pam", "1.3.1-1.fc28", "pam", "1.3.1-1.fc28", rpm.ParserName},
	{"python3-gobject-base", "3.28.3-1.fc28", "pygobject3", "3.28.3-1.fc28", rpm.ParserName},
	{"python3-smartcols", "0.3.0-2.fc28", "python-smartcols", "0.3.0-2.fc28", rpm.ParserName},
	{"python3-iniparse", "0.4-30.fc28", "python-iniparse", "0.4-30.fc28", rpm.ParserName},
	{"openldap", "2.4.46-3.fc28", "openldap", "2.4.46-3.fc28", rpm.ParserName},
	{"libseccomp", "2.3.3-2.fc28", "libseccomp", "2.3.3-2.fc28", rpm.ParserName},
	{"npth", "1.5-4.fc28", "npth", "1.5-4.fc28", rpm.ParserName},
	{"gpgme", "1.10.0-4.fc28", "gpgme", "1.10.0-4.fc28", rpm.ParserName},
	{"json-c", "0.13.1-2.fc28", "json-c", "0.13.1-2.fc28", rpm.ParserName},
	{"libyaml", "0.1.7-5.fc28", "libyaml", "0.1.7-5.fc28", rpm.ParserName},
	{"libpkgconf", "1.4.2-1.fc28", "pkgconf", "1.4.2-1.fc28", rpm.ParserName},
	{"pkgconf-pkg-config", "1.4.2-1.fc28", "pkgconf", "1.4.2-1.fc28", rpm.ParserName},
	{"iptables-libs", "1.6.2-3.fc28", "iptables", "1.6.2-3.fc28", rpm.ParserName},
	{"device-mapper-libs", "1.02.146-5.fc28", "lvm2", "2.02.177-5.fc28", rpm.ParserName},
	{"systemd-pam", "238-9.git0e0aa59.fc28", "systemd", "238-9.git0e0aa59.fc28", rpm.ParserName},
	{"systemd", "238-9.git0e0aa59.fc28", "systemd", "238-9.git0e0aa59.fc28", rpm.ParserName},
	{"elfutils-default-yama-scope", "0.173-1.fc28", "elfutils", "0.173-1.fc28", rpm.ParserName},
	{"libcurl", "7.59.0-6.fc28", "curl", "7.59.0-6.fc28", rpm.ParserName},
	{"python3-librepo", "1.8.1-7.fc28", "librepo", "1.8.1-7.fc28", rpm.ParserName},
	{"rpm-plugin-selinux", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"rpm", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"libdnf", "0.11.1-3.fc28", "libdnf", "0.11.1-3.fc28", rpm.ParserName},
	{"rpm-build-libs", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"python3-rpm", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"dnf", "2.7.5-12.fc28", "dnf", "2.7.5-12.fc28", rpm.ParserName},
	{"deltarpm", "3.6-25.fc28", "deltarpm", "3.6-25.fc28", rpm.ParserName},
	{"sssd-client", "1.16.3-2.fc28", "sssd", "1.16.3-2.fc28", rpm.ParserName},
	{"cracklib-dicts", "2.9.6-13.fc28", "cracklib", "2.9.6-13.fc28", rpm.ParserName},
	{"tar", "2:1.30-3.fc28", "tar", "1.30-3.fc28", rpm.ParserName},
	{"diffutils", "3.6-4.fc28", "diffutils", "3.6-4.fc28", rpm.ParserName},
	{"langpacks-en", "1.0-12.fc28", "langpacks", "1.0-12.fc28", rpm.ParserName},
	{"libgcc", "8.1.1-5.fc28", "gcc", "8.1.1-5.fc28", rpm.ParserName},
	{"pkgconf-m4", "1.4.2-1.fc28", "pkgconf", "1.4.2-1.fc28", rpm.ParserName},
	{"dnf-conf", "2.7.5-12.fc28", "dnf", "2.7.5-12.fc28", rpm.ParserName},
	{"fedora-repos", "28-5", "fedora-repos", "28-5", rpm.ParserName},
	{"setup", "2.11.4-1.fc28", "setup", "2.11.4-1.fc28", rpm.ParserName},
	{"basesystem", "11-5.fc28", "basesystem", "11-5.fc28", rpm.ParserName},
	{"ncurses-base", "6.1-5.20180224.fc28", "ncurses", "6.1-5.20180224.fc28", rpm.ParserName},
	{"libselinux", "2.8-1.fc28", "libselinux", "2.8-1.fc28", rpm.ParserName},
	{"ncurses-libs", "6.1-5.20180224.fc28", "ncurses", "6.1-5.20180224.fc28", rpm.ParserName},
	{"glibc", "2.27-32.fc28", "glibc", "2.27-32.fc28", rpm.ParserName},
	{"libsepol", "2.8-1.fc28", "libsepol", "2.8-1.fc28", rpm.ParserName},
	{"xz-libs", "5.2.4-2.fc28", "xz", "5.2.4-2.fc28", rpm.ParserName},
	{"info", "6.5-4.fc28", "texinfo", "6.5-4.fc28", rpm.ParserName},
	{"libdb", "5.3.28-30.fc28", "libdb", "5.3.28-30.fc28", rpm.ParserName},
	{"elfutils-libelf", "0.173-1.fc28", "elfutils", "0.173-1.fc28", rpm.ParserName},
	{"popt", "1.16-14.fc28", "popt", "1.16-14.fc28", rpm.ParserName},
	{"nspr", "4.19.0-1.fc28", "nspr", "4.19.0-1.fc28", rpm.ParserName},
	{"libxcrypt", "4.1.2-1.fc28", "libxcrypt", "4.1.2-1.fc28", rpm.ParserName},
	{"lua-libs", "5.3.4-10.fc28", "lua", "5.3.4-10.fc28", rpm.ParserName},
	{"libuuid", "2.32.1-1.fc28", "util-linux", "2.32.1-1.fc28", rpm.ParserName},
	{"readline", "7.0-11.fc28", "readline", "7.0-11.fc28", rpm.ParserName},
	{"libattr", "2.4.48-3.fc28", "attr", "2.4.48-3.fc28", rpm.ParserName},
	{"coreutils-single", "8.29-7.fc28", "coreutils", "8.29-7.fc28", rpm.ParserName},
	{"libblkid", "2.32.1-1.fc28", "util-linux", "2.32.1-1.fc28", rpm.ParserName},
	{"gmp", "1:6.1.2-7.fc28", "gmp", "6.1.2-7.fc28", rpm.ParserName},
	{"libunistring", "0.9.10-1.fc28", "libunistring", "0.9.10-1.fc28", rpm.ParserName},
	{"sqlite-libs", "3.22.0-4.fc28", "sqlite", "3.22.0-4.fc28", rpm.ParserName},
	{"audit-libs", "2.8.4-2.fc28", "audit", "2.8.4-2.fc28", rpm.ParserName},
	{"chkconfig", "1.10-4.fc28", "chkconfig", "1.10-4.fc28", rpm.ParserName},
	{"libsmartcols", "2.32.1-1.fc28", "util-linux", "2.32.1-1.fc28", rpm.ParserName},
	{"pcre", "8.42-3.fc28", "pcre", "8.42-3.fc28", rpm.ParserName},
	{"grep", "3.1-5.fc28", "grep", "3.1-5.fc28", rpm.ParserName},
	{"crypto-policies", "20180425-5.git6ad4018.fc28", "crypto-policies", "20180425-5.git6ad4018.fc28", rpm.ParserName},
	{"gdbm-libs", "1:1.14.1-4.fc28", "gdbm", "1.14.1-4.fc28", rpm.ParserName},
	{"p11-kit-trust", "0.23.12-1.fc28", "p11-kit", "0.23.12-1.fc28", rpm.ParserName},
	{"openssl-libs", "1:1.1.0h-3.fc28", "openssl", "1.1.0h-3.fc28", rpm.ParserName},
	{"ima-evm-utils", "1.1-2.fc28", "ima-evm-utils", "1.1-2.fc28", rpm.ParserName},
	{"gdbm", "1:1.14.1-4.fc28", "gdbm", "1.14.1-4.fc28", rpm.ParserName},
	{"gobject-introspection", "1.56.1-1.fc28", "gobject-introspection", "1.56.1-1.fc28", rpm.ParserName},
	{"shadow-utils", "2:4.6-1.fc28", "shadow-utils", "4.6-1.fc28", rpm.ParserName},
	{"libpsl", "0.20.2-2.fc28", "libpsl", "0.20.2-2.fc28", rpm.ParserName},
	{"nettle", "3.4-2.fc28", "nettle", "3.4-2.fc28", rpm.ParserName},
	{"libfdisk", "2.32.1-1.fc28", "util-linux", "2.32.1-1.fc28", rpm.ParserName},
	{"cracklib", "2.9.6-13.fc28", "cracklib", "2.9.6-13.fc28", rpm.ParserName},
	{"libcomps", "0.1.8-11.fc28", "libcomps", "0.1.8-11.fc28", rpm.ParserName},
	{"nss-softokn", "3.38.0-1.0.fc28", "nss-softokn", "3.38.0-1.0.fc28", rpm.ParserName},
	{"nss-sysinit", "3.38.0-1.0.fc28", "nss", "3.38.0-1.0.fc28", rpm.ParserName},
	{"libksba", "1.3.5-7.fc28", "libksba", "1.3.5-7.fc28", rpm.ParserName},
	{"kmod-libs", "25-2.fc28", "kmod", "25-2.fc28", rpm.ParserName},
	{"libsss_nss_idmap", "1.16.3-2.fc28", "sssd", "1.16.3-2.fc28", rpm.ParserName},
	{"libverto", "0.3.0-5.fc28", "libverto", "0.3.0-5.fc28", rpm.ParserName},
	{"gawk", "4.2.1-1.fc28", "gawk", "4.2.1-1.fc28", rpm.ParserName},
	{"libtirpc", "1.0.3-3.rc2.fc28", "libtirpc", "1.0.3-3.rc2.fc28", rpm.ParserName},
	{"python3-libs", "3.6.6-1.fc28", "python3", "3.6.6-1.fc28", rpm.ParserName},
	{"python3-setuptools", "39.2.0-6.fc28", "python-setuptools", "39.2.0-6.fc28", rpm.ParserName},
	{"libpwquality", "1.4.0-7.fc28", "libpwquality", "1.4.0-7.fc28", rpm.ParserName},
	{"util-linux", "2.32.1-1.fc28", "util-linux", "2.32.1-1.fc28", rpm.ParserName},
	{"python3-libcomps", "0.1.8-11.fc28", "libcomps", "0.1.8-11.fc28", rpm.ParserName},
	{"python3-six", "1.11.0-3.fc28", "python-six", "1.11.0-3.fc28", rpm.ParserName},
	{"cyrus-sasl-lib", "2.1.27-0.2rc7.fc28", "cyrus-sasl", "2.1.27-0.2rc7.fc28", rpm.ParserName},
	{"libssh", "0.8.2-1.fc28", "libssh", "0.8.2-1.fc28", rpm.ParserName},
	{"qrencode-libs", "3.4.4-5.fc28", "qrencode", "3.4.4-5.fc28", rpm.ParserName},
	{"gnupg2", "2.2.8-1.fc28", "gnupg2", "2.2.8-1.fc28", rpm.ParserName},
	{"python3-gpg", "1.10.0-4.fc28", "gpgme", "1.10.0-4.fc28", rpm.ParserName},
	{"libargon2", "20161029-5.fc28", "argon2", "20161029-5.fc28", rpm.ParserName},
	{"libmodulemd", "1.6.2-2.fc28", "libmodulemd", "1.6.2-2.fc28", rpm.ParserName},
	{"pkgconf", "1.4.2-1.fc28", "pkgconf", "1.4.2-1.fc28", rpm.ParserName},
	{"libpcap", "14:1.9.0-1.fc28", "libpcap", "1.9.0-1.fc28", rpm.ParserName},
	{"device-mapper", "1.02.146-5.fc28", "lvm2", "2.02.177-5.fc28", rpm.ParserName},
	{"cryptsetup-libs", "2.0.4-1.fc28", "cryptsetup", "2.0.4-1.fc28", rpm.ParserName},
	{"elfutils-libs", "0.173-1.fc28", "elfutils", "0.173-1.fc28", rpm.ParserName},
	{"dbus", "1:1.12.10-1.fc28", "dbus", "1.12.10-1.fc28", rpm.ParserName},
	{"libnghttp2", "1.32.1-1.fc28", "nghttp2", "1.32.1-1.fc28", rpm.ParserName},
	{"librepo", "1.8.1-7.fc28", "librepo", "1.8.1-7.fc28", rpm.ParserName},
	{"curl", "7.59.0-6.fc28", "curl", "7.59.0-6.fc28", rpm.ParserName},
	{"rpm-libs", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"libsolv", "0.6.35-1.fc28", "libsolv", "0.6.35-1.fc28", rpm.ParserName},
	{"python3-hawkey", "0.11.1-3.fc28", "libdnf", "0.11.1-3.fc28", rpm.ParserName},
	{"rpm-sign-libs", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"python3-dnf", "2.7.5-12.fc28", "dnf", "2.7.5-12.fc28", rpm.ParserName},
	{"dnf-yum", "2.7.5-12.fc28", "dnf", "2.7.5-12.fc28", rpm.ParserName},
	{"rpm-plugin-systemd-inhibit", "4.14.1-9.fc28", "rpm", "4.14.1-9.fc28", rpm.ParserName},
	{"nss-tools", "3.38.0-1.0.fc28", "nss", "3.38.0-1.0.fc28", rpm.ParserName},
	{"openssl-pkcs11", "0.4.8-1.fc28", "openssl-pkcs11", "0.4.8-1.fc28", rpm.ParserName},
	{"vim-minimal", "2:8.1.328-1.fc28", "vim", "8.1.328-1.fc28", rpm.ParserName},
	{"glibc-langpack-en", "2.27-32.fc28", "glibc", "2.27-32.fc28", rpm.ParserName},
	{"rootfiles", "8.1-22.fc28", "rootfiles", "8.1-22.fc28", rpm.ParserName},
}

func TestRpmFeatureDetection(t *testing.T) {
	for _, test := range []featurefmt.TestCase{
		{
			"valid small case",
			map[string]string{"var/lib/rpm/Packages": "rpm/testdata/valid"},
			[]database.Feature{
				{"centos-release", "7-1.1503.el7.centos.2.8", "centos-release", "7-1.1503.el7.centos.2.8", rpm.ParserName},
				{"filesystem", "3.2-18.el7", "filesystem", "3.2-18.el7", rpm.ParserName},
			},
		},
		{
			"valid big case",
			map[string]string{"var/lib/rpm/Packages": "rpm/testdata/valid_big"},
			expectedBigCaseInfo,
		},
	} {
		featurefmt.RunTest(t, test, lister{}, rpm.ParserName)
	}
}

func TestParseSourceRPM(t *testing.T) {
	for _, test := range [...]struct {
		sourceRPM string

		expectedName    string
		expectedVersion string
		expectedErr     string
	}{
		// valid cases
		{"publicsuffix-list-20180514-1.fc28.src.rpm", "publicsuffix-list", "20180514-1.fc28", ""},
		{"libreport-2.9.5-1.fc28.src.rpm", "libreport", "2.9.5-1.fc28", ""},
		{"lua-5.3.4-10.fc28.src.rpm", "lua", "5.3.4-10.fc28", ""},
		{"crypto-policies-20180425-5.git6ad4018.fc28.src.rpm", "crypto-policies", "20180425-5.git6ad4018.fc28", ""},

		// invalid cases
		{"crypto-policies-20180425-5.git6ad4018.fc28.src.dpkg", "", "", "unexpected package type, expect: 'rpm', got: 'dpkg'"},
		{"crypto-policies-20180425-5.git6ad4018.fc28.debian-8.rpm", "", "", "unexpected package architecture, expect: 'src' or 'nosrc', got: 'debian-8'"},
		{"fc28.src.rpm", "", "", "unexpected termination while parsing 'Release Token'"},
		{"...", "", "", "unexpected package type, expect: 'rpm', got: ''"},

		// impossible case
		// This illustrates the limitation of this parser, it will not find the
		// error cased by extra '-' in the intended version/expect token. Based
		// on the documentation, this case should never happen and indicates a
		// corrupted rpm database.
		// actual expected: name="lua", version="5.3.4", release="10.fc-28"
		{"lua-5.3.4-10.fc-28.src.rpm", "lua-5.3.4", "10.fc-28", ""},
	} {
		pkg := database.Feature{}
		err := parseSourceRPM(test.sourceRPM, &pkg)
		if test.expectedErr != "" {
			require.EqualError(t, err, test.expectedErr)
			continue
		}

		require.Nil(t, err)
		require.Equal(t, test.expectedName, pkg.SourceName)
		require.Equal(t, test.expectedVersion, pkg.SourceVersion)
	}
}
