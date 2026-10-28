package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// The license that applies to an Implementation Guide (using an SPDX license Identifiers, or 'not-open-source').
//
// The binding is required but new SPDX license Identifiers are allowed to be used (https://spdx.org/licenses/).
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/spdx-license
type SPDXLicense int

const (
	SPDXLicenseNotOpenSource SPDXLicense = iota
	// BSD Zero Clause License.
	SPDXLicense0BSD
	// Attribution Assurance License.
	SPDXLicenseAAL
	// Abstyles License.
	SPDXLicenseAbstyles
	// Adobe Systems Incorporated Source Code License Agreement.
	SPDXLicenseAdobe2006
	// Adobe Glyph List License.
	SPDXLicenseAdobeGlyph
	// Amazon Digital Services License.
	SPDXLicenseADSL
	// Academic Free License v1.1.
	SPDXLicenseAFL1_1
	// Academic Free License v1.2.
	SPDXLicenseAFL1_2
	// Academic Free License json.0.
	SPDXLicenseAFL2_0
	// Academic Free License json.1.
	SPDXLicenseAFL2_1
	// Academic Free License v3.0.
	SPDXLicenseAFL3_0
	// Afmparse License.
	SPDXLicenseAfmparse
	// Affero General Public License v1.0 only.
	SPDXLicenseAGPL1_0Only
	// Affero General Public License v1.0 or later.
	SPDXLicenseAGPL1_0OrLater
	// GNU Affero General Public License v3.0 only.
	SPDXLicenseAGPL3_0Only
	// GNU Affero General Public License v3.0 or later.
	SPDXLicenseAGPL3_0OrLater
	// Aladdin Free Public License.
	SPDXLicenseAladdin
	// AMD's plpa_map.c License.
	SPDXLicenseAMDPLPA
	// Apple MIT License.
	SPDXLicenseAML
	// Academy of Motion Picture Arts and Sciences BSD.
	SPDXLicenseAMPAS
	// ANTLR Software Rights Notice.
	SPDXLicenseANTLRPD
	// Apache License 1.0.
	SPDXLicenseApache1_0
	// Apache License 1.1.
	SPDXLicenseApache1_1
	// Apache License 2.0.
	SPDXLicenseApache2_0
	// Adobe Postscript AFM License.
	SPDXLicenseAPAFML
	// Adaptive Public License 1.0.
	SPDXLicenseAPL1_0
	// Apple Public Source License 1.0.
	SPDXLicenseAPSL1_0
	// Apple Public Source License 1.1.
	SPDXLicenseAPSL1_1
	// Apple Public Source License 1.2.
	SPDXLicenseAPSL1_2
	// Apple Public Source License 2.0.
	SPDXLicenseAPSL2_0
	// Artistic License 1.0 w/clause 8.
	SPDXLicenseArtistic1_0Cl8
	// Artistic License 1.0 (Perl).
	SPDXLicenseArtistic1_0Perl
	// Artistic License 1.0.
	SPDXLicenseArtistic1_0
	// Artistic License 2.0.
	SPDXLicenseArtistic2_0
	// Bahyph License.
	SPDXLicenseBahyph
	// Barr License.
	SPDXLicenseBarr
	// Beerware License.
	SPDXLicenseBeerware
	// BitTorrent Open Source License v1.0.
	SPDXLicenseBitTorrent1_0
	// BitTorrent Open Source License v1.1.
	SPDXLicenseBitTorrent1_1
	// Borceux license.
	SPDXLicenseBorceux
	// BSD 1-Clause License.
	SPDXLicenseBSD1Clause
	// BSD 2-Clause FreeBSD License.
	SPDXLicenseBSD2ClauseFreeBSD
	// BSD 2-Clause NetBSD License.
	SPDXLicenseBSD2ClauseNetBSD
	// BSD-2-Clause Plus Patent License.
	SPDXLicenseBSD2ClausePatent
	// BSD 2-Clause "Simplified" License.
	SPDXLicenseBSD2Clause
	// BSD with attribution.
	SPDXLicenseBSD3ClauseAttribution
	// BSD 3-Clause Clear License.
	SPDXLicenseBSD3ClauseClear
	// Lawrence Berkeley National Labs BSD variant license.
	SPDXLicenseBSD3ClauseLBNL
	// BSD 3-Clause No Nuclear License 2014.
	SPDXLicenseBSD3ClauseNoNuclearLicense2014
	// BSD 3-Clause No Nuclear License.
	SPDXLicenseBSD3ClauseNoNuclearLicense
	// BSD 3-Clause No Nuclear Warranty.
	SPDXLicenseBSD3ClauseNoNuclearWarranty
	// BSD 3-Clause "New" or "Revised" License.
	SPDXLicenseBSD3Clause
	// BSD-4-Clause (University of California-Specific).
	SPDXLicenseBSD4ClauseUC
	// BSD 4-Clause "Original" or "Old" License.
	SPDXLicenseBSD4Clause
	// BSD Protection License.
	SPDXLicenseBSDProtection
	// BSD Source Code Attribution.
	SPDXLicenseBSDSourceCode
	// Boost Software License 1.0.
	SPDXLicenseBSL1_0
	// bzip2 and libbzip2 License v1.0.5.
	SPDXLicenseBzip21_0_5
	// bzip2 and libbzip2 License v1.0.6.
	SPDXLicenseBzip21_0_6
	// Caldera License.
	SPDXLicenseCaldera
	// Computer Associates Trusted Open Source License 1.1.
	SPDXLicenseCATOSL1_1
	// Creative Commons Attribution 1.0 Generic.
	SPDXLicenseCCBY1_0
	// Creative Commons Attribution 2.0 Generic.
	SPDXLicenseCCBY2_0
	// Creative Commons Attribution 2.5 Generic.
	SPDXLicenseCCBY2_5
	// Creative Commons Attribution 3.0 Unported.
	SPDXLicenseCCBY3_0
	// Creative Commons Attribution 4.0 International.
	SPDXLicenseCCBY4_0
	// Creative Commons Attribution Non Commercial 1.0 Generic.
	SPDXLicenseCCBYNC1_0
	// Creative Commons Attribution Non Commercial 2.0 Generic.
	SPDXLicenseCCBYNC2_0
	// Creative Commons Attribution Non Commercial 2.5 Generic.
	SPDXLicenseCCBYNC2_5
	// Creative Commons Attribution Non Commercial 3.0 Unported.
	SPDXLicenseCCBYNC3_0
	// Creative Commons Attribution Non Commercial 4.0 International.
	SPDXLicenseCCBYNC4_0
	// Creative Commons Attribution Non Commercial No Derivatives 1.0 Generic.
	SPDXLicenseCCBYNCND1_0
	// Creative Commons Attribution Non Commercial No Derivatives 2.0 Generic.
	SPDXLicenseCCBYNCND2_0
	// Creative Commons Attribution Non Commercial No Derivatives 2.5 Generic.
	SPDXLicenseCCBYNCND2_5
	// Creative Commons Attribution Non Commercial No Derivatives 3.0 Unported.
	SPDXLicenseCCBYNCND3_0
	// Creative Commons Attribution Non Commercial No Derivatives 4.0 International.
	SPDXLicenseCCBYNCND4_0
	// Creative Commons Attribution Non Commercial Share Alike 1.0 Generic.
	SPDXLicenseCCBYNCSA1_0
	// Creative Commons Attribution Non Commercial Share Alike 2.0 Generic.
	SPDXLicenseCCBYNCSA2_0
	// Creative Commons Attribution Non Commercial Share Alike 2.5 Generic.
	SPDXLicenseCCBYNCSA2_5
	// Creative Commons Attribution Non Commercial Share Alike 3.0 Unported.
	SPDXLicenseCCBYNCSA3_0
	// Creative Commons Attribution Non Commercial Share Alike 4.0 International.
	SPDXLicenseCCBYNCSA4_0
	// Creative Commons Attribution No Derivatives 1.0 Generic.
	SPDXLicenseCCBYND1_0
	// Creative Commons Attribution No Derivatives 2.0 Generic.
	SPDXLicenseCCBYND2_0
	// Creative Commons Attribution No Derivatives 2.5 Generic.
	SPDXLicenseCCBYND2_5
	// Creative Commons Attribution No Derivatives 3.0 Unported.
	SPDXLicenseCCBYND3_0
	// Creative Commons Attribution No Derivatives 4.0 International.
	SPDXLicenseCCBYND4_0
	// Creative Commons Attribution Share Alike 1.0 Generic.
	SPDXLicenseCCBYSA1_0
	// Creative Commons Attribution Share Alike 2.0 Generic.
	SPDXLicenseCCBYSA2_0
	// Creative Commons Attribution Share Alike 2.5 Generic.
	SPDXLicenseCCBYSA2_5
	// Creative Commons Attribution Share Alike 3.0 Unported.
	SPDXLicenseCCBYSA3_0
	// Creative Commons Attribution Share Alike 4.0 International.
	SPDXLicenseCCBYSA4_0
	// Creative Commons Zero v1.0 Universal.
	SPDXLicenseCC01_0
	// Common Development and Distribution License 1.0.
	SPDXLicenseCDDL1_0
	// Common Development and Distribution License 1.1.
	SPDXLicenseCDDL1_1
	// Community Data License Agreement Permissive 1.0.
	SPDXLicenseCDLAPermissive1_0
	// Community Data License Agreement Sharing 1.0.
	SPDXLicenseCDLASharing1_0
	// CeCILL Free Software License Agreement v1.0.
	SPDXLicenseCECILL1_0
	// CeCILL Free Software License Agreement v1.1.
	SPDXLicenseCECILL1_1
	// CeCILL Free Software License Agreement json.0.
	SPDXLicenseCECILL2_0
	// CeCILL Free Software License Agreement json.1.
	SPDXLicenseCECILL2_1
	// CeCILL-B Free Software License Agreement.
	SPDXLicenseCECILLB
	// CeCILL-C Free Software License Agreement.
	SPDXLicenseCECILLC
	// Clarified Artistic License.
	SPDXLicenseClArtistic
	// CNRI Jython License.
	SPDXLicenseCNRIJython
	// CNRI Python Open Source GPL Compatible License Agreement.
	SPDXLicenseCNRIPythonGPLCompatible
	// CNRI Python License.
	SPDXLicenseCNRIPython
	// Condor Public License v1.1.
	SPDXLicenseCondor1_1
	// Common Public Attribution License 1.0.
	SPDXLicenseCPAL1_0
	// Common Public License 1.0.
	SPDXLicenseCPL1_0
	// Code Project Open License 1.02.
	SPDXLicenseCPOL1_02
	// Crossword License.
	SPDXLicenseCrossword
	// CrystalStacker License.
	SPDXLicenseCrystalStacker
	// CUA Office Public License v1.0.
	SPDXLicenseCUAOPL1_0
	// Cube License.
	SPDXLicenseCube
	// curl License.
	SPDXLicenseCurl
	// Deutsche Freie Software Lizenz.
	SPDXLicenseDFSL1_0
	// diffmark license.
	SPDXLicenseDiffmark
	// DOC License.
	SPDXLicenseDOC
	// Dotseqn License.
	SPDXLicenseDotseqn
	// DSDP License.
	SPDXLicenseDSDP
	// dvipdfm License.
	SPDXLicenseDvipdfm
	// Educational Community License v1.0.
	SPDXLicenseECL1_0
	// Educational Community License json.0.
	SPDXLicenseECL2_0
	// Eiffel Forum License v1.0.
	SPDXLicenseEFL1_0
	// Eiffel Forum License json.0.
	SPDXLicenseEFL2_0
	// eGenix.com Public License 1.1.0.
	SPDXLicenseEGenix
	// Entessa Public License v1.0.
	SPDXLicenseEntessa
	// Eclipse Public License 1.0.
	SPDXLicenseEPL1_0
	// Eclipse Public License 2.0.
	SPDXLicenseEPL2_0
	// Erlang Public License v1.1.
	SPDXLicenseErlPL1_1
	// EU DataGrid Software License.
	SPDXLicenseEUDatagrid
	// European Union Public License 1.0.
	SPDXLicenseEUPL1_0
	// European Union Public License 1.1.
	SPDXLicenseEUPL1_1
	// European Union Public License 1.2.
	SPDXLicenseEUPL1_2
	// Eurosym License.
	SPDXLicenseEurosym
	// Fair License.
	SPDXLicenseFair
	// Frameworx Open License 1.0.
	SPDXLicenseFrameworx1_0
	// FreeImage Public License v1.0.
	SPDXLicenseFreeImage
	// FSF All Permissive License.
	SPDXLicenseFSFAP
	// FSF Unlimited License.
	SPDXLicenseFSFUL
	// FSF Unlimited License (with License Retention).
	SPDXLicenseFSFULLR
	// Freetype Project License.
	SPDXLicenseFTL
	// GNU Free Documentation License v1.1 only.
	SPDXLicenseGFDL1_1Only
	// GNU Free Documentation License v1.1 or later.
	SPDXLicenseGFDL1_1OrLater
	// GNU Free Documentation License v1.2 only.
	SPDXLicenseGFDL1_2Only
	// GNU Free Documentation License v1.2 or later.
	SPDXLicenseGFDL1_2OrLater
	// GNU Free Documentation License v1.3 only.
	SPDXLicenseGFDL1_3Only
	// GNU Free Documentation License v1.3 or later.
	SPDXLicenseGFDL1_3OrLater
	// Giftware License.
	SPDXLicenseGiftware
	// GL2PS License.
	SPDXLicenseGL2PS
	// 3dfx Glide License.
	SPDXLicenseGlide
	// Glulxe License.
	SPDXLicenseGlulxe
	// gnuplot License.
	SPDXLicenseGnuplot
	// GNU General Public License v1.0 only.
	SPDXLicenseGPL1_0Only
	// GNU General Public License v1.0 or later.
	SPDXLicenseGPL1_0OrLater
	// GNU General Public License json.0 only.
	SPDXLicenseGPL2_0Only
	// GNU General Public License json.0 or later.
	SPDXLicenseGPL2_0OrLater
	// GNU General Public License v3.0 only.
	SPDXLicenseGPL3_0Only
	// GNU General Public License v3.0 or later.
	SPDXLicenseGPL3_0OrLater
	// gSOAP Public License v1.3b.
	SPDXLicenseGSOAP1_3b
	// Haskell Language Report License.
	SPDXLicenseHaskellReport
	// Historical Permission Notice and Disclaimer.
	SPDXLicenseHPND
	// IBM PowerPC Initialization and Boot Software.
	SPDXLicenseIBMPibs
	// ICU License.
	SPDXLicenseICU
	// Independent JPEG Group License.
	SPDXLicenseIJG
	// ImageMagick License.
	SPDXLicenseImageMagick
	// iMatix Standard Function Library Agreement.
	SPDXLicenseIMatix
	// Imlib2 License.
	SPDXLicenseImlib2
	// Info-ZIP License.
	SPDXLicenseInfoZIP
	// Intel ACPI Software License Agreement.
	SPDXLicenseIntelACPI
	// Intel Open Source License.
	SPDXLicenseIntel
	// Interbase Public License v1.0.
	SPDXLicenseInterbase1_0
	// IPA Font License.
	SPDXLicenseIPA
	// IBM Public License v1.0.
	SPDXLicenseIPL1_0
	// ISC License.
	SPDXLicenseISC
	// JasPer License.
	SPDXLicenseJasPer2_0
	// JSON License.
	SPDXLicenseJSON
	// Licence Art Libre 1.2.
	SPDXLicenseLAL1_2
	// Licence Art Libre 1.3.
	SPDXLicenseLAL1_3
	// Latex2e License.
	SPDXLicenseLatex2e
	// Leptonica License.
	SPDXLicenseLeptonica
	// GNU Library General Public License json.only.
	SPDXLicenseLGPL2_0Only
	// GNU Library General Public License json.or later.
	SPDXLicenseLGPL2_0OrLater
	// GNU Lesser General Public License json.1 only.
	SPDXLicenseLGPL2_1Only
	// GNU Lesser General Public License json.1 or later.
	SPDXLicenseLGPL2_1OrLater
	// GNU Lesser General Public License v3.0 only.
	SPDXLicenseLGPL3_0Only
	// GNU Lesser General Public License v3.0 or later.
	SPDXLicenseLGPL3_0OrLater
	// Lesser General Public License For Linguistic Resources.
	SPDXLicenseLGPLLR
	// libpng License.
	SPDXLicenseLibpng
	// libtiff License.
	SPDXLicenseLibtiff
	// Licence Libre du Québec – Permissive version 1.1.
	SPDXLicenseLiLiQP1_1
	// Licence Libre du Québec – Réciprocité version 1.1.
	SPDXLicenseLiLiQR1_1
	// Licence Libre du Québec – Réciprocité forte version 1.1.
	SPDXLicenseLiLiQRplus1_1
	// Linux Kernel Variant of OpenIB.org license.
	SPDXLicenseLinuxOpenIB
	// Lucent Public License Version 1.0.
	SPDXLicenseLPL1_0
	// Lucent Public License v1.02.
	SPDXLicenseLPL1_02
	// LaTeX Project Public License v1.0.
	SPDXLicenseLPPL1_0
	// LaTeX Project Public License v1.1.
	SPDXLicenseLPPL1_1
	// LaTeX Project Public License v1.2.
	SPDXLicenseLPPL1_2
	// LaTeX Project Public License v1.3a.
	SPDXLicenseLPPL1_3a
	// LaTeX Project Public License v1.3c.
	SPDXLicenseLPPL1_3c
	// MakeIndex License.
	SPDXLicenseMakeIndex
	// MirOS License.
	SPDXLicenseMirOS
	// MIT No Attribution.
	SPDXLicenseMIT0
	// Enlightenment License (e16).
	SPDXLicenseMITAdvertising
	// CMU License.
	SPDXLicenseMITCMU
	// enna License.
	SPDXLicenseMITEnna
	// feh License.
	SPDXLicenseMITFeh
	// MIT License.
	SPDXLicenseMIT
	// MIT +no-false-attribs license.
	SPDXLicenseMITNFA
	// Motosoto License.
	SPDXLicenseMotosoto
	// mpich2 License.
	SPDXLicenseMpich2
	// Mozilla Public License 1.0.
	SPDXLicenseMPL1_0
	// Mozilla Public License 1.1.
	SPDXLicenseMPL1_1
	// Mozilla Public License 2.0 (no copyleft exception).
	SPDXLicenseMPL2_0NoCopyleftException
	// Mozilla Public License 2.0.
	SPDXLicenseMPL2_0
	// Microsoft Public License.
	SPDXLicenseMSPL
	// Microsoft Reciprocal License.
	SPDXLicenseMSRL
	// Matrix Template Library License.
	SPDXLicenseMTLL
	// Multics License.
	SPDXLicenseMultics
	// Mup License.
	SPDXLicenseMup
	// NASA Open Source Agreement 1.3.
	SPDXLicenseNASA1_3
	// Naumen Public License.
	SPDXLicenseNaumen
	// Net Boolean Public License v1.
	SPDXLicenseNBPL1_0
	// University of Illinois/NCSA Open Source License.
	SPDXLicenseNCSA
	// Net-SNMP License.
	SPDXLicenseNetSNMP
	// NetCDF license.
	SPDXLicenseNetCDF
	// Newsletr License.
	SPDXLicenseNewsletr
	// Nethack General Public License.
	SPDXLicenseNGPL
	// Norwegian Licence for Open Government Data.
	SPDXLicenseNLOD1_0
	// No Limit Public License.
	SPDXLicenseNLPL
	// Nokia Open Source License.
	SPDXLicenseNokia
	// Netizen Open Source License.
	SPDXLicenseNOSL
	// Noweb License.
	SPDXLicenseNoweb
	// Netscape Public License v1.0.
	SPDXLicenseNPL1_0
	// Netscape Public License v1.1.
	SPDXLicenseNPL1_1
	// Non-Profit Open Software License 3.0.
	SPDXLicenseNPOSL3_0
	// NRL License.
	SPDXLicenseNRL
	// NTP License.
	SPDXLicenseNTP
	// Open CASCADE Technology Public License.
	SPDXLicenseOCCTPL
	// OCLC Research Public License 2.0.
	SPDXLicenseOCLC2_0
	// ODC Open Database License v1.0.
	SPDXLicenseODbL1_0
	// SIL Open Font License 1.0.
	SPDXLicenseOFL1_0
	// SIL Open Font License 1.1.
	SPDXLicenseOFL1_1
	// Open Group Test Suite License.
	SPDXLicenseOGTSL
	// Open LDAP Public License v1.1.
	SPDXLicenseOLDAP1_1
	// Open LDAP Public License v1.2.
	SPDXLicenseOLDAP1_2
	// Open LDAP Public License v1.3.
	SPDXLicenseOLDAP1_3
	// Open LDAP Public License v1.4.
	SPDXLicenseOLDAP1_4
	// Open LDAP Public License json.0.1.
	SPDXLicenseOLDAP2_0_1
	// Open LDAP Public License json.0 (or possibly 2.0A and 2.0B).
	SPDXLicenseOLDAP2_0
	// Open LDAP Public License json.1.
	SPDXLicenseOLDAP2_1
	// Open LDAP Public License json.2.1.
	SPDXLicenseOLDAP2_2_1
	// Open LDAP Public License 2.2.2.
	SPDXLicenseOLDAP2_2_2
	// Open LDAP Public License json.2.
	SPDXLicenseOLDAP2_2
	// Open LDAP Public License json.3.
	SPDXLicenseOLDAP2_3
	// Open LDAP Public License json.4.
	SPDXLicenseOLDAP2_4
	// Open LDAP Public License json.5.
	SPDXLicenseOLDAP2_5
	// Open LDAP Public License json.6.
	SPDXLicenseOLDAP2_6
	// Open LDAP Public License json.7.
	SPDXLicenseOLDAP2_7
	// Open LDAP Public License json.8.
	SPDXLicenseOLDAP2_8
	// Open Market License.
	SPDXLicenseOML
	// OpenSSL License.
	SPDXLicenseOpenSSL
	// Open Public License v1.0.
	SPDXLicenseOPL1_0
	// OSET Public License version 2.1.
	SPDXLicenseOSETPL2_1
	// Open Software License 1.0.
	SPDXLicenseOSL1_0
	// Open Software License 1.1.
	SPDXLicenseOSL1_1
	// Open Software License 2.0.
	SPDXLicenseOSL2_0
	// Open Software License 2.1.
	SPDXLicenseOSL2_1
	// Open Software License 3.0.
	SPDXLicenseOSL3_0
	// ODC Public Domain Dedication & License 1.0.
	SPDXLicensePDDL1_0
	// PHP License v3.0.
	SPDXLicensePHP3_0
	// PHP License v3.01.
	SPDXLicensePHP3_01
	// Plexus Classworlds License.
	SPDXLicensePlexus
	// PostgreSQL License.
	SPDXLicensePostgreSQL
	// psfrag License.
	SPDXLicensePsfrag
	// psutils License.
	SPDXLicensePsutils
	// Python License 2.0.
	SPDXLicensePython2_0
	// Qhull License.
	SPDXLicenseQhull
	// Q Public License 1.0.
	SPDXLicenseQPL1_0
	// Rdisc License.
	SPDXLicenseRdisc
	// Red Hat eCos Public License v1.1.
	SPDXLicenseRHeCos1_1
	// Reciprocal Public License 1.1.
	SPDXLicenseRPL1_1
	// Reciprocal Public License 1.5.
	SPDXLicenseRPL1_5
	// RealNetworks Public Source License v1.0.
	SPDXLicenseRPSL1_0
	// RSA Message-Digest License.
	SPDXLicenseRSAMD
	// Ricoh Source Code Public License.
	SPDXLicenseRSCPL
	// Ruby License.
	SPDXLicenseRuby
	// Sax Public Domain Notice.
	SPDXLicenseSAXPD
	// Saxpath License.
	SPDXLicenseSaxpath
	// SCEA Shared Source License.
	SPDXLicenseSCEA
	// Sendmail License.
	SPDXLicenseSendmail
	// SGI Free Software License B v1.0.
	SPDXLicenseSGIB1_0
	// SGI Free Software License B v1.1.
	SPDXLicenseSGIB1_1
	// SGI Free Software License B json.0.
	SPDXLicenseSGIB2_0
	// Simple Public License 2.0.
	SPDXLicenseSimPL2_0
	// Sun Industry Standards Source License v1.2.
	SPDXLicenseSISSL1_2
	// Sun Industry Standards Source License v1.1.
	SPDXLicenseSISSL
	// Sleepycat License.
	SPDXLicenseSleepycat
	// Standard ML of New Jersey License.
	SPDXLicenseSMLNJ
	// Secure Messaging Protocol Public License.
	SPDXLicenseSMPPL
	// SNIA Public License 1.1.
	SPDXLicenseSNIA
	// Spencer License 86.
	SPDXLicenseSpencer86
	// Spencer License 94.
	SPDXLicenseSpencer94
	// Spencer License 99.
	SPDXLicenseSpencer99
	// Sun Public License v1.0.
	SPDXLicenseSPL1_0
	// SugarCRM Public License v1.1.3.
	SPDXLicenseSugarCRM1_1_3
	// Scheme Widget Library (SWL) Software License Agreement.
	SPDXLicenseSWL
	// TCL/TK License.
	SPDXLicenseTCL
	// TCP Wrappers License.
	SPDXLicenseTCPWrappers
	// TMate Open Source License.
	SPDXLicenseTMate
	// TORQUE json.5+ Software License v1.1.
	SPDXLicenseTORQUE1_1
	// Trusster Open Source License.
	SPDXLicenseTOSL
	// Unicode License Agreement - Data Files and Software (2015).
	SPDXLicenseUnicodeDFS2015
	// Unicode License Agreement - Data Files and Software (2016).
	SPDXLicenseUnicodeDFS2016
	// Unicode Terms of Use.
	SPDXLicenseUnicodeTOU
	// The Unlicense.
	SPDXLicenseUnlicense
	// Universal Permissive License v1.0.
	SPDXLicenseUPL1_0
	// Vim License.
	SPDXLicenseVim
	// VOSTROM Public License for Open Source.
	SPDXLicenseVOSTROM
	// Vovida Software License v1.0.
	SPDXLicenseVSL1_0
	// W3C Software Notice and License (1998-07-20).
	SPDXLicenseW3C19980720
	// W3C Software Notice and Document License (2015-05-13).
	SPDXLicenseW3C20150513
	// W3C Software Notice and License (2002-12-31).
	SPDXLicenseW3C
	// Sybase Open Watcom Public License 1.0.
	SPDXLicenseWatcom1_0
	// Wsuipa License.
	SPDXLicenseWsuipa
	// Do What The F*ck You Want To Public License.
	SPDXLicenseWTFPL
	// X11 License.
	SPDXLicenseX11
	// Xerox License.
	SPDXLicenseXerox
	// XFree86 License 1.1.
	SPDXLicenseXFree861_1
	// xinetd License.
	SPDXLicenseXinetd
	// X.Net License.
	SPDXLicenseXnet
	// XPP License.
	SPDXLicenseXpp
	// XSkat License.
	SPDXLicenseXSkat
	// Yahoo! Public License v1.0.
	SPDXLicenseYPL1_0
	// Yahoo! Public License v1.1.
	SPDXLicenseYPL1_1
	// Zed License.
	SPDXLicenseZed
	// Zend License json.0.
	SPDXLicenseZend2_0
	// Zimbra Public License v1.3.
	SPDXLicenseZimbra1_3
	// Zimbra Public License v1.4.
	SPDXLicenseZimbra1_4
	// zlib/libpng License with Acknowledgement.
	SPDXLicenseZlibAcknowledgement
	// zlib License.
	SPDXLicenseZlib
	// Zope Public License 1.1.
	SPDXLicenseZPL1_1
	// Zope Public License 2.0.
	SPDXLicenseZPL2_0
	// Zope Public License 2.1.
	SPDXLicenseZPL2_1
)

