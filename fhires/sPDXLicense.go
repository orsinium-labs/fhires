package fhires

// The license that applies to an Implementation Guide (using an SPDX license Identifiers, or 'not-open-source').
//
// The binding is required but new SPDX license Identifiers are allowed to be used (https://spdx.org/licenses/).
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/spdx-license
type SPDXLicense string

const (
	// Not an open source license.
	SPDXLicenseNotOpenSource SPDXLicense = "not-open-source"
	// BSD Zero Clause License.
	SPDXLicense0BSD SPDXLicense = "0BSD"
	// Attribution Assurance License.
	SPDXLicenseAAL SPDXLicense = "AAL"
	// Abstyles License.
	SPDXLicenseAbstyles SPDXLicense = "Abstyles"
	// Adobe Systems Incorporated Source Code License Agreement.
	SPDXLicenseAdobe2006 SPDXLicense = "Adobe-2006"
	// Adobe Glyph List License.
	SPDXLicenseAdobeGlyph SPDXLicense = "Adobe-Glyph"
	// Amazon Digital Services License.
	SPDXLicenseADSL SPDXLicense = "ADSL"
	// Academic Free License v1.1.
	SPDXLicenseAFL1_1 SPDXLicense = "AFL-1.1"
	// Academic Free License v1.2.
	SPDXLicenseAFL1_2 SPDXLicense = "AFL-1.2"
	// Academic Free License v2.0.
	SPDXLicenseAFL2_0 SPDXLicense = "AFL-2.0"
	// Academic Free License v2.1.
	SPDXLicenseAFL2_1 SPDXLicense = "AFL-2.1"
	// Academic Free License v3.0.
	SPDXLicenseAFL3_0 SPDXLicense = "AFL-3.0"
	// Afmparse License.
	SPDXLicenseAfmparse SPDXLicense = "Afmparse"
	// Affero General Public License v1.0 only.
	SPDXLicenseAGPL1_0Only SPDXLicense = "AGPL-1.0-only"
	// Affero General Public License v1.0 or later.
	SPDXLicenseAGPL1_0OrLater SPDXLicense = "AGPL-1.0-or-later"
	// GNU Affero General Public License v3.0 only.
	SPDXLicenseAGPL3_0Only SPDXLicense = "AGPL-3.0-only"
	// GNU Affero General Public License v3.0 or later.
	SPDXLicenseAGPL3_0OrLater SPDXLicense = "AGPL-3.0-or-later"
	// Aladdin Free Public License.
	SPDXLicenseAladdin SPDXLicense = "Aladdin"
	// AMD's plpa_map.c License.
	SPDXLicenseAMDPLPA SPDXLicense = "AMDPLPA"
	// Apple MIT License.
	SPDXLicenseAML SPDXLicense = "AML"
	// Academy of Motion Picture Arts and Sciences BSD.
	SPDXLicenseAMPAS SPDXLicense = "AMPAS"
	// ANTLR Software Rights Notice.
	SPDXLicenseANTLRPD SPDXLicense = "ANTLR-PD"
	// Apache License 1.0.
	SPDXLicenseApache1_0 SPDXLicense = "Apache-1.0"
	// Apache License 1.1.
	SPDXLicenseApache1_1 SPDXLicense = "Apache-1.1"
	// Apache License 2.0.
	SPDXLicenseApache2_0 SPDXLicense = "Apache-2.0"
	// Adobe Postscript AFM License.
	SPDXLicenseAPAFML SPDXLicense = "APAFML"
	// Adaptive Public License 1.0.
	SPDXLicenseAPL1_0 SPDXLicense = "APL-1.0"
	// Apple Public Source License 1.0.
	SPDXLicenseAPSL1_0 SPDXLicense = "APSL-1.0"
	// Apple Public Source License 1.1.
	SPDXLicenseAPSL1_1 SPDXLicense = "APSL-1.1"
	// Apple Public Source License 1.2.
	SPDXLicenseAPSL1_2 SPDXLicense = "APSL-1.2"
	// Apple Public Source License 2.0.
	SPDXLicenseAPSL2_0 SPDXLicense = "APSL-2.0"
	// Artistic License 1.0 w/clause 8.
	SPDXLicenseArtistic1_0Cl8 SPDXLicense = "Artistic-1.0-cl8"
	// Artistic License 1.0 (Perl).
	SPDXLicenseArtistic1_0Perl SPDXLicense = "Artistic-1.0-Perl"
	// Artistic License 1.0.
	SPDXLicenseArtistic1_0 SPDXLicense = "Artistic-1.0"
	// Artistic License 2.0.
	SPDXLicenseArtistic2_0 SPDXLicense = "Artistic-2.0"
	// Bahyph License.
	SPDXLicenseBahyph SPDXLicense = "Bahyph"
	// Barr License.
	SPDXLicenseBarr SPDXLicense = "Barr"
	// Beerware License.
	SPDXLicenseBeerware SPDXLicense = "Beerware"
	// BitTorrent Open Source License v1.0.
	SPDXLicenseBitTorrent1_0 SPDXLicense = "BitTorrent-1.0"
	// BitTorrent Open Source License v1.1.
	SPDXLicenseBitTorrent1_1 SPDXLicense = "BitTorrent-1.1"
	// Borceux license.
	SPDXLicenseBorceux SPDXLicense = "Borceux"
	// BSD 1-Clause License.
	SPDXLicenseBSD1Clause SPDXLicense = "BSD-1-Clause"
	// BSD 2-Clause FreeBSD License.
	SPDXLicenseBSD2ClauseFreeBSD SPDXLicense = "BSD-2-Clause-FreeBSD"
	// BSD 2-Clause NetBSD License.
	SPDXLicenseBSD2ClauseNetBSD SPDXLicense = "BSD-2-Clause-NetBSD"
	// BSD-2-Clause Plus Patent License.
	SPDXLicenseBSD2ClausePatent SPDXLicense = "BSD-2-Clause-Patent"
	// BSD 2-Clause "Simplified" License.
	SPDXLicenseBSD2Clause SPDXLicense = "BSD-2-Clause"
	// BSD with attribution.
	SPDXLicenseBSD3ClauseAttribution SPDXLicense = "BSD-3-Clause-Attribution"
	// BSD 3-Clause Clear License.
	SPDXLicenseBSD3ClauseClear SPDXLicense = "BSD-3-Clause-Clear"
	// Lawrence Berkeley National Labs BSD variant license.
	SPDXLicenseBSD3ClauseLBNL SPDXLicense = "BSD-3-Clause-LBNL"
	// BSD 3-Clause No Nuclear License 2014.
	SPDXLicenseBSD3ClauseNoNuclearLicense2014 SPDXLicense = "BSD-3-Clause-No-Nuclear-License-2014"
	// BSD 3-Clause No Nuclear License.
	SPDXLicenseBSD3ClauseNoNuclearLicense SPDXLicense = "BSD-3-Clause-No-Nuclear-License"
	// BSD 3-Clause No Nuclear Warranty.
	SPDXLicenseBSD3ClauseNoNuclearWarranty SPDXLicense = "BSD-3-Clause-No-Nuclear-Warranty"
	// BSD 3-Clause "New" or "Revised" License.
	SPDXLicenseBSD3Clause SPDXLicense = "BSD-3-Clause"
	// BSD-4-Clause (University of California-Specific).
	SPDXLicenseBSD4ClauseUC SPDXLicense = "BSD-4-Clause-UC"
	// BSD 4-Clause "Original" or "Old" License.
	SPDXLicenseBSD4Clause SPDXLicense = "BSD-4-Clause"
	// BSD Protection License.
	SPDXLicenseBSDProtection SPDXLicense = "BSD-Protection"
	// BSD Source Code Attribution.
	SPDXLicenseBSDSourceCode SPDXLicense = "BSD-Source-Code"
	// Boost Software License 1.0.
	SPDXLicenseBSL1_0 SPDXLicense = "BSL-1.0"
	// bzip2 and libbzip2 License v1.0.5.
	SPDXLicenseBzip21_0_5 SPDXLicense = "bzip2-1.0.5"
	// bzip2 and libbzip2 License v1.0.6.
	SPDXLicenseBzip21_0_6 SPDXLicense = "bzip2-1.0.6"
	// Caldera License.
	SPDXLicenseCaldera SPDXLicense = "Caldera"
	// Computer Associates Trusted Open Source License 1.1.
	SPDXLicenseCATOSL1_1 SPDXLicense = "CATOSL-1.1"
	// Creative Commons Attribution 1.0 Generic.
	SPDXLicenseCCBY1_0 SPDXLicense = "CC-BY-1.0"
	// Creative Commons Attribution 2.0 Generic.
	SPDXLicenseCCBY2_0 SPDXLicense = "CC-BY-2.0"
	// Creative Commons Attribution 2.5 Generic.
	SPDXLicenseCCBY2_5 SPDXLicense = "CC-BY-2.5"
	// Creative Commons Attribution 3.0 Unported.
	SPDXLicenseCCBY3_0 SPDXLicense = "CC-BY-3.0"
	// Creative Commons Attribution 4.0 International.
	SPDXLicenseCCBY4_0 SPDXLicense = "CC-BY-4.0"
	// Creative Commons Attribution Non Commercial 1.0 Generic.
	SPDXLicenseCCBYNC1_0 SPDXLicense = "CC-BY-NC-1.0"
	// Creative Commons Attribution Non Commercial 2.0 Generic.
	SPDXLicenseCCBYNC2_0 SPDXLicense = "CC-BY-NC-2.0"
	// Creative Commons Attribution Non Commercial 2.5 Generic.
	SPDXLicenseCCBYNC2_5 SPDXLicense = "CC-BY-NC-2.5"
	// Creative Commons Attribution Non Commercial 3.0 Unported.
	SPDXLicenseCCBYNC3_0 SPDXLicense = "CC-BY-NC-3.0"
	// Creative Commons Attribution Non Commercial 4.0 International.
	SPDXLicenseCCBYNC4_0 SPDXLicense = "CC-BY-NC-4.0"
	// Creative Commons Attribution Non Commercial No Derivatives 1.0 Generic.
	SPDXLicenseCCBYNCND1_0 SPDXLicense = "CC-BY-NC-ND-1.0"
	// Creative Commons Attribution Non Commercial No Derivatives 2.0 Generic.
	SPDXLicenseCCBYNCND2_0 SPDXLicense = "CC-BY-NC-ND-2.0"
	// Creative Commons Attribution Non Commercial No Derivatives 2.5 Generic.
	SPDXLicenseCCBYNCND2_5 SPDXLicense = "CC-BY-NC-ND-2.5"
	// Creative Commons Attribution Non Commercial No Derivatives 3.0 Unported.
	SPDXLicenseCCBYNCND3_0 SPDXLicense = "CC-BY-NC-ND-3.0"
	// Creative Commons Attribution Non Commercial No Derivatives 4.0 International.
	SPDXLicenseCCBYNCND4_0 SPDXLicense = "CC-BY-NC-ND-4.0"
	// Creative Commons Attribution Non Commercial Share Alike 1.0 Generic.
	SPDXLicenseCCBYNCSA1_0 SPDXLicense = "CC-BY-NC-SA-1.0"
	// Creative Commons Attribution Non Commercial Share Alike 2.0 Generic.
	SPDXLicenseCCBYNCSA2_0 SPDXLicense = "CC-BY-NC-SA-2.0"
	// Creative Commons Attribution Non Commercial Share Alike 2.5 Generic.
	SPDXLicenseCCBYNCSA2_5 SPDXLicense = "CC-BY-NC-SA-2.5"
	// Creative Commons Attribution Non Commercial Share Alike 3.0 Unported.
	SPDXLicenseCCBYNCSA3_0 SPDXLicense = "CC-BY-NC-SA-3.0"
	// Creative Commons Attribution Non Commercial Share Alike 4.0 International.
	SPDXLicenseCCBYNCSA4_0 SPDXLicense = "CC-BY-NC-SA-4.0"
	// Creative Commons Attribution No Derivatives 1.0 Generic.
	SPDXLicenseCCBYND1_0 SPDXLicense = "CC-BY-ND-1.0"
	// Creative Commons Attribution No Derivatives 2.0 Generic.
	SPDXLicenseCCBYND2_0 SPDXLicense = "CC-BY-ND-2.0"
	// Creative Commons Attribution No Derivatives 2.5 Generic.
	SPDXLicenseCCBYND2_5 SPDXLicense = "CC-BY-ND-2.5"
	// Creative Commons Attribution No Derivatives 3.0 Unported.
	SPDXLicenseCCBYND3_0 SPDXLicense = "CC-BY-ND-3.0"
	// Creative Commons Attribution No Derivatives 4.0 International.
	SPDXLicenseCCBYND4_0 SPDXLicense = "CC-BY-ND-4.0"
	// Creative Commons Attribution Share Alike 1.0 Generic.
	SPDXLicenseCCBYSA1_0 SPDXLicense = "CC-BY-SA-1.0"
	// Creative Commons Attribution Share Alike 2.0 Generic.
	SPDXLicenseCCBYSA2_0 SPDXLicense = "CC-BY-SA-2.0"
	// Creative Commons Attribution Share Alike 2.5 Generic.
	SPDXLicenseCCBYSA2_5 SPDXLicense = "CC-BY-SA-2.5"
	// Creative Commons Attribution Share Alike 3.0 Unported.
	SPDXLicenseCCBYSA3_0 SPDXLicense = "CC-BY-SA-3.0"
	// Creative Commons Attribution Share Alike 4.0 International.
	SPDXLicenseCCBYSA4_0 SPDXLicense = "CC-BY-SA-4.0"
	// Creative Commons Zero v1.0 Universal.
	SPDXLicenseCC01_0 SPDXLicense = "CC0-1.0"
	// Common Development and Distribution License 1.0.
	SPDXLicenseCDDL1_0 SPDXLicense = "CDDL-1.0"
	// Common Development and Distribution License 1.1.
	SPDXLicenseCDDL1_1 SPDXLicense = "CDDL-1.1"
	// Community Data License Agreement Permissive 1.0.
	SPDXLicenseCDLAPermissive1_0 SPDXLicense = "CDLA-Permissive-1.0"
	// Community Data License Agreement Sharing 1.0.
	SPDXLicenseCDLASharing1_0 SPDXLicense = "CDLA-Sharing-1.0"
	// CeCILL Free Software License Agreement v1.0.
	SPDXLicenseCECILL1_0 SPDXLicense = "CECILL-1.0"
	// CeCILL Free Software License Agreement v1.1.
	SPDXLicenseCECILL1_1 SPDXLicense = "CECILL-1.1"
	// CeCILL Free Software License Agreement v2.0.
	SPDXLicenseCECILL2_0 SPDXLicense = "CECILL-2.0"
	// CeCILL Free Software License Agreement v2.1.
	SPDXLicenseCECILL2_1 SPDXLicense = "CECILL-2.1"
	// CeCILL-B Free Software License Agreement.
	SPDXLicenseCECILLB SPDXLicense = "CECILL-B"
	// CeCILL-C Free Software License Agreement.
	SPDXLicenseCECILLC SPDXLicense = "CECILL-C"
	// Clarified Artistic License.
	SPDXLicenseClArtistic SPDXLicense = "ClArtistic"
	// CNRI Jython License.
	SPDXLicenseCNRIJython SPDXLicense = "CNRI-Jython"
	// CNRI Python Open Source GPL Compatible License Agreement.
	SPDXLicenseCNRIPythonGPLCompatible SPDXLicense = "CNRI-Python-GPL-Compatible"
	// CNRI Python License.
	SPDXLicenseCNRIPython SPDXLicense = "CNRI-Python"
	// Condor Public License v1.1.
	SPDXLicenseCondor1_1 SPDXLicense = "Condor-1.1"
	// Common Public Attribution License 1.0.
	SPDXLicenseCPAL1_0 SPDXLicense = "CPAL-1.0"
	// Common Public License 1.0.
	SPDXLicenseCPL1_0 SPDXLicense = "CPL-1.0"
	// Code Project Open License 1.02.
	SPDXLicenseCPOL1_02 SPDXLicense = "CPOL-1.02"
	// Crossword License.
	SPDXLicenseCrossword SPDXLicense = "Crossword"
	// CrystalStacker License.
	SPDXLicenseCrystalStacker SPDXLicense = "CrystalStacker"
	// CUA Office Public License v1.0.
	SPDXLicenseCUAOPL1_0 SPDXLicense = "CUA-OPL-1.0"
	// Cube License.
	SPDXLicenseCube SPDXLicense = "Cube"
	// curl License.
	SPDXLicenseCurl SPDXLicense = "curl"
	// Deutsche Freie Software Lizenz.
	SPDXLicenseDFSL1_0 SPDXLicense = "D-FSL-1.0"
	// diffmark license.
	SPDXLicenseDiffmark SPDXLicense = "diffmark"
	// DOC License.
	SPDXLicenseDOC SPDXLicense = "DOC"
	// Dotseqn License.
	SPDXLicenseDotseqn SPDXLicense = "Dotseqn"
	// DSDP License.
	SPDXLicenseDSDP SPDXLicense = "DSDP"
	// dvipdfm License.
	SPDXLicenseDvipdfm SPDXLicense = "dvipdfm"
	// Educational Community License v1.0.
	SPDXLicenseECL1_0 SPDXLicense = "ECL-1.0"
	// Educational Community License v2.0.
	SPDXLicenseECL2_0 SPDXLicense = "ECL-2.0"
	// Eiffel Forum License v1.0.
	SPDXLicenseEFL1_0 SPDXLicense = "EFL-1.0"
	// Eiffel Forum License v2.0.
	SPDXLicenseEFL2_0 SPDXLicense = "EFL-2.0"
	// eGenix.com Public License 1.1.0.
	SPDXLicenseEGenix SPDXLicense = "eGenix"
	// Entessa Public License v1.0.
	SPDXLicenseEntessa SPDXLicense = "Entessa"
	// Eclipse Public License 1.0.
	SPDXLicenseEPL1_0 SPDXLicense = "EPL-1.0"
	// Eclipse Public License 2.0.
	SPDXLicenseEPL2_0 SPDXLicense = "EPL-2.0"
	// Erlang Public License v1.1.
	SPDXLicenseErlPL1_1 SPDXLicense = "ErlPL-1.1"
	// EU DataGrid Software License.
	SPDXLicenseEUDatagrid SPDXLicense = "EUDatagrid"
	// European Union Public License 1.0.
	SPDXLicenseEUPL1_0 SPDXLicense = "EUPL-1.0"
	// European Union Public License 1.1.
	SPDXLicenseEUPL1_1 SPDXLicense = "EUPL-1.1"
	// European Union Public License 1.2.
	SPDXLicenseEUPL1_2 SPDXLicense = "EUPL-1.2"
	// Eurosym License.
	SPDXLicenseEurosym SPDXLicense = "Eurosym"
	// Fair License.
	SPDXLicenseFair SPDXLicense = "Fair"
	// Frameworx Open License 1.0.
	SPDXLicenseFrameworx1_0 SPDXLicense = "Frameworx-1.0"
	// FreeImage Public License v1.0.
	SPDXLicenseFreeImage SPDXLicense = "FreeImage"
	// FSF All Permissive License.
	SPDXLicenseFSFAP SPDXLicense = "FSFAP"
	// FSF Unlimited License.
	SPDXLicenseFSFUL SPDXLicense = "FSFUL"
	// FSF Unlimited License (with License Retention).
	SPDXLicenseFSFULLR SPDXLicense = "FSFULLR"
	// Freetype Project License.
	SPDXLicenseFTL SPDXLicense = "FTL"
	// GNU Free Documentation License v1.1 only.
	SPDXLicenseGFDL1_1Only SPDXLicense = "GFDL-1.1-only"
	// GNU Free Documentation License v1.1 or later.
	SPDXLicenseGFDL1_1OrLater SPDXLicense = "GFDL-1.1-or-later"
	// GNU Free Documentation License v1.2 only.
	SPDXLicenseGFDL1_2Only SPDXLicense = "GFDL-1.2-only"
	// GNU Free Documentation License v1.2 or later.
	SPDXLicenseGFDL1_2OrLater SPDXLicense = "GFDL-1.2-or-later"
	// GNU Free Documentation License v1.3 only.
	SPDXLicenseGFDL1_3Only SPDXLicense = "GFDL-1.3-only"
	// GNU Free Documentation License v1.3 or later.
	SPDXLicenseGFDL1_3OrLater SPDXLicense = "GFDL-1.3-or-later"
	// Giftware License.
	SPDXLicenseGiftware SPDXLicense = "Giftware"
	// GL2PS License.
	SPDXLicenseGL2PS SPDXLicense = "GL2PS"
	// 3dfx Glide License.
	SPDXLicenseGlide SPDXLicense = "Glide"
	// Glulxe License.
	SPDXLicenseGlulxe SPDXLicense = "Glulxe"
	// gnuplot License.
	SPDXLicenseGnuplot SPDXLicense = "gnuplot"
	// GNU General Public License v1.0 only.
	SPDXLicenseGPL1_0Only SPDXLicense = "GPL-1.0-only"
	// GNU General Public License v1.0 or later.
	SPDXLicenseGPL1_0OrLater SPDXLicense = "GPL-1.0-or-later"
	// GNU General Public License v2.0 only.
	SPDXLicenseGPL2_0Only SPDXLicense = "GPL-2.0-only"
	// GNU General Public License v2.0 or later.
	SPDXLicenseGPL2_0OrLater SPDXLicense = "GPL-2.0-or-later"
	// GNU General Public License v3.0 only.
	SPDXLicenseGPL3_0Only SPDXLicense = "GPL-3.0-only"
	// GNU General Public License v3.0 or later.
	SPDXLicenseGPL3_0OrLater SPDXLicense = "GPL-3.0-or-later"
	// gSOAP Public License v1.3b.
	SPDXLicenseGSOAP1_3b SPDXLicense = "gSOAP-1.3b"
	// Haskell Language Report License.
	SPDXLicenseHaskellReport SPDXLicense = "HaskellReport"
	// Historical Permission Notice and Disclaimer.
	SPDXLicenseHPND SPDXLicense = "HPND"
	// IBM PowerPC Initialization and Boot Software.
	SPDXLicenseIBMPibs SPDXLicense = "IBM-pibs"
	// ICU License.
	SPDXLicenseICU SPDXLicense = "ICU"
	// Independent JPEG Group License.
	SPDXLicenseIJG SPDXLicense = "IJG"
	// ImageMagick License.
	SPDXLicenseImageMagick SPDXLicense = "ImageMagick"
	// iMatix Standard Function Library Agreement.
	SPDXLicenseIMatix SPDXLicense = "iMatix"
	// Imlib2 License.
	SPDXLicenseImlib2 SPDXLicense = "Imlib2"
	// Info-ZIP License.
	SPDXLicenseInfoZIP SPDXLicense = "Info-ZIP"
	// Intel ACPI Software License Agreement.
	SPDXLicenseIntelACPI SPDXLicense = "Intel-ACPI"
	// Intel Open Source License.
	SPDXLicenseIntel SPDXLicense = "Intel"
	// Interbase Public License v1.0.
	SPDXLicenseInterbase1_0 SPDXLicense = "Interbase-1.0"
	// IPA Font License.
	SPDXLicenseIPA SPDXLicense = "IPA"
	// IBM Public License v1.0.
	SPDXLicenseIPL1_0 SPDXLicense = "IPL-1.0"
	// ISC License.
	SPDXLicenseISC SPDXLicense = "ISC"
	// JasPer License.
	SPDXLicenseJasPer2_0 SPDXLicense = "JasPer-2.0"
	// JSON License.
	SPDXLicenseJSON SPDXLicense = "JSON"
	// Licence Art Libre 1.2.
	SPDXLicenseLAL1_2 SPDXLicense = "LAL-1.2"
	// Licence Art Libre 1.3.
	SPDXLicenseLAL1_3 SPDXLicense = "LAL-1.3"
	// Latex2e License.
	SPDXLicenseLatex2e SPDXLicense = "Latex2e"
	// Leptonica License.
	SPDXLicenseLeptonica SPDXLicense = "Leptonica"
	// GNU Library General Public License v2 only.
	SPDXLicenseLGPL2_0Only SPDXLicense = "LGPL-2.0-only"
	// GNU Library General Public License v2 or later.
	SPDXLicenseLGPL2_0OrLater SPDXLicense = "LGPL-2.0-or-later"
	// GNU Lesser General Public License v2.1 only.
	SPDXLicenseLGPL2_1Only SPDXLicense = "LGPL-2.1-only"
	// GNU Lesser General Public License v2.1 or later.
	SPDXLicenseLGPL2_1OrLater SPDXLicense = "LGPL-2.1-or-later"
	// GNU Lesser General Public License v3.0 only.
	SPDXLicenseLGPL3_0Only SPDXLicense = "LGPL-3.0-only"
	// GNU Lesser General Public License v3.0 or later.
	SPDXLicenseLGPL3_0OrLater SPDXLicense = "LGPL-3.0-or-later"
	// Lesser General Public License For Linguistic Resources.
	SPDXLicenseLGPLLR SPDXLicense = "LGPLLR"
	// libpng License.
	SPDXLicenseLibpng SPDXLicense = "Libpng"
	// libtiff License.
	SPDXLicenseLibtiff SPDXLicense = "libtiff"
	// Licence Libre du Québec – Permissive version 1.1.
	SPDXLicenseLiLiQP1_1 SPDXLicense = "LiLiQ-P-1.1"
	// Licence Libre du Québec – Réciprocité version 1.1.
	SPDXLicenseLiLiQR1_1 SPDXLicense = "LiLiQ-R-1.1"
	// Licence Libre du Québec – Réciprocité forte version 1.1.
	SPDXLicenseLiLiQRplus1_1 SPDXLicense = "LiLiQ-Rplus-1.1"
	// Linux Kernel Variant of OpenIB.org license.
	SPDXLicenseLinuxOpenIB SPDXLicense = "Linux-OpenIB"
	// Lucent Public License Version 1.0.
	SPDXLicenseLPL1_0 SPDXLicense = "LPL-1.0"
	// Lucent Public License v1.02.
	SPDXLicenseLPL1_02 SPDXLicense = "LPL-1.02"
	// LaTeX Project Public License v1.0.
	SPDXLicenseLPPL1_0 SPDXLicense = "LPPL-1.0"
	// LaTeX Project Public License v1.1.
	SPDXLicenseLPPL1_1 SPDXLicense = "LPPL-1.1"
	// LaTeX Project Public License v1.2.
	SPDXLicenseLPPL1_2 SPDXLicense = "LPPL-1.2"
	// LaTeX Project Public License v1.3a.
	SPDXLicenseLPPL1_3a SPDXLicense = "LPPL-1.3a"
	// LaTeX Project Public License v1.3c.
	SPDXLicenseLPPL1_3c SPDXLicense = "LPPL-1.3c"
	// MakeIndex License.
	SPDXLicenseMakeIndex SPDXLicense = "MakeIndex"
	// MirOS License.
	SPDXLicenseMirOS SPDXLicense = "MirOS"
	// MIT No Attribution.
	SPDXLicenseMIT0 SPDXLicense = "MIT-0"
	// Enlightenment License (e16).
	SPDXLicenseMITAdvertising SPDXLicense = "MIT-advertising"
	// CMU License.
	SPDXLicenseMITCMU SPDXLicense = "MIT-CMU"
	// enna License.
	SPDXLicenseMITEnna SPDXLicense = "MIT-enna"
	// feh License.
	SPDXLicenseMITFeh SPDXLicense = "MIT-feh"
	// MIT License.
	SPDXLicenseMIT SPDXLicense = "MIT"
	// MIT +no-false-attribs license.
	SPDXLicenseMITNFA SPDXLicense = "MITNFA"
	// Motosoto License.
	SPDXLicenseMotosoto SPDXLicense = "Motosoto"
	// mpich2 License.
	SPDXLicenseMpich2 SPDXLicense = "mpich2"
	// Mozilla Public License 1.0.
	SPDXLicenseMPL1_0 SPDXLicense = "MPL-1.0"
	// Mozilla Public License 1.1.
	SPDXLicenseMPL1_1 SPDXLicense = "MPL-1.1"
	// Mozilla Public License 2.0 (no copyleft exception).
	SPDXLicenseMPL2_0NoCopyleftException SPDXLicense = "MPL-2.0-no-copyleft-exception"
	// Mozilla Public License 2.0.
	SPDXLicenseMPL2_0 SPDXLicense = "MPL-2.0"
	// Microsoft Public License.
	SPDXLicenseMSPL SPDXLicense = "MS-PL"
	// Microsoft Reciprocal License.
	SPDXLicenseMSRL SPDXLicense = "MS-RL"
	// Matrix Template Library License.
	SPDXLicenseMTLL SPDXLicense = "MTLL"
	// Multics License.
	SPDXLicenseMultics SPDXLicense = "Multics"
	// Mup License.
	SPDXLicenseMup SPDXLicense = "Mup"
	// NASA Open Source Agreement 1.3.
	SPDXLicenseNASA1_3 SPDXLicense = "NASA-1.3"
	// Naumen Public License.
	SPDXLicenseNaumen SPDXLicense = "Naumen"
	// Net Boolean Public License v1.
	SPDXLicenseNBPL1_0 SPDXLicense = "NBPL-1.0"
	// University of Illinois/NCSA Open Source License.
	SPDXLicenseNCSA SPDXLicense = "NCSA"
	// Net-SNMP License.
	SPDXLicenseNetSNMP SPDXLicense = "Net-SNMP"
	// NetCDF license.
	SPDXLicenseNetCDF SPDXLicense = "NetCDF"
	// Newsletr License.
	SPDXLicenseNewsletr SPDXLicense = "Newsletr"
	// Nethack General Public License.
	SPDXLicenseNGPL SPDXLicense = "NGPL"
	// Norwegian Licence for Open Government Data.
	SPDXLicenseNLOD1_0 SPDXLicense = "NLOD-1.0"
	// No Limit Public License.
	SPDXLicenseNLPL SPDXLicense = "NLPL"
	// Nokia Open Source License.
	SPDXLicenseNokia SPDXLicense = "Nokia"
	// Netizen Open Source License.
	SPDXLicenseNOSL SPDXLicense = "NOSL"
	// Noweb License.
	SPDXLicenseNoweb SPDXLicense = "Noweb"
	// Netscape Public License v1.0.
	SPDXLicenseNPL1_0 SPDXLicense = "NPL-1.0"
	// Netscape Public License v1.1.
	SPDXLicenseNPL1_1 SPDXLicense = "NPL-1.1"
	// Non-Profit Open Software License 3.0.
	SPDXLicenseNPOSL3_0 SPDXLicense = "NPOSL-3.0"
	// NRL License.
	SPDXLicenseNRL SPDXLicense = "NRL"
	// NTP License.
	SPDXLicenseNTP SPDXLicense = "NTP"
	// Open CASCADE Technology Public License.
	SPDXLicenseOCCTPL SPDXLicense = "OCCT-PL"
	// OCLC Research Public License 2.0.
	SPDXLicenseOCLC2_0 SPDXLicense = "OCLC-2.0"
	// ODC Open Database License v1.0.
	SPDXLicenseODbL1_0 SPDXLicense = "ODbL-1.0"
	// SIL Open Font License 1.0.
	SPDXLicenseOFL1_0 SPDXLicense = "OFL-1.0"
	// SIL Open Font License 1.1.
	SPDXLicenseOFL1_1 SPDXLicense = "OFL-1.1"
	// Open Group Test Suite License.
	SPDXLicenseOGTSL SPDXLicense = "OGTSL"
	// Open LDAP Public License v1.1.
	SPDXLicenseOLDAP1_1 SPDXLicense = "OLDAP-1.1"
	// Open LDAP Public License v1.2.
	SPDXLicenseOLDAP1_2 SPDXLicense = "OLDAP-1.2"
	// Open LDAP Public License v1.3.
	SPDXLicenseOLDAP1_3 SPDXLicense = "OLDAP-1.3"
	// Open LDAP Public License v1.4.
	SPDXLicenseOLDAP1_4 SPDXLicense = "OLDAP-1.4"
	// Open LDAP Public License v2.0.1.
	SPDXLicenseOLDAP2_0_1 SPDXLicense = "OLDAP-2.0.1"
	// Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B).
	SPDXLicenseOLDAP2_0 SPDXLicense = "OLDAP-2.0"
	// Open LDAP Public License v2.1.
	SPDXLicenseOLDAP2_1 SPDXLicense = "OLDAP-2.1"
	// Open LDAP Public License v2.2.1.
	SPDXLicenseOLDAP2_2_1 SPDXLicense = "OLDAP-2.2.1"
	// Open LDAP Public License 2.2.2.
	SPDXLicenseOLDAP2_2_2 SPDXLicense = "OLDAP-2.2.2"
	// Open LDAP Public License v2.2.
	SPDXLicenseOLDAP2_2 SPDXLicense = "OLDAP-2.2"
	// Open LDAP Public License v2.3.
	SPDXLicenseOLDAP2_3 SPDXLicense = "OLDAP-2.3"
	// Open LDAP Public License v2.4.
	SPDXLicenseOLDAP2_4 SPDXLicense = "OLDAP-2.4"
	// Open LDAP Public License v2.5.
	SPDXLicenseOLDAP2_5 SPDXLicense = "OLDAP-2.5"
	// Open LDAP Public License v2.6.
	SPDXLicenseOLDAP2_6 SPDXLicense = "OLDAP-2.6"
	// Open LDAP Public License v2.7.
	SPDXLicenseOLDAP2_7 SPDXLicense = "OLDAP-2.7"
	// Open LDAP Public License v2.8.
	SPDXLicenseOLDAP2_8 SPDXLicense = "OLDAP-2.8"
	// Open Market License.
	SPDXLicenseOML SPDXLicense = "OML"
	// OpenSSL License.
	SPDXLicenseOpenSSL SPDXLicense = "OpenSSL"
	// Open Public License v1.0.
	SPDXLicenseOPL1_0 SPDXLicense = "OPL-1.0"
	// OSET Public License version 2.1.
	SPDXLicenseOSETPL2_1 SPDXLicense = "OSET-PL-2.1"
	// Open Software License 1.0.
	SPDXLicenseOSL1_0 SPDXLicense = "OSL-1.0"
	// Open Software License 1.1.
	SPDXLicenseOSL1_1 SPDXLicense = "OSL-1.1"
	// Open Software License 2.0.
	SPDXLicenseOSL2_0 SPDXLicense = "OSL-2.0"
	// Open Software License 2.1.
	SPDXLicenseOSL2_1 SPDXLicense = "OSL-2.1"
	// Open Software License 3.0.
	SPDXLicenseOSL3_0 SPDXLicense = "OSL-3.0"
	// ODC Public Domain Dedication & License 1.0.
	SPDXLicensePDDL1_0 SPDXLicense = "PDDL-1.0"
	// PHP License v3.0.
	SPDXLicensePHP3_0 SPDXLicense = "PHP-3.0"
	// PHP License v3.01.
	SPDXLicensePHP3_01 SPDXLicense = "PHP-3.01"
	// Plexus Classworlds License.
	SPDXLicensePlexus SPDXLicense = "Plexus"
	// PostgreSQL License.
	SPDXLicensePostgreSQL SPDXLicense = "PostgreSQL"
	// psfrag License.
	SPDXLicensePsfrag SPDXLicense = "psfrag"
	// psutils License.
	SPDXLicensePsutils SPDXLicense = "psutils"
	// Python License 2.0.
	SPDXLicensePython2_0 SPDXLicense = "Python-2.0"
	// Qhull License.
	SPDXLicenseQhull SPDXLicense = "Qhull"
	// Q Public License 1.0.
	SPDXLicenseQPL1_0 SPDXLicense = "QPL-1.0"
	// Rdisc License.
	SPDXLicenseRdisc SPDXLicense = "Rdisc"
	// Red Hat eCos Public License v1.1.
	SPDXLicenseRHeCos1_1 SPDXLicense = "RHeCos-1.1"
	// Reciprocal Public License 1.1.
	SPDXLicenseRPL1_1 SPDXLicense = "RPL-1.1"
	// Reciprocal Public License 1.5.
	SPDXLicenseRPL1_5 SPDXLicense = "RPL-1.5"
	// RealNetworks Public Source License v1.0.
	SPDXLicenseRPSL1_0 SPDXLicense = "RPSL-1.0"
	// RSA Message-Digest License.
	SPDXLicenseRSAMD SPDXLicense = "RSA-MD"
	// Ricoh Source Code Public License.
	SPDXLicenseRSCPL SPDXLicense = "RSCPL"
	// Ruby License.
	SPDXLicenseRuby SPDXLicense = "Ruby"
	// Sax Public Domain Notice.
	SPDXLicenseSAXPD SPDXLicense = "SAX-PD"
	// Saxpath License.
	SPDXLicenseSaxpath SPDXLicense = "Saxpath"
	// SCEA Shared Source License.
	SPDXLicenseSCEA SPDXLicense = "SCEA"
	// Sendmail License.
	SPDXLicenseSendmail SPDXLicense = "Sendmail"
	// SGI Free Software License B v1.0.
	SPDXLicenseSGIB1_0 SPDXLicense = "SGI-B-1.0"
	// SGI Free Software License B v1.1.
	SPDXLicenseSGIB1_1 SPDXLicense = "SGI-B-1.1"
	// SGI Free Software License B v2.0.
	SPDXLicenseSGIB2_0 SPDXLicense = "SGI-B-2.0"
	// Simple Public License 2.0.
	SPDXLicenseSimPL2_0 SPDXLicense = "SimPL-2.0"
	// Sun Industry Standards Source License v1.2.
	SPDXLicenseSISSL1_2 SPDXLicense = "SISSL-1.2"
	// Sun Industry Standards Source License v1.1.
	SPDXLicenseSISSL SPDXLicense = "SISSL"
	// Sleepycat License.
	SPDXLicenseSleepycat SPDXLicense = "Sleepycat"
	// Standard ML of New Jersey License.
	SPDXLicenseSMLNJ SPDXLicense = "SMLNJ"
	// Secure Messaging Protocol Public License.
	SPDXLicenseSMPPL SPDXLicense = "SMPPL"
	// SNIA Public License 1.1.
	SPDXLicenseSNIA SPDXLicense = "SNIA"
	// Spencer License 86.
	SPDXLicenseSpencer86 SPDXLicense = "Spencer-86"
	// Spencer License 94.
	SPDXLicenseSpencer94 SPDXLicense = "Spencer-94"
	// Spencer License 99.
	SPDXLicenseSpencer99 SPDXLicense = "Spencer-99"
	// Sun Public License v1.0.
	SPDXLicenseSPL1_0 SPDXLicense = "SPL-1.0"
	// SugarCRM Public License v1.1.3.
	SPDXLicenseSugarCRM1_1_3 SPDXLicense = "SugarCRM-1.1.3"
	// Scheme Widget Library (SWL) Software License Agreement.
	SPDXLicenseSWL SPDXLicense = "SWL"
	// TCL/TK License.
	SPDXLicenseTCL SPDXLicense = "TCL"
	// TCP Wrappers License.
	SPDXLicenseTCPWrappers SPDXLicense = "TCP-wrappers"
	// TMate Open Source License.
	SPDXLicenseTMate SPDXLicense = "TMate"
	// TORQUE v2.5+ Software License v1.1.
	SPDXLicenseTORQUE1_1 SPDXLicense = "TORQUE-1.1"
	// Trusster Open Source License.
	SPDXLicenseTOSL SPDXLicense = "TOSL"
	// Unicode License Agreement - Data Files and Software (2015).
	SPDXLicenseUnicodeDFS2015 SPDXLicense = "Unicode-DFS-2015"
	// Unicode License Agreement - Data Files and Software (2016).
	SPDXLicenseUnicodeDFS2016 SPDXLicense = "Unicode-DFS-2016"
	// Unicode Terms of Use.
	SPDXLicenseUnicodeTOU SPDXLicense = "Unicode-TOU"
	// The Unlicense.
	SPDXLicenseUnlicense SPDXLicense = "Unlicense"
	// Universal Permissive License v1.0.
	SPDXLicenseUPL1_0 SPDXLicense = "UPL-1.0"
	// Vim License.
	SPDXLicenseVim SPDXLicense = "Vim"
	// VOSTROM Public License for Open Source.
	SPDXLicenseVOSTROM SPDXLicense = "VOSTROM"
	// Vovida Software License v1.0.
	SPDXLicenseVSL1_0 SPDXLicense = "VSL-1.0"
	// W3C Software Notice and License (1998-07-20).
	SPDXLicenseW3C19980720 SPDXLicense = "W3C-19980720"
	// W3C Software Notice and Document License (2015-05-13).
	SPDXLicenseW3C20150513 SPDXLicense = "W3C-20150513"
	// W3C Software Notice and License (2002-12-31).
	SPDXLicenseW3C SPDXLicense = "W3C"
	// Sybase Open Watcom Public License 1.0.
	SPDXLicenseWatcom1_0 SPDXLicense = "Watcom-1.0"
	// Wsuipa License.
	SPDXLicenseWsuipa SPDXLicense = "Wsuipa"
	// Do What The F*ck You Want To Public License.
	SPDXLicenseWTFPL SPDXLicense = "WTFPL"
	// X11 License.
	SPDXLicenseX11 SPDXLicense = "X11"
	// Xerox License.
	SPDXLicenseXerox SPDXLicense = "Xerox"
	// XFree86 License 1.1.
	SPDXLicenseXFree861_1 SPDXLicense = "XFree86-1.1"
	// xinetd License.
	SPDXLicenseXinetd SPDXLicense = "xinetd"
	// X.Net License.
	SPDXLicenseXnet SPDXLicense = "Xnet"
	// XPP License.
	SPDXLicenseXpp SPDXLicense = "xpp"
	// XSkat License.
	SPDXLicenseXSkat SPDXLicense = "XSkat"
	// Yahoo! Public License v1.0.
	SPDXLicenseYPL1_0 SPDXLicense = "YPL-1.0"
	// Yahoo! Public License v1.1.
	SPDXLicenseYPL1_1 SPDXLicense = "YPL-1.1"
	// Zed License.
	SPDXLicenseZed SPDXLicense = "Zed"
	// Zend License v2.0.
	SPDXLicenseZend2_0 SPDXLicense = "Zend-2.0"
	// Zimbra Public License v1.3.
	SPDXLicenseZimbra1_3 SPDXLicense = "Zimbra-1.3"
	// Zimbra Public License v1.4.
	SPDXLicenseZimbra1_4 SPDXLicense = "Zimbra-1.4"
	// zlib/libpng License with Acknowledgement.
	SPDXLicenseZlibAcknowledgement SPDXLicense = "zlib-acknowledgement"
	// zlib License.
	SPDXLicenseZlib SPDXLicense = "Zlib"
	// Zope Public License 1.1.
	SPDXLicenseZPL1_1 SPDXLicense = "ZPL-1.1"
	// Zope Public License 2.0.
	SPDXLicenseZPL2_0 SPDXLicense = "ZPL-2.0"
	// Zope Public License 2.1.
	SPDXLicenseZPL2_1 SPDXLicense = "ZPL-2.1"
)