func (code SPDXLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *SPDXLicense) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-open-source":
		*code = SPDXLicenseNotOpenSource
	case "0BSD":
		*code = SPDXLicense0BSD
	case "AAL":
		*code = SPDXLicenseAAL
	case "Abstyles":
		*code = SPDXLicenseAbstyles
	case "Adobe-2006":
		*code = SPDXLicenseAdobe2006
	case "Adobe-Glyph":
		*code = SPDXLicenseAdobeGlyph
	case "ADSL":
		*code = SPDXLicenseADSL
	case "AFL-1.1":
		*code = SPDXLicenseAFL1_1
	case "AFL-1.2":
		*code = SPDXLicenseAFL1_2
	case "AFL-2.0":
		*code = SPDXLicenseAFL2_0
	case "AFL-2.1":
		*code = SPDXLicenseAFL2_1
	case "AFL-3.0":
		*code = SPDXLicenseAFL3_0
	case "Afmparse":
		*code = SPDXLicenseAfmparse
	case "AGPL-1.0-only":
		*code = SPDXLicenseAGPL1_0Only
	case "AGPL-1.0-or-later":
		*code = SPDXLicenseAGPL1_0OrLater
	case "AGPL-3.0-only":
		*code = SPDXLicenseAGPL3_0Only
	case "AGPL-3.0-or-later":
		*code = SPDXLicenseAGPL3_0OrLater
	case "Aladdin":
		*code = SPDXLicenseAladdin
	case "AMDPLPA":
		*code = SPDXLicenseAMDPLPA
	case "AML":
		*code = SPDXLicenseAML
	case "AMPAS":
		*code = SPDXLicenseAMPAS
	case "ANTLR-PD":
		*code = SPDXLicenseANTLRPD
	case "Apache-1.0":
		*code = SPDXLicenseApache1_0
	case "Apache-1.1":
		*code = SPDXLicenseApache1_1
	case "Apache-2.0":
		*code = SPDXLicenseApache2_0
	case "APAFML":
		*code = SPDXLicenseAPAFML
	case "APL-1.0":
		*code = SPDXLicenseAPL1_0
	case "APSL-1.0":
		*code = SPDXLicenseAPSL1_0
	case "APSL-1.1":
		*code = SPDXLicenseAPSL1_1
	case "APSL-1.2":
		*code = SPDXLicenseAPSL1_2
	case "APSL-2.0":
		*code = SPDXLicenseAPSL2_0
	case "Artistic-1.0-cl8":
		*code = SPDXLicenseArtistic1_0Cl8
	case "Artistic-1.0-Perl":
		*code = SPDXLicenseArtistic1_0Perl
	case "Artistic-1.0":
		*code = SPDXLicenseArtistic1_0
	case "Artistic-2.0":
		*code = SPDXLicenseArtistic2_0
	case "Bahyph":
		*code = SPDXLicenseBahyph
	case "Barr":
		*code = SPDXLicenseBarr
	case "Beerware":
		*code = SPDXLicenseBeerware
	case "BitTorrent-1.0":
		*code = SPDXLicenseBitTorrent1_0
	case "BitTorrent-1.1":
		*code = SPDXLicenseBitTorrent1_1
	case "Borceux":
		*code = SPDXLicenseBorceux
	case "BSD-1-Clause":
		*code = SPDXLicenseBSD1Clause
	case "BSD-2-Clause-FreeBSD":
		*code = SPDXLicenseBSD2ClauseFreeBSD
	case "BSD-2-Clause-NetBSD":
		*code = SPDXLicenseBSD2ClauseNetBSD
	case "BSD-2-Clause-Patent":
		*code = SPDXLicenseBSD2ClausePatent
	case "BSD-2-Clause":
		*code = SPDXLicenseBSD2Clause
	case "BSD-3-Clause-Attribution":
		*code = SPDXLicenseBSD3ClauseAttribution
	case "BSD-3-Clause-Clear":
		*code = SPDXLicenseBSD3ClauseClear
	case "BSD-3-Clause-LBNL":
		*code = SPDXLicenseBSD3ClauseLBNL
	case "BSD-3-Clause-No-Nuclear-License-2014":
		*code = SPDXLicenseBSD3ClauseNoNuclearLicense2014
	case "BSD-3-Clause-No-Nuclear-License":
		*code = SPDXLicenseBSD3ClauseNoNuclearLicense
	case "BSD-3-Clause-No-Nuclear-Warranty":
		*code = SPDXLicenseBSD3ClauseNoNuclearWarranty
	case "BSD-3-Clause":
		*code = SPDXLicenseBSD3Clause
	case "BSD-4-Clause-UC":
		*code = SPDXLicenseBSD4ClauseUC
	case "BSD-4-Clause":
		*code = SPDXLicenseBSD4Clause
	case "BSD-Protection":
		*code = SPDXLicenseBSDProtection
	case "BSD-Source-Code":
		*code = SPDXLicenseBSDSourceCode
	case "BSL-1.0":
		*code = SPDXLicenseBSL1_0
	case "bzip2-1.0.5":
		*code = SPDXLicenseBzip21_0_5
	case "bzip2-1.0.6":
		*code = SPDXLicenseBzip21_0_6
	case "Caldera":
		*code = SPDXLicenseCaldera
	case "CATOSL-1.1":
		*code = SPDXLicenseCATOSL1_1
	case "CC-BY-1.0":
		*code = SPDXLicenseCCBY1_0
	case "CC-BY-2.0":
		*code = SPDXLicenseCCBY2_0
	case "CC-BY-2.5":
		*code = SPDXLicenseCCBY2_5
	case "CC-BY-3.0":
		*code = SPDXLicenseCCBY3_0
	case "CC-BY-4.0":
		*code = SPDXLicenseCCBY4_0
	case "CC-BY-NC-1.0":
		*code = SPDXLicenseCCBYNC1_0
	case "CC-BY-NC-2.0":
		*code = SPDXLicenseCCBYNC2_0
	case "CC-BY-NC-2.5":
		*code = SPDXLicenseCCBYNC2_5
	case "CC-BY-NC-3.0":
		*code = SPDXLicenseCCBYNC3_0
	case "CC-BY-NC-4.0":
		*code = SPDXLicenseCCBYNC4_0
	case "CC-BY-NC-ND-1.0":
		*code = SPDXLicenseCCBYNCND1_0
	case "CC-BY-NC-ND-2.0":
		*code = SPDXLicenseCCBYNCND2_0
	case "CC-BY-NC-ND-2.5":
		*code = SPDXLicenseCCBYNCND2_5
	case "CC-BY-NC-ND-3.0":
		*code = SPDXLicenseCCBYNCND3_0
	case "CC-BY-NC-ND-4.0":
		*code = SPDXLicenseCCBYNCND4_0
	case "CC-BY-NC-SA-1.0":
		*code = SPDXLicenseCCBYNCSA1_0
	case "CC-BY-NC-SA-2.0":
		*code = SPDXLicenseCCBYNCSA2_0
	case "CC-BY-NC-SA-2.5":
		*code = SPDXLicenseCCBYNCSA2_5
	case "CC-BY-NC-SA-3.0":
		*code = SPDXLicenseCCBYNCSA3_0
	case "CC-BY-NC-SA-4.0":
		*code = SPDXLicenseCCBYNCSA4_0
	case "CC-BY-ND-1.0":
		*code = SPDXLicenseCCBYND1_0
	case "CC-BY-ND-2.0":
		*code = SPDXLicenseCCBYND2_0
	case "CC-BY-ND-2.5":
		*code = SPDXLicenseCCBYND2_5
	case "CC-BY-ND-3.0":
		*code = SPDXLicenseCCBYND3_0
	case "CC-BY-ND-4.0":
		*code = SPDXLicenseCCBYND4_0
	case "CC-BY-SA-1.0":
		*code = SPDXLicenseCCBYSA1_0
	case "CC-BY-SA-2.0":
		*code = SPDXLicenseCCBYSA2_0
	case "CC-BY-SA-2.5":
		*code = SPDXLicenseCCBYSA2_5
	case "CC-BY-SA-3.0":
		*code = SPDXLicenseCCBYSA3_0
	case "CC-BY-SA-4.0":
		*code = SPDXLicenseCCBYSA4_0
	case "CC0-1.0":
		*code = SPDXLicenseCC01_0
	case "CDDL-1.0":
		*code = SPDXLicenseCDDL1_0
	case "CDDL-1.1":
		*code = SPDXLicenseCDDL1_1
	case "CDLA-Permissive-1.0":
		*code = SPDXLicenseCDLAPermissive1_0
	case "CDLA-Sharing-1.0":
		*code = SPDXLicenseCDLASharing1_0
	case "CECILL-1.0":
		*code = SPDXLicenseCECILL1_0
	case "CECILL-1.1":
		*code = SPDXLicenseCECILL1_1
	case "CECILL-2.0":
		*code = SPDXLicenseCECILL2_0
	case "CECILL-2.1":
		*code = SPDXLicenseCECILL2_1
	case "CECILL-B":
		*code = SPDXLicenseCECILLB
	case "CECILL-C":
		*code = SPDXLicenseCECILLC
	case "ClArtistic":
		*code = SPDXLicenseClArtistic
	case "CNRI-Jython":
		*code = SPDXLicenseCNRIJython
	case "CNRI-Python-GPL-Compatible":
		*code = SPDXLicenseCNRIPythonGPLCompatible
	case "CNRI-Python":
		*code = SPDXLicenseCNRIPython
	case "Condor-1.1":
		*code = SPDXLicenseCondor1_1
	case "CPAL-1.0":
		*code = SPDXLicenseCPAL1_0
	case "CPL-1.0":
		*code = SPDXLicenseCPL1_0
	case "CPOL-1.02":
		*code = SPDXLicenseCPOL1_02
	case "Crossword":
		*code = SPDXLicenseCrossword
	case "CrystalStacker":
		*code = SPDXLicenseCrystalStacker
	case "CUA-OPL-1.0":
		*code = SPDXLicenseCUAOPL1_0
	case "Cube":
		*code = SPDXLicenseCube
	case "curl":
		*code = SPDXLicenseCurl
	case "D-FSL-1.0":
		*code = SPDXLicenseDFSL1_0
	case "diffmark":
		*code = SPDXLicenseDiffmark
	case "DOC":
		*code = SPDXLicenseDOC
	case "Dotseqn":
		*code = SPDXLicenseDotseqn
	case "DSDP":
		*code = SPDXLicenseDSDP
	case "dvipdfm":
		*code = SPDXLicenseDvipdfm
	case "ECL-1.0":
		*code = SPDXLicenseECL1_0
	case "ECL-2.0":
		*code = SPDXLicenseECL2_0
	case "EFL-1.0":
		*code = SPDXLicenseEFL1_0
	case "EFL-2.0":
		*code = SPDXLicenseEFL2_0
	case "eGenix":
		*code = SPDXLicenseEGenix
	case "Entessa":
		*code = SPDXLicenseEntessa
	case "EPL-1.0":
		*code = SPDXLicenseEPL1_0
	case "EPL-2.0":
		*code = SPDXLicenseEPL2_0
	case "ErlPL-1.1":
		*code = SPDXLicenseErlPL1_1
	case "EUDatagrid":
		*code = SPDXLicenseEUDatagrid
	case "EUPL-1.0":
		*code = SPDXLicenseEUPL1_0
	case "EUPL-1.1":
		*code = SPDXLicenseEUPL1_1
	case "EUPL-1.2":
		*code = SPDXLicenseEUPL1_2
	case "Eurosym":
		*code = SPDXLicenseEurosym
	case "Fair":
		*code = SPDXLicenseFair
	case "Frameworx-1.0":
		*code = SPDXLicenseFrameworx1_0
	case "FreeImage":
		*code = SPDXLicenseFreeImage
	case "FSFAP":
		*code = SPDXLicenseFSFAP
	case "FSFUL":
		*code = SPDXLicenseFSFUL
	case "FSFULLR":
		*code = SPDXLicenseFSFULLR
	case "FTL":
		*code = SPDXLicenseFTL
	case "GFDL-1.1-only":
		*code = SPDXLicenseGFDL1_1Only
	case "GFDL-1.1-or-later":
		*code = SPDXLicenseGFDL1_1OrLater
	case "GFDL-1.2-only":
		*code = SPDXLicenseGFDL1_2Only
	case "GFDL-1.2-or-later":
		*code = SPDXLicenseGFDL1_2OrLater
	case "GFDL-1.3-only":
		*code = SPDXLicenseGFDL1_3Only
	case "GFDL-1.3-or-later":
		*code = SPDXLicenseGFDL1_3OrLater
	case "Giftware":
		*code = SPDXLicenseGiftware
	case "GL2PS":
		*code = SPDXLicenseGL2PS
	case "Glide":
		*code = SPDXLicenseGlide
	case "Glulxe":
		*code = SPDXLicenseGlulxe
	case "gnuplot":
		*code = SPDXLicenseGnuplot
	case "GPL-1.0-only":
		*code = SPDXLicenseGPL1_0Only
	case "GPL-1.0-or-later":
		*code = SPDXLicenseGPL1_0OrLater
	case "GPL-2.0-only":
		*code = SPDXLicenseGPL2_0Only
	case "GPL-2.0-or-later":
		*code = SPDXLicenseGPL2_0OrLater
	case "GPL-3.0-only":
		*code = SPDXLicenseGPL3_0Only
	case "GPL-3.0-or-later":
		*code = SPDXLicenseGPL3_0OrLater
	case "gSOAP-1.3b":
		*code = SPDXLicenseGSOAP1_3b
	case "HaskellReport":
		*code = SPDXLicenseHaskellReport
	case "HPND":
		*code = SPDXLicenseHPND
	case "IBM-pibs":
		*code = SPDXLicenseIBMPibs
	case "ICU":
		*code = SPDXLicenseICU
	case "IJG":
		*code = SPDXLicenseIJG
	case "ImageMagick":
		*code = SPDXLicenseImageMagick
	case "iMatix":
		*code = SPDXLicenseIMatix
	case "Imlib2":
		*code = SPDXLicenseImlib2
	case "Info-ZIP":
		*code = SPDXLicenseInfoZIP
	case "Intel-ACPI":
		*code = SPDXLicenseIntelACPI
	case "Intel":
		*code = SPDXLicenseIntel
	case "Interbase-1.0":
		*code = SPDXLicenseInterbase1_0
	case "IPA":
		*code = SPDXLicenseIPA
	case "IPL-1.0":
		*code = SPDXLicenseIPL1_0
	case "ISC":
		*code = SPDXLicenseISC
	case "JasPer-2.0":
		*code = SPDXLicenseJasPer2_0
	case "JSON":
		*code = SPDXLicenseJSON
	case "LAL-1.2":
		*code = SPDXLicenseLAL1_2
	case "LAL-1.3":
		*code = SPDXLicenseLAL1_3
	case "Latex2e":
		*code = SPDXLicenseLatex2e
	case "Leptonica":
		*code = SPDXLicenseLeptonica
	case "LGPL-2.0-only":
		*code = SPDXLicenseLGPL2_0Only
	case "LGPL-2.0-or-later":
		*code = SPDXLicenseLGPL2_0OrLater
	case "LGPL-2.1-only":
		*code = SPDXLicenseLGPL2_1Only
	case "LGPL-2.1-or-later":
		*code = SPDXLicenseLGPL2_1OrLater
	case "LGPL-3.0-only":
		*code = SPDXLicenseLGPL3_0Only
	case "LGPL-3.0-or-later":
		*code = SPDXLicenseLGPL3_0OrLater
	case "LGPLLR":
		*code = SPDXLicenseLGPLLR
	case "Libpng":
		*code = SPDXLicenseLibpng
	case "libtiff":
		*code = SPDXLicenseLibtiff
	case "LiLiQ-P-1.1":
		*code = SPDXLicenseLiLiQP1_1
	case "LiLiQ-R-1.1":
		*code = SPDXLicenseLiLiQR1_1
	case "LiLiQ-Rplus-1.1":
		*code = SPDXLicenseLiLiQRplus1_1
	case "Linux-OpenIB":
		*code = SPDXLicenseLinuxOpenIB
	case "LPL-1.0":
		*code = SPDXLicenseLPL1_0
	case "LPL-1.02":
		*code = SPDXLicenseLPL1_02
	case "LPPL-1.0":
		*code = SPDXLicenseLPPL1_0
	case "LPPL-1.1":
		*code = SPDXLicenseLPPL1_1
	case "LPPL-1.2":
		*code = SPDXLicenseLPPL1_2
	case "LPPL-1.3a":
		*code = SPDXLicenseLPPL1_3a
	case "LPPL-1.3c":
		*code = SPDXLicenseLPPL1_3c
	case "MakeIndex":
		*code = SPDXLicenseMakeIndex
	case "MirOS":
		*code = SPDXLicenseMirOS
	case "MIT-0":
		*code = SPDXLicenseMIT0
	case "MIT-advertising":
		*code = SPDXLicenseMITAdvertising
	case "MIT-CMU":
		*code = SPDXLicenseMITCMU
	case "MIT-enna":
		*code = SPDXLicenseMITEnna
	case "MIT-feh":
		*code = SPDXLicenseMITFeh
	case "MIT":
		*code = SPDXLicenseMIT
	case "MITNFA":
		*code = SPDXLicenseMITNFA
	case "Motosoto":
		*code = SPDXLicenseMotosoto
	case "mpich2":
		*code = SPDXLicenseMpich2
	case "MPL-1.0":
		*code = SPDXLicenseMPL1_0
	case "MPL-1.1":
		*code = SPDXLicenseMPL1_1
	case "MPL-2.0-no-copyleft-exception":
		*code = SPDXLicenseMPL2_0NoCopyleftException
	case "MPL-2.0":
		*code = SPDXLicenseMPL2_0
	case "MS-PL":
		*code = SPDXLicenseMSPL
	case "MS-RL":
		*code = SPDXLicenseMSRL
	case "MTLL":
		*code = SPDXLicenseMTLL
	case "Multics":
		*code = SPDXLicenseMultics
	case "Mup":
		*code = SPDXLicenseMup
	case "NASA-1.3":
		*code = SPDXLicenseNASA1_3
	case "Naumen":
		*code = SPDXLicenseNaumen
	case "NBPL-1.0":
		*code = SPDXLicenseNBPL1_0
	case "NCSA":
		*code = SPDXLicenseNCSA
	case "Net-SNMP":
		*code = SPDXLicenseNetSNMP
	case "NetCDF":
		*code = SPDXLicenseNetCDF
	case "Newsletr":
		*code = SPDXLicenseNewsletr
	case "NGPL":
		*code = SPDXLicenseNGPL
	case "NLOD-1.0":
		*code = SPDXLicenseNLOD1_0
	case "NLPL":
		*code = SPDXLicenseNLPL
	case "Nokia":
		*code = SPDXLicenseNokia
	case "NOSL":
		*code = SPDXLicenseNOSL
	case "Noweb":
		*code = SPDXLicenseNoweb
	case "NPL-1.0":
		*code = SPDXLicenseNPL1_0
	case "NPL-1.1":
		*code = SPDXLicenseNPL1_1
	case "NPOSL-3.0":
		*code = SPDXLicenseNPOSL3_0
	case "NRL":
		*code = SPDXLicenseNRL
	case "NTP":
		*code = SPDXLicenseNTP
	case "OCCT-PL":
		*code = SPDXLicenseOCCTPL
	case "OCLC-2.0":
		*code = SPDXLicenseOCLC2_0
	case "ODbL-1.0":
		*code = SPDXLicenseODbL1_0
	case "OFL-1.0":
		*code = SPDXLicenseOFL1_0
	case "OFL-1.1":
		*code = SPDXLicenseOFL1_1
	case "OGTSL":
		*code = SPDXLicenseOGTSL
	case "OLDAP-1.1":
		*code = SPDXLicenseOLDAP1_1
	case "OLDAP-1.2":
		*code = SPDXLicenseOLDAP1_2
	case "OLDAP-1.3":
		*code = SPDXLicenseOLDAP1_3
	case "OLDAP-1.4":
		*code = SPDXLicenseOLDAP1_4
	case "OLDAP-2.0.1":
		*code = SPDXLicenseOLDAP2_0_1
	case "OLDAP-2.0":
		*code = SPDXLicenseOLDAP2_0
	case "OLDAP-2.1":
		*code = SPDXLicenseOLDAP2_1
	case "OLDAP-2.2.1":
		*code = SPDXLicenseOLDAP2_2_1
	case "OLDAP-2.2.2":
		*code = SPDXLicenseOLDAP2_2_2
	case "OLDAP-2.2":
		*code = SPDXLicenseOLDAP2_2
	case "OLDAP-2.3":
		*code = SPDXLicenseOLDAP2_3
	case "OLDAP-2.4":
		*code = SPDXLicenseOLDAP2_4
	case "OLDAP-2.5":
		*code = SPDXLicenseOLDAP2_5
	case "OLDAP-2.6":
		*code = SPDXLicenseOLDAP2_6
	case "OLDAP-2.7":
		*code = SPDXLicenseOLDAP2_7
	case "OLDAP-2.8":
		*code = SPDXLicenseOLDAP2_8
	case "OML":
		*code = SPDXLicenseOML
	case "OpenSSL":
		*code = SPDXLicenseOpenSSL
	case "OPL-1.0":
		*code = SPDXLicenseOPL1_0
	case "OSET-PL-2.1":
		*code = SPDXLicenseOSETPL2_1
	case "OSL-1.0":
		*code = SPDXLicenseOSL1_0
	case "OSL-1.1":
		*code = SPDXLicenseOSL1_1
	case "OSL-2.0":
		*code = SPDXLicenseOSL2_0
	case "OSL-2.1":
		*code = SPDXLicenseOSL2_1
	case "OSL-3.0":
		*code = SPDXLicenseOSL3_0
	case "PDDL-1.0":
		*code = SPDXLicensePDDL1_0
	case "PHP-3.0":
		*code = SPDXLicensePHP3_0
	case "PHP-3.01":
		*code = SPDXLicensePHP3_01
	case "Plexus":
		*code = SPDXLicensePlexus
	case "PostgreSQL":
		*code = SPDXLicensePostgreSQL
	case "psfrag":
		*code = SPDXLicensePsfrag
	case "psutils":
		*code = SPDXLicensePsutils
	case "Python-2.0":
		*code = SPDXLicensePython2_0
	case "Qhull":
		*code = SPDXLicenseQhull
	case "QPL-1.0":
		*code = SPDXLicenseQPL1_0
	case "Rdisc":
		*code = SPDXLicenseRdisc
	case "RHeCos-1.1":
		*code = SPDXLicenseRHeCos1_1
	case "RPL-1.1":
		*code = SPDXLicenseRPL1_1
	case "RPL-1.5":
		*code = SPDXLicenseRPL1_5
	case "RPSL-1.0":
		*code = SPDXLicenseRPSL1_0
	case "RSA-MD":
		*code = SPDXLicenseRSAMD
	case "RSCPL":
		*code = SPDXLicenseRSCPL
	case "Ruby":
		*code = SPDXLicenseRuby
	case "SAX-PD":
		*code = SPDXLicenseSAXPD
	case "Saxpath":
		*code = SPDXLicenseSaxpath
	case "SCEA":
		*code = SPDXLicenseSCEA
	case "Sendmail":
		*code = SPDXLicenseSendmail
	case "SGI-B-1.0":
		*code = SPDXLicenseSGIB1_0
	case "SGI-B-1.1":
		*code = SPDXLicenseSGIB1_1
	case "SGI-B-2.0":
		*code = SPDXLicenseSGIB2_0
	case "SimPL-2.0":
		*code = SPDXLicenseSimPL2_0
	case "SISSL-1.2":
		*code = SPDXLicenseSISSL1_2
	case "SISSL":
		*code = SPDXLicenseSISSL
	case "Sleepycat":
		*code = SPDXLicenseSleepycat
	case "SMLNJ":
		*code = SPDXLicenseSMLNJ
	case "SMPPL":
		*code = SPDXLicenseSMPPL
	case "SNIA":
		*code = SPDXLicenseSNIA
	case "Spencer-86":
		*code = SPDXLicenseSpencer86
	case "Spencer-94":
		*code = SPDXLicenseSpencer94
	case "Spencer-99":
		*code = SPDXLicenseSpencer99
	case "SPL-1.0":
		*code = SPDXLicenseSPL1_0
	case "SugarCRM-1.1.3":
		*code = SPDXLicenseSugarCRM1_1_3
	case "SWL":
		*code = SPDXLicenseSWL
	case "TCL":
		*code = SPDXLicenseTCL
	case "TCP-wrappers":
		*code = SPDXLicenseTCPWrappers
	case "TMate":
		*code = SPDXLicenseTMate
	case "TORQUE-1.1":
		*code = SPDXLicenseTORQUE1_1
	case "TOSL":
		*code = SPDXLicenseTOSL
	case "Unicode-DFS-2015":
		*code = SPDXLicenseUnicodeDFS2015
	case "Unicode-DFS-2016":
		*code = SPDXLicenseUnicodeDFS2016
	case "Unicode-TOU":
		*code = SPDXLicenseUnicodeTOU
	case "Unlicense":
		*code = SPDXLicenseUnlicense
	case "UPL-1.0":
		*code = SPDXLicenseUPL1_0
	case "Vim":
		*code = SPDXLicenseVim
	case "VOSTROM":
		*code = SPDXLicenseVOSTROM
	case "VSL-1.0":
		*code = SPDXLicenseVSL1_0
	case "W3C-19980720":
		*code = SPDXLicenseW3C19980720
	case "W3C-20150513":
		*code = SPDXLicenseW3C20150513
	case "W3C":
		*code = SPDXLicenseW3C
	case "Watcom-1.0":
		*code = SPDXLicenseWatcom1_0
	case "Wsuipa":
		*code = SPDXLicenseWsuipa
	case "WTFPL":
		*code = SPDXLicenseWTFPL
	case "X11":
		*code = SPDXLicenseX11
	case "Xerox":
		*code = SPDXLicenseXerox
	case "XFree86-1.1":
		*code = SPDXLicenseXFree861_1
	case "xinetd":
		*code = SPDXLicenseXinetd
	case "Xnet":
		*code = SPDXLicenseXnet
	case "xpp":
		*code = SPDXLicenseXpp
	case "XSkat":
		*code = SPDXLicenseXSkat
	case "YPL-1.0":
		*code = SPDXLicenseYPL1_0
	case "YPL-1.1":
		*code = SPDXLicenseYPL1_1
	case "Zed":
		*code = SPDXLicenseZed
	case "Zend-2.0":
		*code = SPDXLicenseZend2_0
	case "Zimbra-1.3":
		*code = SPDXLicenseZimbra1_3
	case "Zimbra-1.4":
		*code = SPDXLicenseZimbra1_4
	case "zlib-acknowledgement":
		*code = SPDXLicenseZlibAcknowledgement
	case "Zlib":
		*code = SPDXLicenseZlib
	case "ZPL-1.1":
		*code = SPDXLicenseZPL1_1
	case "ZPL-2.0":
		*code = SPDXLicenseZPL2_0
	case "ZPL-2.1":
		*code = SPDXLicenseZPL2_1
	default:
		return fmt.Errorf("unknown SPDXLicense code `%s`", s)
	}
	return nil
}