func (code SPDXLicense) String() string {
	return code.Code()
}

func (code SPDXLicense) Code() string {
	return string(code)
}

func (code SPDXLicense) Display() string {
	switch code {
	case SPDXLicenseNotOpenSource:
		return "Not open source"
	case SPDXLicense0BSD:
		return "BSD Zero Clause License"
	case SPDXLicenseAAL:
		return "Attribution Assurance License"
	case SPDXLicenseAbstyles:
		return "Abstyles License"
	case SPDXLicenseAdobe2006:
		return "Adobe Systems Incorporated Source Code License Agreement"
	case SPDXLicenseAdobeGlyph:
		return "Adobe Glyph List License"
	case SPDXLicenseADSL:
		return "Amazon Digital Services License"
	case SPDXLicenseAFL1_1:
		return "Academic Free License v1.1"
	case SPDXLicenseAFL1_2:
		return "Academic Free License v1.2"
	case SPDXLicenseAFL2_0:
		return "Academic Free License v2.0"
	case SPDXLicenseAFL2_1:
		return "Academic Free License v2.1"
	case SPDXLicenseAFL3_0:
		return "Academic Free License v3.0"
	case SPDXLicenseAfmparse:
		return "Afmparse License"
	case SPDXLicenseAGPL1_0Only:
		return "Affero General Public License v1.0 only"
	case SPDXLicenseAGPL1_0OrLater:
		return "Affero General Public License v1.0 or later"
	case SPDXLicenseAGPL3_0Only:
		return "GNU Affero General Public License v3.0 only"
	case SPDXLicenseAGPL3_0OrLater:
		return "GNU Affero General Public License v3.0 or later"
	case SPDXLicenseAladdin:
		return "Aladdin Free Public License"
	case SPDXLicenseAMDPLPA:
		return "AMD's plpa_map.c License"
	case SPDXLicenseAML:
		return "Apple MIT License"
	case SPDXLicenseAMPAS:
		return "Academy of Motion Picture Arts and Sciences BSD"
	case SPDXLicenseANTLRPD:
		return "ANTLR Software Rights Notice"
	case SPDXLicenseApache1_0:
		return "Apache License 1.0"
	case SPDXLicenseApache1_1:
		return "Apache License 1.1"
	case SPDXLicenseApache2_0:
		return "Apache License 2.0"
	case SPDXLicenseAPAFML:
		return "Adobe Postscript AFM License"
	case SPDXLicenseAPL1_0:
		return "Adaptive Public License 1.0"
	case SPDXLicenseAPSL1_0:
		return "Apple Public Source License 1.0"
	case SPDXLicenseAPSL1_1:
		return "Apple Public Source License 1.1"
	case SPDXLicenseAPSL1_2:
		return "Apple Public Source License 1.2"
	case SPDXLicenseAPSL2_0:
		return "Apple Public Source License 2.0"
	case SPDXLicenseArtistic1_0Cl8:
		return "Artistic License 1.0 w/clause 8"
	case SPDXLicenseArtistic1_0Perl:
		return "Artistic License 1.0 (Perl)"
	case SPDXLicenseArtistic1_0:
		return "Artistic License 1.0"
	case SPDXLicenseArtistic2_0:
		return "Artistic License 2.0"
	case SPDXLicenseBahyph:
		return "Bahyph License"
	case SPDXLicenseBarr:
		return "Barr License"
	case SPDXLicenseBeerware:
		return "Beerware License"
	case SPDXLicenseBitTorrent1_0:
		return "BitTorrent Open Source License v1.0"
	case SPDXLicenseBitTorrent1_1:
		return "BitTorrent Open Source License v1.1"
	case SPDXLicenseBorceux:
		return "Borceux license"
	case SPDXLicenseBSD1Clause:
		return "BSD 1-Clause License"
	case SPDXLicenseBSD2ClauseFreeBSD:
		return "BSD 2-Clause FreeBSD License"
	case SPDXLicenseBSD2ClauseNetBSD:
		return "BSD 2-Clause NetBSD License"
	case SPDXLicenseBSD2ClausePatent:
		return "BSD-2-Clause Plus Patent License"
	case SPDXLicenseBSD2Clause:
		return "BSD 2-Clause \"Simplified\" License"
	case SPDXLicenseBSD3ClauseAttribution:
		return "BSD with attribution"
	case SPDXLicenseBSD3ClauseClear:
		return "BSD 3-Clause Clear License"
	case SPDXLicenseBSD3ClauseLBNL:
		return "Lawrence Berkeley National Labs BSD variant license"
	case SPDXLicenseBSD3ClauseNoNuclearLicense2014:
		return "BSD 3-Clause No Nuclear License 2014"
	case SPDXLicenseBSD3ClauseNoNuclearLicense:
		return "BSD 3-Clause No Nuclear License"
	case SPDXLicenseBSD3ClauseNoNuclearWarranty:
		return "BSD 3-Clause No Nuclear Warranty"
	case SPDXLicenseBSD3Clause:
		return "BSD 3-Clause \"New\" or \"Revised\" License"
	case SPDXLicenseBSD4ClauseUC:
		return "BSD-4-Clause (University of California-Specific)"
	case SPDXLicenseBSD4Clause:
		return "BSD 4-Clause \"Original\" or \"Old\" License"
	case SPDXLicenseBSDProtection:
		return "BSD Protection License"
	case SPDXLicenseBSDSourceCode:
		return "BSD Source Code Attribution"
	case SPDXLicenseBSL1_0:
		return "Boost Software License 1.0"
	case SPDXLicenseBzip21_0_5:
		return "bzip2 and libbzip2 License v1.0.5"
	case SPDXLicenseBzip21_0_6:
		return "bzip2 and libbzip2 License v1.0.6"
	case SPDXLicenseCaldera:
		return "Caldera License"
	case SPDXLicenseCATOSL1_1:
		return "Computer Associates Trusted Open Source License 1.1"
	case SPDXLicenseCCBY1_0:
		return "Creative Commons Attribution 1.0 Generic"
	case SPDXLicenseCCBY2_0:
		return "Creative Commons Attribution 2.0 Generic"
	case SPDXLicenseCCBY2_5:
		return "Creative Commons Attribution 2.5 Generic"
	case SPDXLicenseCCBY3_0:
		return "Creative Commons Attribution 3.0 Unported"
	case SPDXLicenseCCBY4_0:
		return "Creative Commons Attribution 4.0 International"
	case SPDXLicenseCCBYNC1_0:
		return "Creative Commons Attribution Non Commercial 1.0 Generic"
	case SPDXLicenseCCBYNC2_0:
		return "Creative Commons Attribution Non Commercial 2.0 Generic"
	case SPDXLicenseCCBYNC2_5:
		return "Creative Commons Attribution Non Commercial 2.5 Generic"
	case SPDXLicenseCCBYNC3_0:
		return "Creative Commons Attribution Non Commercial 3.0 Unported"
	case SPDXLicenseCCBYNC4_0:
		return "Creative Commons Attribution Non Commercial 4.0 International"
	case SPDXLicenseCCBYNCND1_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 1.0 Generic"
	case SPDXLicenseCCBYNCND2_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 2.0 Generic"
	case SPDXLicenseCCBYNCND2_5:
		return "Creative Commons Attribution Non Commercial No Derivatives 2.5 Generic"
	case SPDXLicenseCCBYNCND3_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 3.0 Unported"
	case SPDXLicenseCCBYNCND4_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 4.0 International"
	case SPDXLicenseCCBYNCSA1_0:
		return "Creative Commons Attribution Non Commercial Share Alike 1.0 Generic"
	case SPDXLicenseCCBYNCSA2_0:
		return "Creative Commons Attribution Non Commercial Share Alike 2.0 Generic"
	case SPDXLicenseCCBYNCSA2_5:
		return "Creative Commons Attribution Non Commercial Share Alike 2.5 Generic"
	case SPDXLicenseCCBYNCSA3_0:
		return "Creative Commons Attribution Non Commercial Share Alike 3.0 Unported"
	case SPDXLicenseCCBYNCSA4_0:
		return "Creative Commons Attribution Non Commercial Share Alike 4.0 International"
	case SPDXLicenseCCBYND1_0:
		return "Creative Commons Attribution No Derivatives 1.0 Generic"
	case SPDXLicenseCCBYND2_0:
		return "Creative Commons Attribution No Derivatives 2.0 Generic"
	case SPDXLicenseCCBYND2_5:
		return "Creative Commons Attribution No Derivatives 2.5 Generic"
	case SPDXLicenseCCBYND3_0:
		return "Creative Commons Attribution No Derivatives 3.0 Unported"
	case SPDXLicenseCCBYND4_0:
		return "Creative Commons Attribution No Derivatives 4.0 International"
	case SPDXLicenseCCBYSA1_0:
		return "Creative Commons Attribution Share Alike 1.0 Generic"
	case SPDXLicenseCCBYSA2_0:
		return "Creative Commons Attribution Share Alike 2.0 Generic"
	case SPDXLicenseCCBYSA2_5:
		return "Creative Commons Attribution Share Alike 2.5 Generic"
	case SPDXLicenseCCBYSA3_0:
		return "Creative Commons Attribution Share Alike 3.0 Unported"
	case SPDXLicenseCCBYSA4_0:
		return "Creative Commons Attribution Share Alike 4.0 International"
	case SPDXLicenseCC01_0:
		return "Creative Commons Zero v1.0 Universal"
	case SPDXLicenseCDDL1_0:
		return "Common Development and Distribution License 1.0"
	case SPDXLicenseCDDL1_1:
		return "Common Development and Distribution License 1.1"
	case SPDXLicenseCDLAPermissive1_0:
		return "Community Data License Agreement Permissive 1.0"
	case SPDXLicenseCDLASharing1_0:
		return "Community Data License Agreement Sharing 1.0"
	case SPDXLicenseCECILL1_0:
		return "CeCILL Free Software License Agreement v1.0"
	case SPDXLicenseCECILL1_1:
		return "CeCILL Free Software License Agreement v1.1"
	case SPDXLicenseCECILL2_0:
		return "CeCILL Free Software License Agreement v2.0"
	case SPDXLicenseCECILL2_1:
		return "CeCILL Free Software License Agreement v2.1"
	case SPDXLicenseCECILLB:
		return "CeCILL-B Free Software License Agreement"
	case SPDXLicenseCECILLC:
		return "CeCILL-C Free Software License Agreement"
	case SPDXLicenseClArtistic:
		return "Clarified Artistic License"
	case SPDXLicenseCNRIJython:
		return "CNRI Jython License"
	case SPDXLicenseCNRIPythonGPLCompatible:
		return "CNRI Python Open Source GPL Compatible License Agreement"
	case SPDXLicenseCNRIPython:
		return "CNRI Python License"
	case SPDXLicenseCondor1_1:
		return "Condor Public License v1.1"
	case SPDXLicenseCPAL1_0:
		return "Common Public Attribution License 1.0"
	case SPDXLicenseCPL1_0:
		return "Common Public License 1.0"
	case SPDXLicenseCPOL1_02:
		return "Code Project Open License 1.02"
	case SPDXLicenseCrossword:
		return "Crossword License"
	case SPDXLicenseCrystalStacker:
		return "CrystalStacker License"
	case SPDXLicenseCUAOPL1_0:
		return "CUA Office Public License v1.0"
	case SPDXLicenseCube:
		return "Cube License"
	case SPDXLicenseCurl:
		return "curl License"
	case SPDXLicenseDFSL1_0:
		return "Deutsche Freie Software Lizenz"
	case SPDXLicenseDiffmark:
		return "diffmark license"
	case SPDXLicenseDOC:
		return "DOC License"
	case SPDXLicenseDotseqn:
		return "Dotseqn License"
	case SPDXLicenseDSDP:
		return "DSDP License"
	case SPDXLicenseDvipdfm:
		return "dvipdfm License"
	case SPDXLicenseECL1_0:
		return "Educational Community License v1.0"
	case SPDXLicenseECL2_0:
		return "Educational Community License v2.0"
	case SPDXLicenseEFL1_0:
		return "Eiffel Forum License v1.0"
	case SPDXLicenseEFL2_0:
		return "Eiffel Forum License v2.0"
	case SPDXLicenseEGenix:
		return "eGenix.com Public License 1.1.0"
	case SPDXLicenseEntessa:
		return "Entessa Public License v1.0"
	case SPDXLicenseEPL1_0:
		return "Eclipse Public License 1.0"
	case SPDXLicenseEPL2_0:
		return "Eclipse Public License 2.0"
	case SPDXLicenseErlPL1_1:
		return "Erlang Public License v1.1"
	case SPDXLicenseEUDatagrid:
		return "EU DataGrid Software License"
	case SPDXLicenseEUPL1_0:
		return "European Union Public License 1.0"
	case SPDXLicenseEUPL1_1:
		return "European Union Public License 1.1"
	case SPDXLicenseEUPL1_2:
		return "European Union Public License 1.2"
	case SPDXLicenseEurosym:
		return "Eurosym License"
	case SPDXLicenseFair:
		return "Fair License"
	case SPDXLicenseFrameworx1_0:
		return "Frameworx Open License 1.0"
	case SPDXLicenseFreeImage:
		return "FreeImage Public License v1.0"
	case SPDXLicenseFSFAP:
		return "FSF All Permissive License"
	case SPDXLicenseFSFUL:
		return "FSF Unlimited License"
	case SPDXLicenseFSFULLR:
		return "FSF Unlimited License (with License Retention)"
	case SPDXLicenseFTL:
		return "Freetype Project License"
	case SPDXLicenseGFDL1_1Only:
		return "GNU Free Documentation License v1.1 only"
	case SPDXLicenseGFDL1_1OrLater:
		return "GNU Free Documentation License v1.1 or later"
	case SPDXLicenseGFDL1_2Only:
		return "GNU Free Documentation License v1.2 only"
	case SPDXLicenseGFDL1_2OrLater:
		return "GNU Free Documentation License v1.2 or later"
	case SPDXLicenseGFDL1_3Only:
		return "GNU Free Documentation License v1.3 only"
	case SPDXLicenseGFDL1_3OrLater:
		return "GNU Free Documentation License v1.3 or later"
	case SPDXLicenseGiftware:
		return "Giftware License"
	case SPDXLicenseGL2PS:
		return "GL2PS License"
	case SPDXLicenseGlide:
		return "3dfx Glide License"
	case SPDXLicenseGlulxe:
		return "Glulxe License"
	case SPDXLicenseGnuplot:
		return "gnuplot License"
	case SPDXLicenseGPL1_0Only:
		return "GNU General Public License v1.0 only"
	case SPDXLicenseGPL1_0OrLater:
		return "GNU General Public License v1.0 or later"
	case SPDXLicenseGPL2_0Only:
		return "GNU General Public License v2.0 only"
	case SPDXLicenseGPL2_0OrLater:
		return "GNU General Public License v2.0 or later"
	case SPDXLicenseGPL3_0Only:
		return "GNU General Public License v3.0 only"
	case SPDXLicenseGPL3_0OrLater:
		return "GNU General Public License v3.0 or later"
	case SPDXLicenseGSOAP1_3b:
		return "gSOAP Public License v1.3b"
	case SPDXLicenseHaskellReport:
		return "Haskell Language Report License"
	case SPDXLicenseHPND:
		return "Historical Permission Notice and Disclaimer"
	case SPDXLicenseIBMPibs:
		return "IBM PowerPC Initialization and Boot Software"
	case SPDXLicenseICU:
		return "ICU License"
	case SPDXLicenseIJG:
		return "Independent JPEG Group License"
	case SPDXLicenseImageMagick:
		return "ImageMagick License"
	case SPDXLicenseIMatix:
		return "iMatix Standard Function Library Agreement"
	case SPDXLicenseImlib2:
		return "Imlib2 License"
	case SPDXLicenseInfoZIP:
		return "Info-ZIP License"
	case SPDXLicenseIntelACPI:
		return "Intel ACPI Software License Agreement"
	case SPDXLicenseIntel:
		return "Intel Open Source License"
	case SPDXLicenseInterbase1_0:
		return "Interbase Public License v1.0"
	case SPDXLicenseIPA:
		return "IPA Font License"
	case SPDXLicenseIPL1_0:
		return "IBM Public License v1.0"
	case SPDXLicenseISC:
		return "ISC License"
	case SPDXLicenseJasPer2_0:
		return "JasPer License"
	case SPDXLicenseJSON:
		return "JSON License"
	case SPDXLicenseLAL1_2:
		return "Licence Art Libre 1.2"
	case SPDXLicenseLAL1_3:
		return "Licence Art Libre 1.3"
	case SPDXLicenseLatex2e:
		return "Latex2e License"
	case SPDXLicenseLeptonica:
		return "Leptonica License"
	case SPDXLicenseLGPL2_0Only:
		return "GNU Library General Public License v2 only"
	case SPDXLicenseLGPL2_0OrLater:
		return "GNU Library General Public License v2 or later"
	case SPDXLicenseLGPL2_1Only:
		return "GNU Lesser General Public License v2.1 only"
	case SPDXLicenseLGPL2_1OrLater:
		return "GNU Lesser General Public License v2.1 or later"
	case SPDXLicenseLGPL3_0Only:
		return "GNU Lesser General Public License v3.0 only"
	case SPDXLicenseLGPL3_0OrLater:
		return "GNU Lesser General Public License v3.0 or later"
	case SPDXLicenseLGPLLR:
		return "Lesser General Public License For Linguistic Resources"
	case SPDXLicenseLibpng:
		return "libpng License"
	case SPDXLicenseLibtiff:
		return "libtiff License"
	case SPDXLicenseLiLiQP1_1:
		return "Licence Libre du Québec – Permissive version 1.1"
	case SPDXLicenseLiLiQR1_1:
		return "Licence Libre du Québec – Réciprocité version 1.1"
	case SPDXLicenseLiLiQRplus1_1:
		return "Licence Libre du Québec – Réciprocité forte version 1.1"
	case SPDXLicenseLinuxOpenIB:
		return "Linux Kernel Variant of OpenIB.org license"
	case SPDXLicenseLPL1_0:
		return "Lucent Public License Version 1.0"
	case SPDXLicenseLPL1_02:
		return "Lucent Public License v1.02"
	case SPDXLicenseLPPL1_0:
		return "LaTeX Project Public License v1.0"
	case SPDXLicenseLPPL1_1:
		return "LaTeX Project Public License v1.1"
	case SPDXLicenseLPPL1_2:
		return "LaTeX Project Public License v1.2"
	case SPDXLicenseLPPL1_3a:
		return "LaTeX Project Public License v1.3a"
	case SPDXLicenseLPPL1_3c:
		return "LaTeX Project Public License v1.3c"
	case SPDXLicenseMakeIndex:
		return "MakeIndex License"
	case SPDXLicenseMirOS:
		return "MirOS License"
	case SPDXLicenseMIT0:
		return "MIT No Attribution"
	case SPDXLicenseMITAdvertising:
		return "Enlightenment License (e16)"
	case SPDXLicenseMITCMU:
		return "CMU License"
	case SPDXLicenseMITEnna:
		return "enna License"
	case SPDXLicenseMITFeh:
		return "feh License"
	case SPDXLicenseMIT:
		return "MIT License"
	case SPDXLicenseMITNFA:
		return "MIT +no-false-attribs license"
	case SPDXLicenseMotosoto:
		return "Motosoto License"
	case SPDXLicenseMpich2:
		return "mpich2 License"
	case SPDXLicenseMPL1_0:
		return "Mozilla Public License 1.0"
	case SPDXLicenseMPL1_1:
		return "Mozilla Public License 1.1"
	case SPDXLicenseMPL2_0NoCopyleftException:
		return "Mozilla Public License 2.0 (no copyleft exception)"
	case SPDXLicenseMPL2_0:
		return "Mozilla Public License 2.0"
	case SPDXLicenseMSPL:
		return "Microsoft Public License"
	case SPDXLicenseMSRL:
		return "Microsoft Reciprocal License"
	case SPDXLicenseMTLL:
		return "Matrix Template Library License"
	case SPDXLicenseMultics:
		return "Multics License"
	case SPDXLicenseMup:
		return "Mup License"
	case SPDXLicenseNASA1_3:
		return "NASA Open Source Agreement 1.3"
	case SPDXLicenseNaumen:
		return "Naumen Public License"
	case SPDXLicenseNBPL1_0:
		return "Net Boolean Public License v1"
	case SPDXLicenseNCSA:
		return "University of Illinois/NCSA Open Source License"
	case SPDXLicenseNetSNMP:
		return "Net-SNMP License"
	case SPDXLicenseNetCDF:
		return "NetCDF license"
	case SPDXLicenseNewsletr:
		return "Newsletr License"
	case SPDXLicenseNGPL:
		return "Nethack General Public License"
	case SPDXLicenseNLOD1_0:
		return "Norwegian Licence for Open Government Data"
	case SPDXLicenseNLPL:
		return "No Limit Public License"
	case SPDXLicenseNokia:
		return "Nokia Open Source License"
	case SPDXLicenseNOSL:
		return "Netizen Open Source License"
	case SPDXLicenseNoweb:
		return "Noweb License"
	case SPDXLicenseNPL1_0:
		return "Netscape Public License v1.0"
	case SPDXLicenseNPL1_1:
		return "Netscape Public License v1.1"
	case SPDXLicenseNPOSL3_0:
		return "Non-Profit Open Software License 3.0"
	case SPDXLicenseNRL:
		return "NRL License"
	case SPDXLicenseNTP:
		return "NTP License"
	case SPDXLicenseOCCTPL:
		return "Open CASCADE Technology Public License"
	case SPDXLicenseOCLC2_0:
		return "OCLC Research Public License 2.0"
	case SPDXLicenseODbL1_0:
		return "ODC Open Database License v1.0"
	case SPDXLicenseOFL1_0:
		return "SIL Open Font License 1.0"
	case SPDXLicenseOFL1_1:
		return "SIL Open Font License 1.1"
	case SPDXLicenseOGTSL:
		return "Open Group Test Suite License"
	case SPDXLicenseOLDAP1_1:
		return "Open LDAP Public License v1.1"
	case SPDXLicenseOLDAP1_2:
		return "Open LDAP Public License v1.2"
	case SPDXLicenseOLDAP1_3:
		return "Open LDAP Public License v1.3"
	case SPDXLicenseOLDAP1_4:
		return "Open LDAP Public License v1.4"
	case SPDXLicenseOLDAP2_0_1:
		return "Open LDAP Public License v2.0.1"
	case SPDXLicenseOLDAP2_0:
		return "Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B)"
	case SPDXLicenseOLDAP2_1:
		return "Open LDAP Public License v2.1"
	case SPDXLicenseOLDAP2_2_1:
		return "Open LDAP Public License v2.2.1"
	case SPDXLicenseOLDAP2_2_2:
		return "Open LDAP Public License 2.2.2"
	case SPDXLicenseOLDAP2_2:
		return "Open LDAP Public License v2.2"
	case SPDXLicenseOLDAP2_3:
		return "Open LDAP Public License v2.3"
	case SPDXLicenseOLDAP2_4:
		return "Open LDAP Public License v2.4"
	case SPDXLicenseOLDAP2_5:
		return "Open LDAP Public License v2.5"
	case SPDXLicenseOLDAP2_6:
		return "Open LDAP Public License v2.6"
	case SPDXLicenseOLDAP2_7:
		return "Open LDAP Public License v2.7"
	case SPDXLicenseOLDAP2_8:
		return "Open LDAP Public License v2.8"
	case SPDXLicenseOML:
		return "Open Market License"
	case SPDXLicenseOpenSSL:
		return "OpenSSL License"
	case SPDXLicenseOPL1_0:
		return "Open Public License v1.0"
	case SPDXLicenseOSETPL2_1:
		return "OSET Public License version 2.1"
	case SPDXLicenseOSL1_0:
		return "Open Software License 1.0"
	case SPDXLicenseOSL1_1:
		return "Open Software License 1.1"
	case SPDXLicenseOSL2_0:
		return "Open Software License 2.0"
	case SPDXLicenseOSL2_1:
		return "Open Software License 2.1"
	case SPDXLicenseOSL3_0:
		return "Open Software License 3.0"
	case SPDXLicensePDDL1_0:
		return "ODC Public Domain Dedication & License 1.0"
	case SPDXLicensePHP3_0:
		return "PHP License v3.0"
	case SPDXLicensePHP3_01:
		return "PHP License v3.01"
	case SPDXLicensePlexus:
		return "Plexus Classworlds License"
	case SPDXLicensePostgreSQL:
		return "PostgreSQL License"
	case SPDXLicensePsfrag:
		return "psfrag License"
	case SPDXLicensePsutils:
		return "psutils License"
	case SPDXLicensePython2_0:
		return "Python License 2.0"
	case SPDXLicenseQhull:
		return "Qhull License"
	case SPDXLicenseQPL1_0:
		return "Q Public License 1.0"
	case SPDXLicenseRdisc:
		return "Rdisc License"
	case SPDXLicenseRHeCos1_1:
		return "Red Hat eCos Public License v1.1"
	case SPDXLicenseRPL1_1:
		return "Reciprocal Public License 1.1"
	case SPDXLicenseRPL1_5:
		return "Reciprocal Public License 1.5"
	case SPDXLicenseRPSL1_0:
		return "RealNetworks Public Source License v1.0"
	case SPDXLicenseRSAMD:
		return "RSA Message-Digest License"
	case SPDXLicenseRSCPL:
		return "Ricoh Source Code Public License"
	case SPDXLicenseRuby:
		return "Ruby License"
	case SPDXLicenseSAXPD:
		return "Sax Public Domain Notice"
	case SPDXLicenseSaxpath:
		return "Saxpath License"
	case SPDXLicenseSCEA:
		return "SCEA Shared Source License"
	case SPDXLicenseSendmail:
		return "Sendmail License"
	case SPDXLicenseSGIB1_0:
		return "SGI Free Software License B v1.0"
	case SPDXLicenseSGIB1_1:
		return "SGI Free Software License B v1.1"
	case SPDXLicenseSGIB2_0:
		return "SGI Free Software License B v2.0"
	case SPDXLicenseSimPL2_0:
		return "Simple Public License 2.0"
	case SPDXLicenseSISSL1_2:
		return "Sun Industry Standards Source License v1.2"
	case SPDXLicenseSISSL:
		return "Sun Industry Standards Source License v1.1"
	case SPDXLicenseSleepycat:
		return "Sleepycat License"
	case SPDXLicenseSMLNJ:
		return "Standard ML of New Jersey License"
	case SPDXLicenseSMPPL:
		return "Secure Messaging Protocol Public License"
	case SPDXLicenseSNIA:
		return "SNIA Public License 1.1"
	case SPDXLicenseSpencer86:
		return "Spencer License 86"
	case SPDXLicenseSpencer94:
		return "Spencer License 94"
	case SPDXLicenseSpencer99:
		return "Spencer License 99"
	case SPDXLicenseSPL1_0:
		return "Sun Public License v1.0"
	case SPDXLicenseSugarCRM1_1_3:
		return "SugarCRM Public License v1.1.3"
	case SPDXLicenseSWL:
		return "Scheme Widget Library (SWL) Software License Agreement"
	case SPDXLicenseTCL:
		return "TCL/TK License"
	case SPDXLicenseTCPWrappers:
		return "TCP Wrappers License"
	case SPDXLicenseTMate:
		return "TMate Open Source License"
	case SPDXLicenseTORQUE1_1:
		return "TORQUE v2.5+ Software License v1.1"
	case SPDXLicenseTOSL:
		return "Trusster Open Source License"
	case SPDXLicenseUnicodeDFS2015:
		return "Unicode License Agreement - Data Files and Software (2015)"
	case SPDXLicenseUnicodeDFS2016:
		return "Unicode License Agreement - Data Files and Software (2016)"
	case SPDXLicenseUnicodeTOU:
		return "Unicode Terms of Use"
	case SPDXLicenseUnlicense:
		return "The Unlicense"
	case SPDXLicenseUPL1_0:
		return "Universal Permissive License v1.0"
	case SPDXLicenseVim:
		return "Vim License"
	case SPDXLicenseVOSTROM:
		return "VOSTROM Public License for Open Source"
	case SPDXLicenseVSL1_0:
		return "Vovida Software License v1.0"
	case SPDXLicenseW3C19980720:
		return "W3C Software Notice and License (1998-07-20)"
	case SPDXLicenseW3C20150513:
		return "W3C Software Notice and Document License (2015-05-13)"
	case SPDXLicenseW3C:
		return "W3C Software Notice and License (2002-12-31)"
	case SPDXLicenseWatcom1_0:
		return "Sybase Open Watcom Public License 1.0"
	case SPDXLicenseWsuipa:
		return "Wsuipa License"
	case SPDXLicenseWTFPL:
		return "Do What The F*ck You Want To Public License"
	case SPDXLicenseX11:
		return "X11 License"
	case SPDXLicenseXerox:
		return "Xerox License"
	case SPDXLicenseXFree861_1:
		return "XFree86 License 1.1"
	case SPDXLicenseXinetd:
		return "xinetd License"
	case SPDXLicenseXnet:
		return "X.Net License"
	case SPDXLicenseXpp:
		return "XPP License"
	case SPDXLicenseXSkat:
		return "XSkat License"
	case SPDXLicenseYPL1_0:
		return "Yahoo! Public License v1.0"
	case SPDXLicenseYPL1_1:
		return "Yahoo! Public License v1.1"
	case SPDXLicenseZed:
		return "Zed License"
	case SPDXLicenseZend2_0:
		return "Zend License v2.0"
	case SPDXLicenseZimbra1_3:
		return "Zimbra Public License v1.3"
	case SPDXLicenseZimbra1_4:
		return "Zimbra Public License v1.4"
	case SPDXLicenseZlibAcknowledgement:
		return "zlib/libpng License with Acknowledgement"
	case SPDXLicenseZlib:
		return "zlib License"
	case SPDXLicenseZPL1_1:
		return "Zope Public License 1.1"
	case SPDXLicenseZPL2_0:
		return "Zope Public License 2.0"
	case SPDXLicenseZPL2_1:
		return "Zope Public License 2.1"
	}
	return "<unknown>"
}

func (code SPDXLicense) Definition() string {
	switch code {
	case SPDXLicenseNotOpenSource:
		return "Not an open source license."
	case SPDXLicense0BSD:
		return "BSD Zero Clause License."
	case SPDXLicenseAAL:
		return "Attribution Assurance License."
	case SPDXLicenseAbstyles:
		return "Abstyles License."
	case SPDXLicenseAdobe2006:
		return "Adobe Systems Incorporated Source Code License Agreement."
	case SPDXLicenseAdobeGlyph:
		return "Adobe Glyph List License."
	case SPDXLicenseADSL:
		return "Amazon Digital Services License."
	case SPDXLicenseAFL1_1:
		return "Academic Free License v1.1."
	case SPDXLicenseAFL1_2:
		return "Academic Free License v1.2."
	case SPDXLicenseAFL2_0:
		return "Academic Free License v2.0."
	case SPDXLicenseAFL2_1:
		return "Academic Free License v2.1."
	case SPDXLicenseAFL3_0:
		return "Academic Free License v3.0."
	case SPDXLicenseAfmparse:
		return "Afmparse License."
	case SPDXLicenseAGPL1_0Only:
		return "Affero General Public License v1.0 only."
	case SPDXLicenseAGPL1_0OrLater:
		return "Affero General Public License v1.0 or later."
	case SPDXLicenseAGPL3_0Only:
		return "GNU Affero General Public License v3.0 only."
	case SPDXLicenseAGPL3_0OrLater:
		return "GNU Affero General Public License v3.0 or later."
	case SPDXLicenseAladdin:
		return "Aladdin Free Public License."
	case SPDXLicenseAMDPLPA:
		return "AMD's plpa_map.c License."
	case SPDXLicenseAML:
		return "Apple MIT License."
	case SPDXLicenseAMPAS:
		return "Academy of Motion Picture Arts and Sciences BSD."
	case SPDXLicenseANTLRPD:
		return "ANTLR Software Rights Notice."
	case SPDXLicenseApache1_0:
		return "Apache License 1.0."
	case SPDXLicenseApache1_1:
		return "Apache License 1.1."
	case SPDXLicenseApache2_0:
		return "Apache License 2.0."
	case SPDXLicenseAPAFML:
		return "Adobe Postscript AFM License."
	case SPDXLicenseAPL1_0:
		return "Adaptive Public License 1.0."
	case SPDXLicenseAPSL1_0:
		return "Apple Public Source License 1.0."
	case SPDXLicenseAPSL1_1:
		return "Apple Public Source License 1.1."
	case SPDXLicenseAPSL1_2:
		return "Apple Public Source License 1.2."
	case SPDXLicenseAPSL2_0:
		return "Apple Public Source License 2.0."
	case SPDXLicenseArtistic1_0Cl8:
		return "Artistic License 1.0 w/clause 8."
	case SPDXLicenseArtistic1_0Perl:
		return "Artistic License 1.0 (Perl)."
	case SPDXLicenseArtistic1_0:
		return "Artistic License 1.0."
	case SPDXLicenseArtistic2_0:
		return "Artistic License 2.0."
	case SPDXLicenseBahyph:
		return "Bahyph License."
	case SPDXLicenseBarr:
		return "Barr License."
	case SPDXLicenseBeerware:
		return "Beerware License."
	case SPDXLicenseBitTorrent1_0:
		return "BitTorrent Open Source License v1.0."
	case SPDXLicenseBitTorrent1_1:
		return "BitTorrent Open Source License v1.1."
	case SPDXLicenseBorceux:
		return "Borceux license."
	case SPDXLicenseBSD1Clause:
		return "BSD 1-Clause License."
	case SPDXLicenseBSD2ClauseFreeBSD:
		return "BSD 2-Clause FreeBSD License."
	case SPDXLicenseBSD2ClauseNetBSD:
		return "BSD 2-Clause NetBSD License."
	case SPDXLicenseBSD2ClausePatent:
		return "BSD-2-Clause Plus Patent License."
	case SPDXLicenseBSD2Clause:
		return "BSD 2-Clause \"Simplified\" License."
	case SPDXLicenseBSD3ClauseAttribution:
		return "BSD with attribution."
	case SPDXLicenseBSD3ClauseClear:
		return "BSD 3-Clause Clear License."
	case SPDXLicenseBSD3ClauseLBNL:
		return "Lawrence Berkeley National Labs BSD variant license."
	case SPDXLicenseBSD3ClauseNoNuclearLicense2014:
		return "BSD 3-Clause No Nuclear License 2014."
	case SPDXLicenseBSD3ClauseNoNuclearLicense:
		return "BSD 3-Clause No Nuclear License."
	case SPDXLicenseBSD3ClauseNoNuclearWarranty:
		return "BSD 3-Clause No Nuclear Warranty."
	case SPDXLicenseBSD3Clause:
		return "BSD 3-Clause \"New\" or \"Revised\" License."
	case SPDXLicenseBSD4ClauseUC:
		return "BSD-4-Clause (University of California-Specific)."
	case SPDXLicenseBSD4Clause:
		return "BSD 4-Clause \"Original\" or \"Old\" License."
	case SPDXLicenseBSDProtection:
		return "BSD Protection License."
	case SPDXLicenseBSDSourceCode:
		return "BSD Source Code Attribution."
	case SPDXLicenseBSL1_0:
		return "Boost Software License 1.0."
	case SPDXLicenseBzip21_0_5:
		return "bzip2 and libbzip2 License v1.0.5."
	case SPDXLicenseBzip21_0_6:
		return "bzip2 and libbzip2 License v1.0.6."
	case SPDXLicenseCaldera:
		return "Caldera License."
	case SPDXLicenseCATOSL1_1:
		return "Computer Associates Trusted Open Source License 1.1."
	case SPDXLicenseCCBY1_0:
		return "Creative Commons Attribution 1.0 Generic."
	case SPDXLicenseCCBY2_0:
		return "Creative Commons Attribution 2.0 Generic."
	case SPDXLicenseCCBY2_5:
		return "Creative Commons Attribution 2.5 Generic."
	case SPDXLicenseCCBY3_0:
		return "Creative Commons Attribution 3.0 Unported."
	case SPDXLicenseCCBY4_0:
		return "Creative Commons Attribution 4.0 International."
	case SPDXLicenseCCBYNC1_0:
		return "Creative Commons Attribution Non Commercial 1.0 Generic."
	case SPDXLicenseCCBYNC2_0:
		return "Creative Commons Attribution Non Commercial 2.0 Generic."
	case SPDXLicenseCCBYNC2_5:
		return "Creative Commons Attribution Non Commercial 2.5 Generic."
	case SPDXLicenseCCBYNC3_0:
		return "Creative Commons Attribution Non Commercial 3.0 Unported."
	case SPDXLicenseCCBYNC4_0:
		return "Creative Commons Attribution Non Commercial 4.0 International."
	case SPDXLicenseCCBYNCND1_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 1.0 Generic."
	case SPDXLicenseCCBYNCND2_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 2.0 Generic."
	case SPDXLicenseCCBYNCND2_5:
		return "Creative Commons Attribution Non Commercial No Derivatives 2.5 Generic."
	case SPDXLicenseCCBYNCND3_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 3.0 Unported."
	case SPDXLicenseCCBYNCND4_0:
		return "Creative Commons Attribution Non Commercial No Derivatives 4.0 International."
	case SPDXLicenseCCBYNCSA1_0:
		return "Creative Commons Attribution Non Commercial Share Alike 1.0 Generic."
	case SPDXLicenseCCBYNCSA2_0:
		return "Creative Commons Attribution Non Commercial Share Alike 2.0 Generic."
	case SPDXLicenseCCBYNCSA2_5:
		return "Creative Commons Attribution Non Commercial Share Alike 2.5 Generic."
	case SPDXLicenseCCBYNCSA3_0:
		return "Creative Commons Attribution Non Commercial Share Alike 3.0 Unported."
	case SPDXLicenseCCBYNCSA4_0:
		return "Creative Commons Attribution Non Commercial Share Alike 4.0 International."
	case SPDXLicenseCCBYND1_0:
		return "Creative Commons Attribution No Derivatives 1.0 Generic."
	case SPDXLicenseCCBYND2_0:
		return "Creative Commons Attribution No Derivatives 2.0 Generic."
	case SPDXLicenseCCBYND2_5:
		return "Creative Commons Attribution No Derivatives 2.5 Generic."
	case SPDXLicenseCCBYND3_0:
		return "Creative Commons Attribution No Derivatives 3.0 Unported."
	case SPDXLicenseCCBYND4_0:
		return "Creative Commons Attribution No Derivatives 4.0 International."
	case SPDXLicenseCCBYSA1_0:
		return "Creative Commons Attribution Share Alike 1.0 Generic."
	case SPDXLicenseCCBYSA2_0:
		return "Creative Commons Attribution Share Alike 2.0 Generic."
	case SPDXLicenseCCBYSA2_5:
		return "Creative Commons Attribution Share Alike 2.5 Generic."
	case SPDXLicenseCCBYSA3_0:
		return "Creative Commons Attribution Share Alike 3.0 Unported."
	case SPDXLicenseCCBYSA4_0:
		return "Creative Commons Attribution Share Alike 4.0 International."
	case SPDXLicenseCC01_0:
		return "Creative Commons Zero v1.0 Universal."
	case SPDXLicenseCDDL1_0:
		return "Common Development and Distribution License 1.0."
	case SPDXLicenseCDDL1_1:
		return "Common Development and Distribution License 1.1."
	case SPDXLicenseCDLAPermissive1_0:
		return "Community Data License Agreement Permissive 1.0."
	case SPDXLicenseCDLASharing1_0:
		return "Community Data License Agreement Sharing 1.0."
	case SPDXLicenseCECILL1_0:
		return "CeCILL Free Software License Agreement v1.0."
	case SPDXLicenseCECILL1_1:
		return "CeCILL Free Software License Agreement v1.1."
	case SPDXLicenseCECILL2_0:
		return "CeCILL Free Software License Agreement v2.0."
	case SPDXLicenseCECILL2_1:
		return "CeCILL Free Software License Agreement v2.1."
	case SPDXLicenseCECILLB:
		return "CeCILL-B Free Software License Agreement."
	case SPDXLicenseCECILLC:
		return "CeCILL-C Free Software License Agreement."
	case SPDXLicenseClArtistic:
		return "Clarified Artistic License."
	case SPDXLicenseCNRIJython:
		return "CNRI Jython License."
	case SPDXLicenseCNRIPythonGPLCompatible:
		return "CNRI Python Open Source GPL Compatible License Agreement."
	case SPDXLicenseCNRIPython:
		return "CNRI Python License."
	case SPDXLicenseCondor1_1:
		return "Condor Public License v1.1."
	case SPDXLicenseCPAL1_0:
		return "Common Public Attribution License 1.0."
	case SPDXLicenseCPL1_0:
		return "Common Public License 1.0."
	case SPDXLicenseCPOL1_02:
		return "Code Project Open License 1.02."
	case SPDXLicenseCrossword:
		return "Crossword License."
	case SPDXLicenseCrystalStacker:
		return "CrystalStacker License."
	case SPDXLicenseCUAOPL1_0:
		return "CUA Office Public License v1.0."
	case SPDXLicenseCube:
		return "Cube License."
	case SPDXLicenseCurl:
		return "curl License."
	case SPDXLicenseDFSL1_0:
		return "Deutsche Freie Software Lizenz."
	case SPDXLicenseDiffmark:
		return "diffmark license."
	case SPDXLicenseDOC:
		return "DOC License."
	case SPDXLicenseDotseqn:
		return "Dotseqn License."
	case SPDXLicenseDSDP:
		return "DSDP License."
	case SPDXLicenseDvipdfm:
		return "dvipdfm License."
	case SPDXLicenseECL1_0:
		return "Educational Community License v1.0."
	case SPDXLicenseECL2_0:
		return "Educational Community License v2.0."
	case SPDXLicenseEFL1_0:
		return "Eiffel Forum License v1.0."
	case SPDXLicenseEFL2_0:
		return "Eiffel Forum License v2.0."
	case SPDXLicenseEGenix:
		return "eGenix.com Public License 1.1.0."
	case SPDXLicenseEntessa:
		return "Entessa Public License v1.0."
	case SPDXLicenseEPL1_0:
		return "Eclipse Public License 1.0."
	case SPDXLicenseEPL2_0:
		return "Eclipse Public License 2.0."
	case SPDXLicenseErlPL1_1:
		return "Erlang Public License v1.1."
	case SPDXLicenseEUDatagrid:
		return "EU DataGrid Software License."
	case SPDXLicenseEUPL1_0:
		return "European Union Public License 1.0."
	case SPDXLicenseEUPL1_1:
		return "European Union Public License 1.1."
	case SPDXLicenseEUPL1_2:
		return "European Union Public License 1.2."
	case SPDXLicenseEurosym:
		return "Eurosym License."
	case SPDXLicenseFair:
		return "Fair License."
	case SPDXLicenseFrameworx1_0:
		return "Frameworx Open License 1.0."
	case SPDXLicenseFreeImage:
		return "FreeImage Public License v1.0."
	case SPDXLicenseFSFAP:
		return "FSF All Permissive License."
	case SPDXLicenseFSFUL:
		return "FSF Unlimited License."
	case SPDXLicenseFSFULLR:
		return "FSF Unlimited License (with License Retention)."
	case SPDXLicenseFTL:
		return "Freetype Project License."
	case SPDXLicenseGFDL1_1Only:
		return "GNU Free Documentation License v1.1 only."
	case SPDXLicenseGFDL1_1OrLater:
		return "GNU Free Documentation License v1.1 or later."
	case SPDXLicenseGFDL1_2Only:
		return "GNU Free Documentation License v1.2 only."
	case SPDXLicenseGFDL1_2OrLater:
		return "GNU Free Documentation License v1.2 or later."
	case SPDXLicenseGFDL1_3Only:
		return "GNU Free Documentation License v1.3 only."
	case SPDXLicenseGFDL1_3OrLater:
		return "GNU Free Documentation License v1.3 or later."
	case SPDXLicenseGiftware:
		return "Giftware License."
	case SPDXLicenseGL2PS:
		return "GL2PS License."
	case SPDXLicenseGlide:
		return "3dfx Glide License."
	case SPDXLicenseGlulxe:
		return "Glulxe License."
	case SPDXLicenseGnuplot:
		return "gnuplot License."
	case SPDXLicenseGPL1_0Only:
		return "GNU General Public License v1.0 only."
	case SPDXLicenseGPL1_0OrLater:
		return "GNU General Public License v1.0 or later."
	case SPDXLicenseGPL2_0Only:
		return "GNU General Public License v2.0 only."
	case SPDXLicenseGPL2_0OrLater:
		return "GNU General Public License v2.0 or later."
	case SPDXLicenseGPL3_0Only:
		return "GNU General Public License v3.0 only."
	case SPDXLicenseGPL3_0OrLater:
		return "GNU General Public License v3.0 or later."
	case SPDXLicenseGSOAP1_3b:
		return "gSOAP Public License v1.3b."
	case SPDXLicenseHaskellReport:
		return "Haskell Language Report License."
	case SPDXLicenseHPND:
		return "Historical Permission Notice and Disclaimer."
	case SPDXLicenseIBMPibs:
		return "IBM PowerPC Initialization and Boot Software."
	case SPDXLicenseICU:
		return "ICU License."
	case SPDXLicenseIJG:
		return "Independent JPEG Group License."
	case SPDXLicenseImageMagick:
		return "ImageMagick License."
	case SPDXLicenseIMatix:
		return "iMatix Standard Function Library Agreement."
	case SPDXLicenseImlib2:
		return "Imlib2 License."
	case SPDXLicenseInfoZIP:
		return "Info-ZIP License."
	case SPDXLicenseIntelACPI:
		return "Intel ACPI Software License Agreement."
	case SPDXLicenseIntel:
		return "Intel Open Source License."
	case SPDXLicenseInterbase1_0:
		return "Interbase Public License v1.0."
	case SPDXLicenseIPA:
		return "IPA Font License."
	case SPDXLicenseIPL1_0:
		return "IBM Public License v1.0."
	case SPDXLicenseISC:
		return "ISC License."
	case SPDXLicenseJasPer2_0:
		return "JasPer License."
	case SPDXLicenseJSON:
		return "JSON License."
	case SPDXLicenseLAL1_2:
		return "Licence Art Libre 1.2."
	case SPDXLicenseLAL1_3:
		return "Licence Art Libre 1.3."
	case SPDXLicenseLatex2e:
		return "Latex2e License."
	case SPDXLicenseLeptonica:
		return "Leptonica License."
	case SPDXLicenseLGPL2_0Only:
		return "GNU Library General Public License v2 only."
	case SPDXLicenseLGPL2_0OrLater:
		return "GNU Library General Public License v2 or later."
	case SPDXLicenseLGPL2_1Only:
		return "GNU Lesser General Public License v2.1 only."
	case SPDXLicenseLGPL2_1OrLater:
		return "GNU Lesser General Public License v2.1 or later."
	case SPDXLicenseLGPL3_0Only:
		return "GNU Lesser General Public License v3.0 only."
	case SPDXLicenseLGPL3_0OrLater:
		return "GNU Lesser General Public License v3.0 or later."
	case SPDXLicenseLGPLLR:
		return "Lesser General Public License For Linguistic Resources."
	case SPDXLicenseLibpng:
		return "libpng License."
	case SPDXLicenseLibtiff:
		return "libtiff License."
	case SPDXLicenseLiLiQP1_1:
		return "Licence Libre du Québec – Permissive version 1.1."
	case SPDXLicenseLiLiQR1_1:
		return "Licence Libre du Québec – Réciprocité version 1.1."
	case SPDXLicenseLiLiQRplus1_1:
		return "Licence Libre du Québec – Réciprocité forte version 1.1."
	case SPDXLicenseLinuxOpenIB:
		return "Linux Kernel Variant of OpenIB.org license."
	case SPDXLicenseLPL1_0:
		return "Lucent Public License Version 1.0."
	case SPDXLicenseLPL1_02:
		return "Lucent Public License v1.02."
	case SPDXLicenseLPPL1_0:
		return "LaTeX Project Public License v1.0."
	case SPDXLicenseLPPL1_1:
		return "LaTeX Project Public License v1.1."
	case SPDXLicenseLPPL1_2:
		return "LaTeX Project Public License v1.2."
	case SPDXLicenseLPPL1_3a:
		return "LaTeX Project Public License v1.3a."
	case SPDXLicenseLPPL1_3c:
		return "LaTeX Project Public License v1.3c."
	case SPDXLicenseMakeIndex:
		return "MakeIndex License."
	case SPDXLicenseMirOS:
		return "MirOS License."
	case SPDXLicenseMIT0:
		return "MIT No Attribution."
	case SPDXLicenseMITAdvertising:
		return "Enlightenment License (e16)."
	case SPDXLicenseMITCMU:
		return "CMU License."
	case SPDXLicenseMITEnna:
		return "enna License."
	case SPDXLicenseMITFeh:
		return "feh License."
	case SPDXLicenseMIT:
		return "MIT License."
	case SPDXLicenseMITNFA:
		return "MIT +no-false-attribs license."
	case SPDXLicenseMotosoto:
		return "Motosoto License."
	case SPDXLicenseMpich2:
		return "mpich2 License."
	case SPDXLicenseMPL1_0:
		return "Mozilla Public License 1.0."
	case SPDXLicenseMPL1_1:
		return "Mozilla Public License 1.1."
	case SPDXLicenseMPL2_0NoCopyleftException:
		return "Mozilla Public License 2.0 (no copyleft exception)."
	case SPDXLicenseMPL2_0:
		return "Mozilla Public License 2.0."
	case SPDXLicenseMSPL:
		return "Microsoft Public License."
	case SPDXLicenseMSRL:
		return "Microsoft Reciprocal License."
	case SPDXLicenseMTLL:
		return "Matrix Template Library License."
	case SPDXLicenseMultics:
		return "Multics License."
	case SPDXLicenseMup:
		return "Mup License."
	case SPDXLicenseNASA1_3:
		return "NASA Open Source Agreement 1.3."
	case SPDXLicenseNaumen:
		return "Naumen Public License."
	case SPDXLicenseNBPL1_0:
		return "Net Boolean Public License v1."
	case SPDXLicenseNCSA:
		return "University of Illinois/NCSA Open Source License."
	case SPDXLicenseNetSNMP:
		return "Net-SNMP License."
	case SPDXLicenseNetCDF:
		return "NetCDF license."
	case SPDXLicenseNewsletr:
		return "Newsletr License."
	case SPDXLicenseNGPL:
		return "Nethack General Public License."
	case SPDXLicenseNLOD1_0:
		return "Norwegian Licence for Open Government Data."
	case SPDXLicenseNLPL:
		return "No Limit Public License."
	case SPDXLicenseNokia:
		return "Nokia Open Source License."
	case SPDXLicenseNOSL:
		return "Netizen Open Source License."
	case SPDXLicenseNoweb:
		return "Noweb License."
	case SPDXLicenseNPL1_0:
		return "Netscape Public License v1.0."
	case SPDXLicenseNPL1_1:
		return "Netscape Public License v1.1."
	case SPDXLicenseNPOSL3_0:
		return "Non-Profit Open Software License 3.0."
	case SPDXLicenseNRL:
		return "NRL License."
	case SPDXLicenseNTP:
		return "NTP License."
	case SPDXLicenseOCCTPL:
		return "Open CASCADE Technology Public License."
	case SPDXLicenseOCLC2_0:
		return "OCLC Research Public License 2.0."
	case SPDXLicenseODbL1_0:
		return "ODC Open Database License v1.0."
	case SPDXLicenseOFL1_0:
		return "SIL Open Font License 1.0."
	case SPDXLicenseOFL1_1:
		return "SIL Open Font License 1.1."
	case SPDXLicenseOGTSL:
		return "Open Group Test Suite License."
	case SPDXLicenseOLDAP1_1:
		return "Open LDAP Public License v1.1."
	case SPDXLicenseOLDAP1_2:
		return "Open LDAP Public License v1.2."
	case SPDXLicenseOLDAP1_3:
		return "Open LDAP Public License v1.3."
	case SPDXLicenseOLDAP1_4:
		return "Open LDAP Public License v1.4."
	case SPDXLicenseOLDAP2_0_1:
		return "Open LDAP Public License v2.0.1."
	case SPDXLicenseOLDAP2_0:
		return "Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B)."
	case SPDXLicenseOLDAP2_1:
		return "Open LDAP Public License v2.1."
	case SPDXLicenseOLDAP2_2_1:
		return "Open LDAP Public License v2.2.1."
	case SPDXLicenseOLDAP2_2_2:
		return "Open LDAP Public License 2.2.2."
	case SPDXLicenseOLDAP2_2:
		return "Open LDAP Public License v2.2."
	case SPDXLicenseOLDAP2_3:
		return "Open LDAP Public License v2.3."
	case SPDXLicenseOLDAP2_4:
		return "Open LDAP Public License v2.4."
	case SPDXLicenseOLDAP2_5:
		return "Open LDAP Public License v2.5."
	case SPDXLicenseOLDAP2_6:
		return "Open LDAP Public License v2.6."
	case SPDXLicenseOLDAP2_7:
		return "Open LDAP Public License v2.7."
	case SPDXLicenseOLDAP2_8:
		return "Open LDAP Public License v2.8."
	case SPDXLicenseOML:
		return "Open Market License."
	case SPDXLicenseOpenSSL:
		return "OpenSSL License."
	case SPDXLicenseOPL1_0:
		return "Open Public License v1.0."
	case SPDXLicenseOSETPL2_1:
		return "OSET Public License version 2.1."
	case SPDXLicenseOSL1_0:
		return "Open Software License 1.0."
	case SPDXLicenseOSL1_1:
		return "Open Software License 1.1."
	case SPDXLicenseOSL2_0:
		return "Open Software License 2.0."
	case SPDXLicenseOSL2_1:
		return "Open Software License 2.1."
	case SPDXLicenseOSL3_0:
		return "Open Software License 3.0."
	case SPDXLicensePDDL1_0:
		return "ODC Public Domain Dedication & License 1.0."
	case SPDXLicensePHP3_0:
		return "PHP License v3.0."
	case SPDXLicensePHP3_01:
		return "PHP License v3.01."
	case SPDXLicensePlexus:
		return "Plexus Classworlds License."
	case SPDXLicensePostgreSQL:
		return "PostgreSQL License."
	case SPDXLicensePsfrag:
		return "psfrag License."
	case SPDXLicensePsutils:
		return "psutils License."
	case SPDXLicensePython2_0:
		return "Python License 2.0."
	case SPDXLicenseQhull:
		return "Qhull License."
	case SPDXLicenseQPL1_0:
		return "Q Public License 1.0."
	case SPDXLicenseRdisc:
		return "Rdisc License."
	case SPDXLicenseRHeCos1_1:
		return "Red Hat eCos Public License v1.1."
	case SPDXLicenseRPL1_1:
		return "Reciprocal Public License 1.1."
	case SPDXLicenseRPL1_5:
		return "Reciprocal Public License 1.5."
	case SPDXLicenseRPSL1_0:
		return "RealNetworks Public Source License v1.0."
	case SPDXLicenseRSAMD:
		return "RSA Message-Digest License."
	case SPDXLicenseRSCPL:
		return "Ricoh Source Code Public License."
	case SPDXLicenseRuby:
		return "Ruby License."
	case SPDXLicenseSAXPD:
		return "Sax Public Domain Notice."
	case SPDXLicenseSaxpath:
		return "Saxpath License."
	case SPDXLicenseSCEA:
		return "SCEA Shared Source License."
	case SPDXLicenseSendmail:
		return "Sendmail License."
	case SPDXLicenseSGIB1_0:
		return "SGI Free Software License B v1.0."
	case SPDXLicenseSGIB1_1:
		return "SGI Free Software License B v1.1."
	case SPDXLicenseSGIB2_0:
		return "SGI Free Software License B v2.0."
	case SPDXLicenseSimPL2_0:
		return "Simple Public License 2.0."
	case SPDXLicenseSISSL1_2:
		return "Sun Industry Standards Source License v1.2."
	case SPDXLicenseSISSL:
		return "Sun Industry Standards Source License v1.1."
	case SPDXLicenseSleepycat:
		return "Sleepycat License."
	case SPDXLicenseSMLNJ:
		return "Standard ML of New Jersey License."
	case SPDXLicenseSMPPL:
		return "Secure Messaging Protocol Public License."
	case SPDXLicenseSNIA:
		return "SNIA Public License 1.1."
	case SPDXLicenseSpencer86:
		return "Spencer License 86."
	case SPDXLicenseSpencer94:
		return "Spencer License 94."
	case SPDXLicenseSpencer99:
		return "Spencer License 99."
	case SPDXLicenseSPL1_0:
		return "Sun Public License v1.0."
	case SPDXLicenseSugarCRM1_1_3:
		return "SugarCRM Public License v1.1.3."
	case SPDXLicenseSWL:
		return "Scheme Widget Library (SWL) Software License Agreement."
	case SPDXLicenseTCL:
		return "TCL/TK License."
	case SPDXLicenseTCPWrappers:
		return "TCP Wrappers License."
	case SPDXLicenseTMate:
		return "TMate Open Source License."
	case SPDXLicenseTORQUE1_1:
		return "TORQUE v2.5+ Software License v1.1."
	case SPDXLicenseTOSL:
		return "Trusster Open Source License."
	case SPDXLicenseUnicodeDFS2015:
		return "Unicode License Agreement - Data Files and Software (2015)."
	case SPDXLicenseUnicodeDFS2016:
		return "Unicode License Agreement - Data Files and Software (2016)."
	case SPDXLicenseUnicodeTOU:
		return "Unicode Terms of Use."
	case SPDXLicenseUnlicense:
		return "The Unlicense."
	case SPDXLicenseUPL1_0:
		return "Universal Permissive License v1.0."
	case SPDXLicenseVim:
		return "Vim License."
	case SPDXLicenseVOSTROM:
		return "VOSTROM Public License for Open Source."
	case SPDXLicenseVSL1_0:
		return "Vovida Software License v1.0."
	case SPDXLicenseW3C19980720:
		return "W3C Software Notice and License (1998-07-20)."
	case SPDXLicenseW3C20150513:
		return "W3C Software Notice and Document License (2015-05-13)."
	case SPDXLicenseW3C:
		return "W3C Software Notice and License (2002-12-31)."
	case SPDXLicenseWatcom1_0:
		return "Sybase Open Watcom Public License 1.0."
	case SPDXLicenseWsuipa:
		return "Wsuipa License."
	case SPDXLicenseWTFPL:
		return "Do What The F*ck You Want To Public License."
	case SPDXLicenseX11:
		return "X11 License."
	case SPDXLicenseXerox:
		return "Xerox License."
	case SPDXLicenseXFree861_1:
		return "XFree86 License 1.1."
	case SPDXLicenseXinetd:
		return "xinetd License."
	case SPDXLicenseXnet:
		return "X.Net License."
	case SPDXLicenseXpp:
		return "XPP License."
	case SPDXLicenseXSkat:
		return "XSkat License."
	case SPDXLicenseYPL1_0:
		return "Yahoo! Public License v1.0."
	case SPDXLicenseYPL1_1:
		return "Yahoo! Public License v1.1."
	case SPDXLicenseZed:
		return "Zed License."
	case SPDXLicenseZend2_0:
		return "Zend License v2.0."
	case SPDXLicenseZimbra1_3:
		return "Zimbra Public License v1.3."
	case SPDXLicenseZimbra1_4:
		return "Zimbra Public License v1.4."
	case SPDXLicenseZlibAcknowledgement:
		return "zlib/libpng License with Acknowledgement."
	case SPDXLicenseZlib:
		return "zlib License."
	case SPDXLicenseZPL1_1:
		return "Zope Public License 1.1."
	case SPDXLicenseZPL2_0:
		return "Zope Public License 2.0."
	case SPDXLicenseZPL2_1:
		return "Zope Public License 2.1."
	}
	return "<unknown>"
}