func (code SPDXLicense) String() string {
	return code.Code()
}

func (code SPDXLicense) Code() string {
	switch code {
	case SPDXLicenseNotOpenSource:
		return "not-open-source"
	case SPDXLicense0BSD:
		return "0BSD"
	case SPDXLicenseAAL:
		return "AAL"
	case SPDXLicenseAbstyles:
		return "Abstyles"
	case SPDXLicenseAdobe2006:
		return "Adobe-2006"
	case SPDXLicenseAdobeGlyph:
		return "Adobe-Glyph"
	case SPDXLicenseADSL:
		return "ADSL"
	case SPDXLicenseAFL1_1:
		return "AFL-1.1"
	case SPDXLicenseAFL1_2:
		return "AFL-1.2"
	case SPDXLicenseAFL2_0:
		return "AFL-2.0"
	case SPDXLicenseAFL2_1:
		return "AFL-2.1"
	case SPDXLicenseAFL3_0:
		return "AFL-3.0"
	case SPDXLicenseAfmparse:
		return "Afmparse"
	case SPDXLicenseAGPL1_0Only:
		return "AGPL-1.0-only"
	case SPDXLicenseAGPL1_0OrLater:
		return "AGPL-1.0-or-later"
	case SPDXLicenseAGPL3_0Only:
		return "AGPL-3.0-only"
	case SPDXLicenseAGPL3_0OrLater:
		return "AGPL-3.0-or-later"
	case SPDXLicenseAladdin:
		return "Aladdin"
	case SPDXLicenseAMDPLPA:
		return "AMDPLPA"
	case SPDXLicenseAML:
		return "AML"
	case SPDXLicenseAMPAS:
		return "AMPAS"
	case SPDXLicenseANTLRPD:
		return "ANTLR-PD"
	case SPDXLicenseApache1_0:
		return "Apache-1.0"
	case SPDXLicenseApache1_1:
		return "Apache-1.1"
	case SPDXLicenseApache2_0:
		return "Apache-2.0"
	case SPDXLicenseAPAFML:
		return "APAFML"
	case SPDXLicenseAPL1_0:
		return "APL-1.0"
	case SPDXLicenseAPSL1_0:
		return "APSL-1.0"
	case SPDXLicenseAPSL1_1:
		return "APSL-1.1"
	case SPDXLicenseAPSL1_2:
		return "APSL-1.2"
	case SPDXLicenseAPSL2_0:
		return "APSL-2.0"
	case SPDXLicenseArtistic1_0Cl8:
		return "Artistic-1.0-cl8"
	case SPDXLicenseArtistic1_0Perl:
		return "Artistic-1.0-Perl"
	case SPDXLicenseArtistic1_0:
		return "Artistic-1.0"
	case SPDXLicenseArtistic2_0:
		return "Artistic-2.0"
	case SPDXLicenseBahyph:
		return "Bahyph"
	case SPDXLicenseBarr:
		return "Barr"
	case SPDXLicenseBeerware:
		return "Beerware"
	case SPDXLicenseBitTorrent1_0:
		return "BitTorrent-1.0"
	case SPDXLicenseBitTorrent1_1:
		return "BitTorrent-1.1"
	case SPDXLicenseBorceux:
		return "Borceux"
	case SPDXLicenseBSD1Clause:
		return "BSD-1-Clause"
	case SPDXLicenseBSD2ClauseFreeBSD:
		return "BSD-2-Clause-FreeBSD"
	case SPDXLicenseBSD2ClauseNetBSD:
		return "BSD-2-Clause-NetBSD"
	case SPDXLicenseBSD2ClausePatent:
		return "BSD-2-Clause-Patent"
	case SPDXLicenseBSD2Clause:
		return "BSD-2-Clause"
	case SPDXLicenseBSD3ClauseAttribution:
		return "BSD-3-Clause-Attribution"
	case SPDXLicenseBSD3ClauseClear:
		return "BSD-3-Clause-Clear"
	case SPDXLicenseBSD3ClauseLBNL:
		return "BSD-3-Clause-LBNL"
	case SPDXLicenseBSD3ClauseNoNuclearLicense2014:
		return "BSD-3-Clause-No-Nuclear-License-2014"
	case SPDXLicenseBSD3ClauseNoNuclearLicense:
		return "BSD-3-Clause-No-Nuclear-License"
	case SPDXLicenseBSD3ClauseNoNuclearWarranty:
		return "BSD-3-Clause-No-Nuclear-Warranty"
	case SPDXLicenseBSD3Clause:
		return "BSD-3-Clause"
	case SPDXLicenseBSD4ClauseUC:
		return "BSD-4-Clause-UC"
	case SPDXLicenseBSD4Clause:
		return "BSD-4-Clause"
	case SPDXLicenseBSDProtection:
		return "BSD-Protection"
	case SPDXLicenseBSDSourceCode:
		return "BSD-Source-Code"
	case SPDXLicenseBSL1_0:
		return "BSL-1.0"
	case SPDXLicenseBzip21_0_5:
		return "bzip2-1.0.5"
	case SPDXLicenseBzip21_0_6:
		return "bzip2-1.0.6"
	case SPDXLicenseCaldera:
		return "Caldera"
	case SPDXLicenseCATOSL1_1:
		return "CATOSL-1.1"
	case SPDXLicenseCCBY1_0:
		return "CC-BY-1.0"
	case SPDXLicenseCCBY2_0:
		return "CC-BY-2.0"
	case SPDXLicenseCCBY2_5:
		return "CC-BY-2.5"
	case SPDXLicenseCCBY3_0:
		return "CC-BY-3.0"
	case SPDXLicenseCCBY4_0:
		return "CC-BY-4.0"
	case SPDXLicenseCCBYNC1_0:
		return "CC-BY-NC-1.0"
	case SPDXLicenseCCBYNC2_0:
		return "CC-BY-NC-2.0"
	case SPDXLicenseCCBYNC2_5:
		return "CC-BY-NC-2.5"
	case SPDXLicenseCCBYNC3_0:
		return "CC-BY-NC-3.0"
	case SPDXLicenseCCBYNC4_0:
		return "CC-BY-NC-4.0"
	case SPDXLicenseCCBYNCND1_0:
		return "CC-BY-NC-ND-1.0"
	case SPDXLicenseCCBYNCND2_0:
		return "CC-BY-NC-ND-2.0"
	case SPDXLicenseCCBYNCND2_5:
		return "CC-BY-NC-ND-2.5"
	case SPDXLicenseCCBYNCND3_0:
		return "CC-BY-NC-ND-3.0"
	case SPDXLicenseCCBYNCND4_0:
		return "CC-BY-NC-ND-4.0"
	case SPDXLicenseCCBYNCSA1_0:
		return "CC-BY-NC-SA-1.0"
	case SPDXLicenseCCBYNCSA2_0:
		return "CC-BY-NC-SA-2.0"
	case SPDXLicenseCCBYNCSA2_5:
		return "CC-BY-NC-SA-2.5"
	case SPDXLicenseCCBYNCSA3_0:
		return "CC-BY-NC-SA-3.0"
	case SPDXLicenseCCBYNCSA4_0:
		return "CC-BY-NC-SA-4.0"
	case SPDXLicenseCCBYND1_0:
		return "CC-BY-ND-1.0"
	case SPDXLicenseCCBYND2_0:
		return "CC-BY-ND-2.0"
	case SPDXLicenseCCBYND2_5:
		return "CC-BY-ND-2.5"
	case SPDXLicenseCCBYND3_0:
		return "CC-BY-ND-3.0"
	case SPDXLicenseCCBYND4_0:
		return "CC-BY-ND-4.0"
	case SPDXLicenseCCBYSA1_0:
		return "CC-BY-SA-1.0"
	case SPDXLicenseCCBYSA2_0:
		return "CC-BY-SA-2.0"
	case SPDXLicenseCCBYSA2_5:
		return "CC-BY-SA-2.5"
	case SPDXLicenseCCBYSA3_0:
		return "CC-BY-SA-3.0"
	case SPDXLicenseCCBYSA4_0:
		return "CC-BY-SA-4.0"
	case SPDXLicenseCC01_0:
		return "CC0-1.0"
	case SPDXLicenseCDDL1_0:
		return "CDDL-1.0"
	case SPDXLicenseCDDL1_1:
		return "CDDL-1.1"
	case SPDXLicenseCDLAPermissive1_0:
		return "CDLA-Permissive-1.0"
	case SPDXLicenseCDLASharing1_0:
		return "CDLA-Sharing-1.0"
	case SPDXLicenseCECILL1_0:
		return "CECILL-1.0"
	case SPDXLicenseCECILL1_1:
		return "CECILL-1.1"
	case SPDXLicenseCECILL2_0:
		return "CECILL-2.0"
	case SPDXLicenseCECILL2_1:
		return "CECILL-2.1"
	case SPDXLicenseCECILLB:
		return "CECILL-B"
	case SPDXLicenseCECILLC:
		return "CECILL-C"
	case SPDXLicenseClArtistic:
		return "ClArtistic"
	case SPDXLicenseCNRIJython:
		return "CNRI-Jython"
	case SPDXLicenseCNRIPythonGPLCompatible:
		return "CNRI-Python-GPL-Compatible"
	case SPDXLicenseCNRIPython:
		return "CNRI-Python"
	case SPDXLicenseCondor1_1:
		return "Condor-1.1"
	case SPDXLicenseCPAL1_0:
		return "CPAL-1.0"
	case SPDXLicenseCPL1_0:
		return "CPL-1.0"
	case SPDXLicenseCPOL1_02:
		return "CPOL-1.02"
	case SPDXLicenseCrossword:
		return "Crossword"
	case SPDXLicenseCrystalStacker:
		return "CrystalStacker"
	case SPDXLicenseCUAOPL1_0:
		return "CUA-OPL-1.0"
	case SPDXLicenseCube:
		return "Cube"
	case SPDXLicenseCurl:
		return "curl"
	case SPDXLicenseDFSL1_0:
		return "D-FSL-1.0"
	case SPDXLicenseDiffmark:
		return "diffmark"
	case SPDXLicenseDOC:
		return "DOC"
	case SPDXLicenseDotseqn:
		return "Dotseqn"
	case SPDXLicenseDSDP:
		return "DSDP"
	case SPDXLicenseDvipdfm:
		return "dvipdfm"
	case SPDXLicenseECL1_0:
		return "ECL-1.0"
	case SPDXLicenseECL2_0:
		return "ECL-2.0"
	case SPDXLicenseEFL1_0:
		return "EFL-1.0"
	case SPDXLicenseEFL2_0:
		return "EFL-2.0"
	case SPDXLicenseEGenix:
		return "eGenix"
	case SPDXLicenseEntessa:
		return "Entessa"
	case SPDXLicenseEPL1_0:
		return "EPL-1.0"
	case SPDXLicenseEPL2_0:
		return "EPL-2.0"
	case SPDXLicenseErlPL1_1:
		return "ErlPL-1.1"
	case SPDXLicenseEUDatagrid:
		return "EUDatagrid"
	case SPDXLicenseEUPL1_0:
		return "EUPL-1.0"
	case SPDXLicenseEUPL1_1:
		return "EUPL-1.1"
	case SPDXLicenseEUPL1_2:
		return "EUPL-1.2"
	case SPDXLicenseEurosym:
		return "Eurosym"
	case SPDXLicenseFair:
		return "Fair"
	case SPDXLicenseFrameworx1_0:
		return "Frameworx-1.0"
	case SPDXLicenseFreeImage:
		return "FreeImage"
	case SPDXLicenseFSFAP:
		return "FSFAP"
	case SPDXLicenseFSFUL:
		return "FSFUL"
	case SPDXLicenseFSFULLR:
		return "FSFULLR"
	case SPDXLicenseFTL:
		return "FTL"
	case SPDXLicenseGFDL1_1Only:
		return "GFDL-1.1-only"
	case SPDXLicenseGFDL1_1OrLater:
		return "GFDL-1.1-or-later"
	case SPDXLicenseGFDL1_2Only:
		return "GFDL-1.2-only"
	case SPDXLicenseGFDL1_2OrLater:
		return "GFDL-1.2-or-later"
	case SPDXLicenseGFDL1_3Only:
		return "GFDL-1.3-only"
	case SPDXLicenseGFDL1_3OrLater:
		return "GFDL-1.3-or-later"
	case SPDXLicenseGiftware:
		return "Giftware"
	case SPDXLicenseGL2PS:
		return "GL2PS"
	case SPDXLicenseGlide:
		return "Glide"
	case SPDXLicenseGlulxe:
		return "Glulxe"
	case SPDXLicenseGnuplot:
		return "gnuplot"
	case SPDXLicenseGPL1_0Only:
		return "GPL-1.0-only"
	case SPDXLicenseGPL1_0OrLater:
		return "GPL-1.0-or-later"
	case SPDXLicenseGPL2_0Only:
		return "GPL-2.0-only"
	case SPDXLicenseGPL2_0OrLater:
		return "GPL-2.0-or-later"
	case SPDXLicenseGPL3_0Only:
		return "GPL-3.0-only"
	case SPDXLicenseGPL3_0OrLater:
		return "GPL-3.0-or-later"
	case SPDXLicenseGSOAP1_3b:
		return "gSOAP-1.3b"
	case SPDXLicenseHaskellReport:
		return "HaskellReport"
	case SPDXLicenseHPND:
		return "HPND"
	case SPDXLicenseIBMPibs:
		return "IBM-pibs"
	case SPDXLicenseICU:
		return "ICU"
	case SPDXLicenseIJG:
		return "IJG"
	case SPDXLicenseImageMagick:
		return "ImageMagick"
	case SPDXLicenseIMatix:
		return "iMatix"
	case SPDXLicenseImlib2:
		return "Imlib2"
	case SPDXLicenseInfoZIP:
		return "Info-ZIP"
	case SPDXLicenseIntelACPI:
		return "Intel-ACPI"
	case SPDXLicenseIntel:
		return "Intel"
	case SPDXLicenseInterbase1_0:
		return "Interbase-1.0"
	case SPDXLicenseIPA:
		return "IPA"
	case SPDXLicenseIPL1_0:
		return "IPL-1.0"
	case SPDXLicenseISC:
		return "ISC"
	case SPDXLicenseJasPer2_0:
		return "JasPer-2.0"
	case SPDXLicenseJSON:
		return "JSON"
	case SPDXLicenseLAL1_2:
		return "LAL-1.2"
	case SPDXLicenseLAL1_3:
		return "LAL-1.3"
	case SPDXLicenseLatex2e:
		return "Latex2e"
	case SPDXLicenseLeptonica:
		return "Leptonica"
	case SPDXLicenseLGPL2_0Only:
		return "LGPL-2.0-only"
	case SPDXLicenseLGPL2_0OrLater:
		return "LGPL-2.0-or-later"
	case SPDXLicenseLGPL2_1Only:
		return "LGPL-2.1-only"
	case SPDXLicenseLGPL2_1OrLater:
		return "LGPL-2.1-or-later"
	case SPDXLicenseLGPL3_0Only:
		return "LGPL-3.0-only"
	case SPDXLicenseLGPL3_0OrLater:
		return "LGPL-3.0-or-later"
	case SPDXLicenseLGPLLR:
		return "LGPLLR"
	case SPDXLicenseLibpng:
		return "Libpng"
	case SPDXLicenseLibtiff:
		return "libtiff"
	case SPDXLicenseLiLiQP1_1:
		return "LiLiQ-P-1.1"
	case SPDXLicenseLiLiQR1_1:
		return "LiLiQ-R-1.1"
	case SPDXLicenseLiLiQRplus1_1:
		return "LiLiQ-Rplus-1.1"
	case SPDXLicenseLinuxOpenIB:
		return "Linux-OpenIB"
	case SPDXLicenseLPL1_0:
		return "LPL-1.0"
	case SPDXLicenseLPL1_02:
		return "LPL-1.02"
	case SPDXLicenseLPPL1_0:
		return "LPPL-1.0"
	case SPDXLicenseLPPL1_1:
		return "LPPL-1.1"
	case SPDXLicenseLPPL1_2:
		return "LPPL-1.2"
	case SPDXLicenseLPPL1_3a:
		return "LPPL-1.3a"
	case SPDXLicenseLPPL1_3c:
		return "LPPL-1.3c"
	case SPDXLicenseMakeIndex:
		return "MakeIndex"
	case SPDXLicenseMirOS:
		return "MirOS"
	case SPDXLicenseMIT0:
		return "MIT-0"
	case SPDXLicenseMITAdvertising:
		return "MIT-advertising"
	case SPDXLicenseMITCMU:
		return "MIT-CMU"
	case SPDXLicenseMITEnna:
		return "MIT-enna"
	case SPDXLicenseMITFeh:
		return "MIT-feh"
	case SPDXLicenseMIT:
		return "MIT"
	case SPDXLicenseMITNFA:
		return "MITNFA"
	case SPDXLicenseMotosoto:
		return "Motosoto"
	case SPDXLicenseMpich2:
		return "mpich2"
	case SPDXLicenseMPL1_0:
		return "MPL-1.0"
	case SPDXLicenseMPL1_1:
		return "MPL-1.1"
	case SPDXLicenseMPL2_0NoCopyleftException:
		return "MPL-2.0-no-copyleft-exception"
	case SPDXLicenseMPL2_0:
		return "MPL-2.0"
	case SPDXLicenseMSPL:
		return "MS-PL"
	case SPDXLicenseMSRL:
		return "MS-RL"
	case SPDXLicenseMTLL:
		return "MTLL"
	case SPDXLicenseMultics:
		return "Multics"
	case SPDXLicenseMup:
		return "Mup"
	case SPDXLicenseNASA1_3:
		return "NASA-1.3"
	case SPDXLicenseNaumen:
		return "Naumen"
	case SPDXLicenseNBPL1_0:
		return "NBPL-1.0"
	case SPDXLicenseNCSA:
		return "NCSA"
	case SPDXLicenseNetSNMP:
		return "Net-SNMP"
	case SPDXLicenseNetCDF:
		return "NetCDF"
	case SPDXLicenseNewsletr:
		return "Newsletr"
	case SPDXLicenseNGPL:
		return "NGPL"
	case SPDXLicenseNLOD1_0:
		return "NLOD-1.0"
	case SPDXLicenseNLPL:
		return "NLPL"
	case SPDXLicenseNokia:
		return "Nokia"
	case SPDXLicenseNOSL:
		return "NOSL"
	case SPDXLicenseNoweb:
		return "Noweb"
	case SPDXLicenseNPL1_0:
		return "NPL-1.0"
	case SPDXLicenseNPL1_1:
		return "NPL-1.1"
	case SPDXLicenseNPOSL3_0:
		return "NPOSL-3.0"
	case SPDXLicenseNRL:
		return "NRL"
	case SPDXLicenseNTP:
		return "NTP"
	case SPDXLicenseOCCTPL:
		return "OCCT-PL"
	case SPDXLicenseOCLC2_0:
		return "OCLC-2.0"
	case SPDXLicenseODbL1_0:
		return "ODbL-1.0"
	case SPDXLicenseOFL1_0:
		return "OFL-1.0"
	case SPDXLicenseOFL1_1:
		return "OFL-1.1"
	case SPDXLicenseOGTSL:
		return "OGTSL"
	case SPDXLicenseOLDAP1_1:
		return "OLDAP-1.1"
	case SPDXLicenseOLDAP1_2:
		return "OLDAP-1.2"
	case SPDXLicenseOLDAP1_3:
		return "OLDAP-1.3"
	case SPDXLicenseOLDAP1_4:
		return "OLDAP-1.4"
	case SPDXLicenseOLDAP2_0_1:
		return "OLDAP-2.0.1"
	case SPDXLicenseOLDAP2_0:
		return "OLDAP-2.0"
	case SPDXLicenseOLDAP2_1:
		return "OLDAP-2.1"
	case SPDXLicenseOLDAP2_2_1:
		return "OLDAP-2.2.1"
	case SPDXLicenseOLDAP2_2_2:
		return "OLDAP-2.2.2"
	case SPDXLicenseOLDAP2_2:
		return "OLDAP-2.2"
	case SPDXLicenseOLDAP2_3:
		return "OLDAP-2.3"
	case SPDXLicenseOLDAP2_4:
		return "OLDAP-2.4"
	case SPDXLicenseOLDAP2_5:
		return "OLDAP-2.5"
	case SPDXLicenseOLDAP2_6:
		return "OLDAP-2.6"
	case SPDXLicenseOLDAP2_7:
		return "OLDAP-2.7"
	case SPDXLicenseOLDAP2_8:
		return "OLDAP-2.8"
	case SPDXLicenseOML:
		return "OML"
	case SPDXLicenseOpenSSL:
		return "OpenSSL"
	case SPDXLicenseOPL1_0:
		return "OPL-1.0"
	case SPDXLicenseOSETPL2_1:
		return "OSET-PL-2.1"
	case SPDXLicenseOSL1_0:
		return "OSL-1.0"
	case SPDXLicenseOSL1_1:
		return "OSL-1.1"
	case SPDXLicenseOSL2_0:
		return "OSL-2.0"
	case SPDXLicenseOSL2_1:
		return "OSL-2.1"
	case SPDXLicenseOSL3_0:
		return "OSL-3.0"
	case SPDXLicensePDDL1_0:
		return "PDDL-1.0"
	case SPDXLicensePHP3_0:
		return "PHP-3.0"
	case SPDXLicensePHP3_01:
		return "PHP-3.01"
	case SPDXLicensePlexus:
		return "Plexus"
	case SPDXLicensePostgreSQL:
		return "PostgreSQL"
	case SPDXLicensePsfrag:
		return "psfrag"
	case SPDXLicensePsutils:
		return "psutils"
	case SPDXLicensePython2_0:
		return "Python-2.0"
	case SPDXLicenseQhull:
		return "Qhull"
	case SPDXLicenseQPL1_0:
		return "QPL-1.0"
	case SPDXLicenseRdisc:
		return "Rdisc"
	case SPDXLicenseRHeCos1_1:
		return "RHeCos-1.1"
	case SPDXLicenseRPL1_1:
		return "RPL-1.1"
	case SPDXLicenseRPL1_5:
		return "RPL-1.5"
	case SPDXLicenseRPSL1_0:
		return "RPSL-1.0"
	case SPDXLicenseRSAMD:
		return "RSA-MD"
	case SPDXLicenseRSCPL:
		return "RSCPL"
	case SPDXLicenseRuby:
		return "Ruby"
	case SPDXLicenseSAXPD:
		return "SAX-PD"
	case SPDXLicenseSaxpath:
		return "Saxpath"
	case SPDXLicenseSCEA:
		return "SCEA"
	case SPDXLicenseSendmail:
		return "Sendmail"
	case SPDXLicenseSGIB1_0:
		return "SGI-B-1.0"
	case SPDXLicenseSGIB1_1:
		return "SGI-B-1.1"
	case SPDXLicenseSGIB2_0:
		return "SGI-B-2.0"
	case SPDXLicenseSimPL2_0:
		return "SimPL-2.0"
	case SPDXLicenseSISSL1_2:
		return "SISSL-1.2"
	case SPDXLicenseSISSL:
		return "SISSL"
	case SPDXLicenseSleepycat:
		return "Sleepycat"
	case SPDXLicenseSMLNJ:
		return "SMLNJ"
	case SPDXLicenseSMPPL:
		return "SMPPL"
	case SPDXLicenseSNIA:
		return "SNIA"
	case SPDXLicenseSpencer86:
		return "Spencer-86"
	case SPDXLicenseSpencer94:
		return "Spencer-94"
	case SPDXLicenseSpencer99:
		return "Spencer-99"
	case SPDXLicenseSPL1_0:
		return "SPL-1.0"
	case SPDXLicenseSugarCRM1_1_3:
		return "SugarCRM-1.1.3"
	case SPDXLicenseSWL:
		return "SWL"
	case SPDXLicenseTCL:
		return "TCL"
	case SPDXLicenseTCPWrappers:
		return "TCP-wrappers"
	case SPDXLicenseTMate:
		return "TMate"
	case SPDXLicenseTORQUE1_1:
		return "TORQUE-1.1"
	case SPDXLicenseTOSL:
		return "TOSL"
	case SPDXLicenseUnicodeDFS2015:
		return "Unicode-DFS-2015"
	case SPDXLicenseUnicodeDFS2016:
		return "Unicode-DFS-2016"
	case SPDXLicenseUnicodeTOU:
		return "Unicode-TOU"
	case SPDXLicenseUnlicense:
		return "Unlicense"
	case SPDXLicenseUPL1_0:
		return "UPL-1.0"
	case SPDXLicenseVim:
		return "Vim"
	case SPDXLicenseVOSTROM:
		return "VOSTROM"
	case SPDXLicenseVSL1_0:
		return "VSL-1.0"
	case SPDXLicenseW3C19980720:
		return "W3C-19980720"
	case SPDXLicenseW3C20150513:
		return "W3C-20150513"
	case SPDXLicenseW3C:
		return "W3C"
	case SPDXLicenseWatcom1_0:
		return "Watcom-1.0"
	case SPDXLicenseWsuipa:
		return "Wsuipa"
	case SPDXLicenseWTFPL:
		return "WTFPL"
	case SPDXLicenseX11:
		return "X11"
	case SPDXLicenseXerox:
		return "Xerox"
	case SPDXLicenseXFree861_1:
		return "XFree86-1.1"
	case SPDXLicenseXinetd:
		return "xinetd"
	case SPDXLicenseXnet:
		return "Xnet"
	case SPDXLicenseXpp:
		return "xpp"
	case SPDXLicenseXSkat:
		return "XSkat"
	case SPDXLicenseYPL1_0:
		return "YPL-1.0"
	case SPDXLicenseYPL1_1:
		return "YPL-1.1"
	case SPDXLicenseZed:
		return "Zed"
	case SPDXLicenseZend2_0:
		return "Zend-2.0"
	case SPDXLicenseZimbra1_3:
		return "Zimbra-1.3"
	case SPDXLicenseZimbra1_4:
		return "Zimbra-1.4"
	case SPDXLicenseZlibAcknowledgement:
		return "zlib-acknowledgement"
	case SPDXLicenseZlib:
		return "Zlib"
	case SPDXLicenseZPL1_1:
		return "ZPL-1.1"
	case SPDXLicenseZPL2_0:
		return "ZPL-2.0"
	case SPDXLicenseZPL2_1:
		return "ZPL-2.1"
	}
	return "<unknown>"
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
		return "Academic Free License json.0"
	case SPDXLicenseAFL2_1:
		return "Academic Free License json.1"
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
		return "CeCILL Free Software License Agreement json.0"
	case SPDXLicenseCECILL2_1:
		return "CeCILL Free Software License Agreement json.1"
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
		return "Educational Community License json.0"
	case SPDXLicenseEFL1_0:
		return "Eiffel Forum License v1.0"
	case SPDXLicenseEFL2_0:
		return "Eiffel Forum License json.0"
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
		return "GNU General Public License json.0 only"
	case SPDXLicenseGPL2_0OrLater:
		return "GNU General Public License json.0 or later"
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
		return "GNU Library General Public License json.only"
	case SPDXLicenseLGPL2_0OrLater:
		return "GNU Library General Public License json.or later"
	case SPDXLicenseLGPL2_1Only:
		return "GNU Lesser General Public License json.1 only"
	case SPDXLicenseLGPL2_1OrLater:
		return "GNU Lesser General Public License json.1 or later"
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
		return "Open LDAP Public License json.0.1"
	case SPDXLicenseOLDAP2_0:
		return "Open LDAP Public License json.0 (or possibly 2.0A and 2.0B)"
	case SPDXLicenseOLDAP2_1:
		return "Open LDAP Public License json.1"
	case SPDXLicenseOLDAP2_2_1:
		return "Open LDAP Public License json.2.1"
	case SPDXLicenseOLDAP2_2_2:
		return "Open LDAP Public License 2.2.2"
	case SPDXLicenseOLDAP2_2:
		return "Open LDAP Public License json.2"
	case SPDXLicenseOLDAP2_3:
		return "Open LDAP Public License json.3"
	case SPDXLicenseOLDAP2_4:
		return "Open LDAP Public License json.4"
	case SPDXLicenseOLDAP2_5:
		return "Open LDAP Public License json.5"
	case SPDXLicenseOLDAP2_6:
		return "Open LDAP Public License json.6"
	case SPDXLicenseOLDAP2_7:
		return "Open LDAP Public License json.7"
	case SPDXLicenseOLDAP2_8:
		return "Open LDAP Public License json.8"
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
		return "SGI Free Software License B json.0"
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
		return "TORQUE json.5+ Software License v1.1"
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
		return "Zend License json.0"
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
		return "Academic Free License json.0."
	case SPDXLicenseAFL2_1:
		return "Academic Free License json.1."
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
		return "CeCILL Free Software License Agreement json.0."
	case SPDXLicenseCECILL2_1:
		return "CeCILL Free Software License Agreement json.1."
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
		return "Educational Community License json.0."
	case SPDXLicenseEFL1_0:
		return "Eiffel Forum License v1.0."
	case SPDXLicenseEFL2_0:
		return "Eiffel Forum License json.0."
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
		return "GNU General Public License json.0 only."
	case SPDXLicenseGPL2_0OrLater:
		return "GNU General Public License json.0 or later."
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
		return "GNU Library General Public License json.only."
	case SPDXLicenseLGPL2_0OrLater:
		return "GNU Library General Public License json.or later."
	case SPDXLicenseLGPL2_1Only:
		return "GNU Lesser General Public License json.1 only."
	case SPDXLicenseLGPL2_1OrLater:
		return "GNU Lesser General Public License json.1 or later."
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
		return "Open LDAP Public License json.0.1."
	case SPDXLicenseOLDAP2_0:
		return "Open LDAP Public License json.0 (or possibly 2.0A and 2.0B)."
	case SPDXLicenseOLDAP2_1:
		return "Open LDAP Public License json.1."
	case SPDXLicenseOLDAP2_2_1:
		return "Open LDAP Public License json.2.1."
	case SPDXLicenseOLDAP2_2_2:
		return "Open LDAP Public License 2.2.2."
	case SPDXLicenseOLDAP2_2:
		return "Open LDAP Public License json.2."
	case SPDXLicenseOLDAP2_3:
		return "Open LDAP Public License json.3."
	case SPDXLicenseOLDAP2_4:
		return "Open LDAP Public License json.4."
	case SPDXLicenseOLDAP2_5:
		return "Open LDAP Public License json.5."
	case SPDXLicenseOLDAP2_6:
		return "Open LDAP Public License json.6."
	case SPDXLicenseOLDAP2_7:
		return "Open LDAP Public License json.7."
	case SPDXLicenseOLDAP2_8:
		return "Open LDAP Public License json.8."
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
		return "SGI Free Software License B json.0."
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
		return "TORQUE json.5+ Software License v1.1."
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
		return "Zend License json.0."
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
