package fileutil

import (
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ghetzel/go-stockutil/stringutil"
	"github.com/h2non/filetype"
)

var once sync.Once // guards InitMime

var mimeFiles = []string{
	`/etc/mime.types`,
}

func AddMimeTypeFile(filename string) {
	if IsNonemptyFile(filename) {
		mimeFiles = append(mimeFiles, filename)
		InitMime()
	}
}

func init() {
	InitMime()
}

func InitMime() {
	// generated from Ubuntu 16.04.1 LTS /etc/mime.types
	mime.AddExtensionType(`.323`, `text/h323`)
	mime.AddExtensionType(`.3gp`, `video/3gpp`)
	mime.AddExtensionType(`.abw`, `application/x-abiword`)
	mime.AddExtensionType(`.ai`, `application/postscript`)
	mime.AddExtensionType(`.aif`, `audio/x-aiff`)
	mime.AddExtensionType(`.aifc`, `audio/x-aiff`)
	mime.AddExtensionType(`.aiff`, `audio/x-aiff`)
	mime.AddExtensionType(`.alc`, `chemical/x-alchemy`)
	mime.AddExtensionType(`.amr`, `audio/amr`)
	mime.AddExtensionType(`.anx`, `application/annodex`)
	mime.AddExtensionType(`.apk`, `application/vnd.android.package-archive`)
	mime.AddExtensionType(`.appcache`, `text/cache-manifest`)
	mime.AddExtensionType(`.~`, `application/x-trash`)
	mime.AddExtensionType(`.%`, `application/x-trash`)
	mime.AddExtensionType(`.art`, `image/x-jg`)
	mime.AddExtensionType(`.asc`, `text/plain`)
	mime.AddExtensionType(`.asf`, `video/x-ms-asf`)
	mime.AddExtensionType(`.asn`, `chemical/x-ncbi-asn1`)
	mime.AddExtensionType(`.asx`, `video/x-ms-asf`)
	mime.AddExtensionType(`.atom`, `application/atom+xml`)
	mime.AddExtensionType(`.atomcat`, `application/atomcat+xml`)
	mime.AddExtensionType(`.au`, `audio/basic`)
	mime.AddExtensionType(`.avi`, `video/x-msvideo`)
	mime.AddExtensionType(`.awb`, `audio/amr-wb`)
	mime.AddExtensionType(`.axa`, `audio/annodex`)
	mime.AddExtensionType(`.axv`, `video/annodex`)
	mime.AddExtensionType(`.bak`, `application/x-trash`)
	mime.AddExtensionType(`.b`, `chemical/x-molconn-Z`)
	mime.AddExtensionType(`.bcpio`, `application/x-bcpio`)
	mime.AddExtensionType(`.bib`, `text/x-bibtex`)
	mime.AddExtensionType(`.bmp`, `image/x-ms-bmp`)
	mime.AddExtensionType(`.book`, `application/x-maker`)
	mime.AddExtensionType(`.boo`, `text/x-boo`)
	mime.AddExtensionType(`.brf`, `text/plain`)
	mime.AddExtensionType(`.bsd`, `chemical/x-crossfire`)
	mime.AddExtensionType(`.c3d`, `chemical/x-chem3d`)
	mime.AddExtensionType(`.cab`, `application/x-cab`)
	mime.AddExtensionType(`.cac`, `chemical/x-cache`)
	mime.AddExtensionType(`.cache`, `chemical/x-cache`)
	mime.AddExtensionType(`.cap`, `application/vnd.tcpdump.pcap`)
	mime.AddExtensionType(`.cat`, `application/vnd.ms-pki.seccat`)
	mime.AddExtensionType(`.cbr`, `application/x-cbr`)
	mime.AddExtensionType(`.cbz`, `application/x-cbz`)
	mime.AddExtensionType(`.cc`, `text/x-c++src`)
	mime.AddExtensionType(`.cda`, `application/x-cdf`)
	mime.AddExtensionType(`.cdf`, `application/x-cdf`)
	mime.AddExtensionType(`.cdr`, `image/x-coreldraw`)
	mime.AddExtensionType(`.cdx`, `chemical/x-cdx`)
	mime.AddExtensionType(`.cdy`, `application/vnd.cinderella`)
	mime.AddExtensionType(`.cef`, `chemical/x-cxf`)
	mime.AddExtensionType(`.cer`, `chemical/x-cerius`)
	mime.AddExtensionType(`.chm`, `chemical/x-chemdraw`)
	mime.AddExtensionType(`.chrt`, `application/x-kchart`)
	mime.AddExtensionType(`.cif`, `chemical/x-cif`)
	mime.AddExtensionType(`.class`, `application/java-vm`)
	mime.AddExtensionType(`.cls`, `text/x-tex`)
	mime.AddExtensionType(`.cmdf`, `chemical/x-cmdf`)
	mime.AddExtensionType(`.cml`, `chemical/x-cml`)
	mime.AddExtensionType(`.cod`, `application/vnd.rim.cod`)
	mime.AddExtensionType(`.cpa`, `chemical/x-compass`)
	mime.AddExtensionType(`.cpio`, `application/x-cpio`)
	mime.AddExtensionType(`.cpp`, `text/x-c++src`)
	mime.AddExtensionType(`.cpt`, `image/x-corelphotopaint`)
	mime.AddExtensionType(`.cr2`, `image/x-canon-cr2`)
	mime.AddExtensionType(`.crl`, `application/x-pkcs7-crl`)
	mime.AddExtensionType(`.crw`, `image/x-canon-crw`)
	mime.AddExtensionType(`.csd`, `audio/csound`)
	mime.AddExtensionType(`.csf`, `chemical/x-cache-csf`)
	mime.AddExtensionType(`.csh`, `application/x-csh`)
	mime.AddExtensionType(`.csh`, `text/x-csh`)
	mime.AddExtensionType(`.csm`, `chemical/x-csml`)
	mime.AddExtensionType(`.csml`, `chemical/x-csml`)
	mime.AddExtensionType(`.css`, `text/css`)
	mime.AddExtensionType(`.csv`, `text/csv`)
	mime.AddExtensionType(`.c`, `text/x-csrc`)
	mime.AddExtensionType(`.c++`, `text/x-c++src`)
	mime.AddExtensionType(`.ctx`, `chemical/x-ctx`)
	mime.AddExtensionType(`.cu`, `application/cu-seeme`)
	mime.AddExtensionType(`.cxf`, `chemical/x-cxf`)
	mime.AddExtensionType(`.cxx`, `text/x-c++src`)
	mime.AddExtensionType(`.dcm`, `application/dicom`)
	mime.AddExtensionType(`.dcr`, `application/x-director`)
	mime.AddExtensionType(`.ddeb`, `application/vnd.debian.binary-package`)
	mime.AddExtensionType(`.deb`, `application/vnd.debian.binary-package`)
	mime.AddExtensionType(`.diff`, `text/x-diff`)
	mime.AddExtensionType(`.dif`, `video/dv`)
	mime.AddExtensionType(`.dir`, `application/x-director`)
	mime.AddExtensionType(`.djv`, `image/vnd.djvu`)
	mime.AddExtensionType(`.djvu`, `image/vnd.djvu`)
	mime.AddExtensionType(`.dl`, `video/dl`)
	mime.AddExtensionType(`.dms`, `application/x-dms`)
	mime.AddExtensionType(`.doc`, `application/msword`)
	mime.AddExtensionType(`.docm`, `application/vnd.ms-word.document.macroEnabled.12`)
	mime.AddExtensionType(`.dot`, `application/msword`)
	mime.AddExtensionType(`.dotm`, `application/vnd.ms-word.template.macroEnabled.12`)
	mime.AddExtensionType(`.d`, `text/x-dsrc`)
	mime.AddExtensionType(`.dvi`, `application/x-dvi`)
	mime.AddExtensionType(`.dv`, `video/dv`)
	mime.AddExtensionType(`.dx`, `chemical/x-jcamp-dx`)
	mime.AddExtensionType(`.dxr`, `application/x-director`)
	mime.AddExtensionType(`.eml`, `message/rfc822`)
	mime.AddExtensionType(`.ent`, `chemical/x-pdb`)
	mime.AddExtensionType(`.eot`, `application/vnd.ms-fontobject`)
	mime.AddExtensionType(`.eps2`, `application/postscript`)
	mime.AddExtensionType(`.eps3`, `application/postscript`)
	mime.AddExtensionType(`.eps`, `application/postscript`)
	mime.AddExtensionType(`.epsf`, `application/postscript`)
	mime.AddExtensionType(`.epsi`, `application/postscript`)
	mime.AddExtensionType(`.erf`, `image/x-epson-erf`)
	mime.AddExtensionType(`.es`, `application/ecmascript`)
	mime.AddExtensionType(`.etx`, `text/x-setext`)
	mime.AddExtensionType(`.fb`, `application/x-maker`)
	mime.AddExtensionType(`.fbdoc`, `application/x-maker`)
	mime.AddExtensionType(`.fig`, `application/x-xfig`)
	mime.AddExtensionType(`.flac`, `audio/flac`)
	mime.AddExtensionType(`.fli`, `video/fli`)
	mime.AddExtensionType(`.flv`, `video/x-flv`)
	mime.AddExtensionType(`.fm`, `application/x-maker`)
	mime.AddExtensionType(`.frame`, `application/x-maker`)
	mime.AddExtensionType(`.frm`, `application/x-maker`)
	mime.AddExtensionType(`.gal`, `chemical/x-gaussian-log`)
	mime.AddExtensionType(`.gam`, `chemical/x-gamess-input`)
	mime.AddExtensionType(`.gamin`, `chemical/x-gamess-input`)
	mime.AddExtensionType(`.gcd`, `text/x-pcs-gcd`)
	mime.AddExtensionType(`.gen`, `chemical/x-genbank`)
	mime.AddExtensionType(`.gf`, `application/x-tex-gf`)
	mime.AddExtensionType(`.gif`, `image/gif`)
	mime.AddExtensionType(`.gl`, `video/gl`)
	mime.AddExtensionType(`.gnumeric`, `application/x-gnumeric`)
	mime.AddExtensionType(`.gpt`, `chemical/x-mopac-graph`)
	mime.AddExtensionType(`.gsf`, `application/x-font`)
	mime.AddExtensionType(`.gsm`, `audio/x-gsm`)
	mime.AddExtensionType(`.gtar`, `application/x-gtar`)
	mime.AddExtensionType(`.gz`, `application/gzip`)
	mime.AddExtensionType(`.hdf`, `application/x-hdf`)
	mime.AddExtensionType(`.hh`, `text/x-c++hdr`)
	mime.AddExtensionType(`.hin`, `chemical/x-hin`)
	mime.AddExtensionType(`.hpp`, `text/x-c++hdr`)
	mime.AddExtensionType(`.hs`, `text/x-haskell`)
	mime.AddExtensionType(`.hta`, `application/hta`)
	mime.AddExtensionType(`.htc`, `text/x-component`)
	mime.AddExtensionType(`.h`, `text/x-chdr`)
	mime.AddExtensionType(`.h++`, `text/x-c++hdr`)
	mime.AddExtensionType(`.html`, `text/html`)
	mime.AddExtensionType(`.htm`, `text/html`)
	mime.AddExtensionType(`.hwp`, `application/x-hwp`)
	mime.AddExtensionType(`.hxx`, `text/x-c++hdr`)
	mime.AddExtensionType(`.ica`, `application/x-ica`)
	mime.AddExtensionType(`.ice`, `x-conference/x-cooltalk`)
	mime.AddExtensionType(`.ics`, `text/calendar`)
	mime.AddExtensionType(`.icz`, `text/calendar`)
	mime.AddExtensionType(`.ief`, `image/ief`)
	mime.AddExtensionType(`.iges`, `model/iges`)
	mime.AddExtensionType(`.igs`, `model/iges`)
	mime.AddExtensionType(`.iii`, `application/x-iphone`)
	mime.AddExtensionType(`.info`, `application/x-info`)
	mime.AddExtensionType(`.inp`, `chemical/x-gamess-input`)
	mime.AddExtensionType(`.ist`, `chemical/x-isostar`)
	mime.AddExtensionType(`.istr`, `chemical/x-isostar`)
	mime.AddExtensionType(`.jam`, `application/x-jam`)
	mime.AddExtensionType(`.java`, `text/x-java`)
	mime.AddExtensionType(`.jdx`, `chemical/x-jcamp-dx`)
	mime.AddExtensionType(`.jmz`, `application/x-jmol`)
	mime.AddExtensionType(`.jng`, `image/x-jng`)
	mime.AddExtensionType(`.jp2`, `image/jp2`)
	mime.AddExtensionType(`.jpeg`, `image/jpeg`)
	mime.AddExtensionType(`.jpe`, `image/jpeg`)
	mime.AddExtensionType(`.jpf`, `image/jpx`)
	mime.AddExtensionType(`.jpg2`, `image/jp2`)
	mime.AddExtensionType(`.jpg`, `image/jpeg`)
	mime.AddExtensionType(`.jpm`, `image/jpm`)
	mime.AddExtensionType(`.jpx`, `image/jpx`)
	mime.AddExtensionType(`.js`, `application/javascript`)
	mime.AddExtensionType(`.json`, `application/json`)
	mime.AddExtensionType(`.kar`, `audio/midi`)
	mime.AddExtensionType(`.key`, `application/pgp-keys`)
	mime.AddExtensionType(`.kin`, `chemical/x-kinemage`)
	mime.AddExtensionType(`.kml`, `application/vnd.google-earth.kml+xml`)
	mime.AddExtensionType(`.kmz`, `application/vnd.google-earth.kmz`)
	mime.AddExtensionType(`.ksp`, `application/x-kspread`)
	mime.AddExtensionType(`.kwd`, `application/x-kword`)
	mime.AddExtensionType(`.kwt`, `application/x-kword`)
	mime.AddExtensionType(`.latex`, `application/x-latex`)
	mime.AddExtensionType(`.lha`, `application/x-lha`)
	mime.AddExtensionType(`.lhs`, `text/x-literate-haskell`)
	mime.AddExtensionType(`.lin`, `application/bbolin`)
	mime.AddExtensionType(`.lsf`, `video/x-la-asf`)
	mime.AddExtensionType(`.lsx`, `video/x-la-asf`)
	mime.AddExtensionType(`.ltx`, `text/x-tex`)
	mime.AddExtensionType(`.ly`, `text/x-lilypond`)
	mime.AddExtensionType(`.lyx`, `application/x-lyx`)
	mime.AddExtensionType(`.lzh`, `application/x-lzh`)
	mime.AddExtensionType(`.lzx`, `application/x-lzx`)
	mime.AddExtensionType(`.m3g`, `application/m3g`)
	mime.AddExtensionType(`.m3u8`, `application/x-mpegURL`)
	mime.AddExtensionType(`.m3u`, `audio/mpegurl`)
	mime.AddExtensionType(`.m3u`, `audio/x-mpegurl`)
	mime.AddExtensionType(`.m4a`, `audio/mpeg`)
	mime.AddExtensionType(`.maker`, `application/x-maker`)
	mime.AddExtensionType(`.man`, `application/x-troff-man`)
	mime.AddExtensionType(`.mbox`, `application/mbox`)
	mime.AddExtensionType(`.mcif`, `chemical/x-mmcif`)
	mime.AddExtensionType(`.mcm`, `chemical/x-macmolecule`)
	mime.AddExtensionType(`.mdb`, `application/msaccess`)
	mime.AddExtensionType(`.me`, `application/x-troff-me`)
	mime.AddExtensionType(`.mesh`, `model/mesh`)
	mime.AddExtensionType(`.mid`, `audio/midi`)
	mime.AddExtensionType(`.midi`, `audio/midi`)
	mime.AddExtensionType(`.mif`, `application/x-mif`)
	mime.AddExtensionType(`.mkv`, `video/x-matroska`)
	mime.AddExtensionType(`.mm`, `application/x-freemind`)
	mime.AddExtensionType(`.mmf`, `application/vnd.smaf`)
	mime.AddExtensionType(`.mml`, `text/mathml`)
	mime.AddExtensionType(`.mng`, `video/x-mng`)
	mime.AddExtensionType(`.moc`, `text/x-moc`)
	mime.AddExtensionType(`.mol2`, `chemical/x-mol2`)
	mime.AddExtensionType(`.mol`, `chemical/x-mdl-molfile`)
	mime.AddExtensionType(`.moo`, `chemical/x-mopac-out`)
	mime.AddExtensionType(`.mop`, `chemical/x-mopac-input`)
	mime.AddExtensionType(`.mopcrt`, `chemical/x-mopac-input`)
	mime.AddExtensionType(`.movie`, `video/x-sgi-movie`)
	mime.AddExtensionType(`.mov`, `video/quicktime`)
	mime.AddExtensionType(`.mp2`, `audio/mpeg`)
	mime.AddExtensionType(`.mp3`, `audio/mpeg`)
	mime.AddExtensionType(`.mp4`, `video/mp4`)
	mime.AddExtensionType(`.mpc`, `chemical/x-mopac-input`)
	mime.AddExtensionType(`.mpega`, `audio/mpeg`)
	mime.AddExtensionType(`.mpeg`, `video/mpeg`)
	mime.AddExtensionType(`.mpe`, `video/mpeg`)
	mime.AddExtensionType(`.mpga`, `audio/mpeg`)
	mime.AddExtensionType(`.mpg`, `video/mpeg`)
	mime.AddExtensionType(`.mph`, `application/x-comsol`)
	mime.AddExtensionType(`.mpv`, `video/x-matroska`)
	mime.AddExtensionType(`.ms`, `application/x-troff-ms`)
	mime.AddExtensionType(`.msh`, `model/mesh`)
	mime.AddExtensionType(`.msi`, `application/x-msi`)
	mime.AddExtensionType(`.mvb`, `chemical/x-mopac-vib`)
	mime.AddExtensionType(`.mxf`, `application/mxf`)
	mime.AddExtensionType(`.mxu`, `video/vnd.mpegurl`)
	mime.AddExtensionType(`.nb`, `application/mathematica`)
	mime.AddExtensionType(`.nbp`, `application/mathematica`)
	mime.AddExtensionType(`.nc`, `application/x-netcdf`)
	mime.AddExtensionType(`.nef`, `image/x-nikon-nef`)
	mime.AddExtensionType(`.nwc`, `application/x-nwc`)
	mime.AddExtensionType(`.o`, `application/x-object`)
	mime.AddExtensionType(`.oda`, `application/oda`)
	mime.AddExtensionType(`.odb`, `application/vnd.oasis.opendocument.database`)
	mime.AddExtensionType(`.odc`, `application/vnd.oasis.opendocument.chart`)
	mime.AddExtensionType(`.odf`, `application/vnd.oasis.opendocument.formula`)
	mime.AddExtensionType(`.odg`, `application/vnd.oasis.opendocument.graphics`)
	mime.AddExtensionType(`.odi`, `application/vnd.oasis.opendocument.image`)
	mime.AddExtensionType(`.odm`, `application/vnd.oasis.opendocument.text-master`)
	mime.AddExtensionType(`.odp`, `application/vnd.oasis.opendocument.presentation`)
	mime.AddExtensionType(`.ods`, `application/vnd.oasis.opendocument.spreadsheet`)
	mime.AddExtensionType(`.odt`, `application/vnd.oasis.opendocument.text`)
	mime.AddExtensionType(`.oga`, `audio/ogg`)
	mime.AddExtensionType(`.ogg`, `audio/ogg`)
	mime.AddExtensionType(`.ogv`, `video/ogg`)
	mime.AddExtensionType(`.ogx`, `application/ogg`)
	mime.AddExtensionType(`.old`, `application/x-trash`)
	mime.AddExtensionType(`.one`, `application/onenote`)
	mime.AddExtensionType(`.onepkg`, `application/onenote`)
	mime.AddExtensionType(`.onetmp`, `application/onenote`)
	mime.AddExtensionType(`.onetoc2`, `application/onenote`)
	mime.AddExtensionType(`.opus`, `audio/ogg`)
	mime.AddExtensionType(`.orc`, `audio/csound`)
	mime.AddExtensionType(`.orf`, `image/x-olympus-orf`)
	mime.AddExtensionType(`.otf`, `application/font-sfnt`)
	mime.AddExtensionType(`.otg`, `application/vnd.oasis.opendocument.graphics-template`)
	mime.AddExtensionType(`.oth`, `application/vnd.oasis.opendocument.text-web`)
	mime.AddExtensionType(`.ots`, `application/vnd.oasis.opendocument.spreadsheet-template`)
	mime.AddExtensionType(`.ott`, `application/vnd.oasis.opendocument.text-template`)
	mime.AddExtensionType(`.pas`, `text/x-pascal`)
	mime.AddExtensionType(`.patch`, `text/x-diff`)
	mime.AddExtensionType(`.pbm`, `image/x-portable-bitmap`)
	mime.AddExtensionType(`.pcap`, `application/vnd.tcpdump.pcap`)
	mime.AddExtensionType(`.pcf`, `application/x-font-pcf`)
	mime.AddExtensionType(`.pcf.Z`, `application/x-font-pcf`)
	mime.AddExtensionType(`.pcx`, `image/pcx`)
	mime.AddExtensionType(`.pdb`, `chemical/x-pdb`)
	mime.AddExtensionType(`.pdf`, `application/pdf`)
	mime.AddExtensionType(`.pfa`, `application/x-font`)
	mime.AddExtensionType(`.pfb`, `application/x-font`)
	mime.AddExtensionType(`.pfr`, `application/font-tdpfr`)
	mime.AddExtensionType(`.pgn`, `application/x-chess-pgn`)
	mime.AddExtensionType(`.pk`, `application/x-tex-pk`)
	mime.AddExtensionType(`.pls`, `audio/x-scpls`)
	mime.AddExtensionType(`.pl`, `text/x-perl`)
	mime.AddExtensionType(`.pm`, `text/x-perl`)
	mime.AddExtensionType(`.png`, `image/png`)
	mime.AddExtensionType(`.pnm`, `image/x-portable-anymap`)
	mime.AddExtensionType(`.potm`, `application/vnd.ms-powerpoint.template.macroEnabled.12`)
	mime.AddExtensionType(`.pot`, `text/plain`)
	mime.AddExtensionType(`.ppam`, `application/vnd.ms-powerpoint.addin.macroEnabled.12`)
	mime.AddExtensionType(`.ppm`, `image/x-portable-pixmap`)
	mime.AddExtensionType(`.pps`, `application/vnd.ms-powerpoint`)
	mime.AddExtensionType(`.ppsm`, `application/vnd.ms-powerpoint.slideshow.macroEnabled.12`)
	mime.AddExtensionType(`.ppt`, `application/vnd.ms-powerpoint`)
	mime.AddExtensionType(`.prf`, `application/pics-rules`)
	mime.AddExtensionType(`.ps`, `application/postscript`)
	mime.AddExtensionType(`.psd`, `image/x-photoshop`)
	mime.AddExtensionType(`.p`, `text/x-pascal`)
	mime.AddExtensionType(`.py`, `text/x-python`)
	mime.AddExtensionType(`.qgs`, `application/x-qgis`)
	mime.AddExtensionType(`.qml`, `text/x-qml`)
	mime.AddExtensionType(`.qt`, `video/quicktime`)
	mime.AddExtensionType(`.ra`, `audio/x-pn-realaudio`)
	mime.AddExtensionType(`.ra`, `audio/x-realaudio`)
	mime.AddExtensionType(`.ram`, `audio/x-pn-realaudio`)
	mime.AddExtensionType(`.rar`, `application/rar`)
	mime.AddExtensionType(`.ras`, `image/x-cmu-raster`)
	mime.AddExtensionType(`.rb`, `application/x-ruby`)
	mime.AddExtensionType(`.rd`, `chemical/x-mdl-rdfile`)
	mime.AddExtensionType(`.rdf`, `application/rdf+xml`)
	mime.AddExtensionType(`.rdp`, `application/x-rdp`)
	mime.AddExtensionType(`.rgb`, `image/x-rgb`)
	mime.AddExtensionType(`.rm`, `audio/x-pn-realaudio`)
	mime.AddExtensionType(`.roff`, `application/x-troff`)
	mime.AddExtensionType(`.ros`, `chemical/x-rosdal`)
	mime.AddExtensionType(`.rss`, `application/x-rss+xml`)
	mime.AddExtensionType(`.rtf`, `application/rtf`)
	mime.AddExtensionType(`.rtx`, `text/richtext`)
	mime.AddExtensionType(`.rxn`, `chemical/x-mdl-rxnfile`)
	mime.AddExtensionType(`.scala`, `text/x-scala`)
	mime.AddExtensionType(`.sce`, `application/x-scilab`)
	mime.AddExtensionType(`.sci`, `application/x-scilab`)
	mime.AddExtensionType(`.sco`, `audio/csound`)
	mime.AddExtensionType(`.sct`, `text/scriptlet`)
	mime.AddExtensionType(`.sd2`, `audio/x-sd2`)
	mime.AddExtensionType(`.sda`, `application/vnd.stardivision.draw`)
	mime.AddExtensionType(`.sdc`, `application/vnd.stardivision.calc`)
	mime.AddExtensionType(`.sd`, `chemical/x-mdl-sdfile`)
	mime.AddExtensionType(`.sdd`, `application/vnd.stardivision.impress`)
	mime.AddExtensionType(`.sdf`, `application/vnd.stardivision.math`)
	mime.AddExtensionType(`.sdf`, `chemical/x-mdl-sdfile`)
	mime.AddExtensionType(`.sds`, `application/vnd.stardivision.chart`)
	mime.AddExtensionType(`.sdw`, `application/vnd.stardivision.writer`)
	mime.AddExtensionType(`.sfd`, `application/vnd.font-fontforge-sfd`)
	mime.AddExtensionType(`.sfv`, `text/x-sfv`)
	mime.AddExtensionType(`.sgf`, `application/x-go-sgf`)
	mime.AddExtensionType(`.sgl`, `application/vnd.stardivision.writer-global`)
	mime.AddExtensionType(`.sh`, `application/x-sh`)
	mime.AddExtensionType(`.shar`, `application/x-shar`)
	mime.AddExtensionType(`.shp`, `application/x-qgis`)
	mime.AddExtensionType(`.sh`, `text/x-sh`)
	mime.AddExtensionType(`.shtml`, `text/html`)
	mime.AddExtensionType(`.shx`, `application/x-qgis`)
	mime.AddExtensionType(`.sid`, `audio/prs.sid`)
	mime.AddExtensionType(`.sik`, `application/x-trash`)
	mime.AddExtensionType(`.silo`, `model/mesh`)
	mime.AddExtensionType(`.sis`, `application/vnd.symbian.install`)
	mime.AddExtensionType(`.sisx`, `x-epoc/x-sisx-app`)
	mime.AddExtensionType(`.sit`, `application/x-stuffit`)
	mime.AddExtensionType(`.sitx`, `application/x-stuffit`)
	mime.AddExtensionType(`.skd`, `application/x-koan`)
	mime.AddExtensionType(`.skm`, `application/x-koan`)
	mime.AddExtensionType(`.skp`, `application/x-koan`)
	mime.AddExtensionType(`.skt`, `application/x-koan`)
	mime.AddExtensionType(`.sldm`, `application/vnd.ms-powerpoint.slide.macroEnabled.12`)
	mime.AddExtensionType(`.smi`, `application/smil+xml`)
	mime.AddExtensionType(`.smil`, `application/smil+xml`)
	mime.AddExtensionType(`.snd`, `audio/basic`)
	mime.AddExtensionType(`.spc`, `chemical/x-galactic-spc`)
	mime.AddExtensionType(`.spx`, `audio/ogg`)
	mime.AddExtensionType(`.sql`, `application/x-sql`)
	mime.AddExtensionType(`.srt`, `text/plain`)
	mime.AddExtensionType(`.stc`, `application/vnd.sun.xml.calc.template`)
	mime.AddExtensionType(`.std`, `application/vnd.sun.xml.draw.template`)
	mime.AddExtensionType(`.sti`, `application/vnd.sun.xml.impress.template`)
	mime.AddExtensionType(`.stl`, `application/sla`)
	mime.AddExtensionType(`.stw`, `application/vnd.sun.xml.writer.template`)
	mime.AddExtensionType(`.sty`, `text/x-tex`)
	mime.AddExtensionType(`.sv4cpio`, `application/x-sv4cpio`)
	mime.AddExtensionType(`.sv4crc`, `application/x-sv4crc`)
	mime.AddExtensionType(`.svg`, `image/svg+xml`)
	mime.AddExtensionType(`.svgz`, `image/svg+xml`)
	mime.AddExtensionType(`.sw`, `chemical/x-swissprot`)
	mime.AddExtensionType(`.sxc`, `application/vnd.sun.xml.calc`)
	mime.AddExtensionType(`.sxd`, `application/vnd.sun.xml.draw`)
	mime.AddExtensionType(`.sxg`, `application/vnd.sun.xml.writer.global`)
	mime.AddExtensionType(`.sxi`, `application/vnd.sun.xml.impress`)
	mime.AddExtensionType(`.sxm`, `application/vnd.sun.xml.math`)
	mime.AddExtensionType(`.sxw`, `application/vnd.sun.xml.writer`)
	mime.AddExtensionType(`.t`, `application/x-troff`)
	mime.AddExtensionType(`.tar`, `application/x-tar`)
	mime.AddExtensionType(`.tcl`, `application/x-tcl`)
	mime.AddExtensionType(`.tcl`, `text/x-tcl`)
	mime.AddExtensionType(`.texi`, `application/x-texinfo`)
	mime.AddExtensionType(`.texinfo`, `application/x-texinfo`)
	mime.AddExtensionType(`.tex`, `text/x-tex`)
	mime.AddExtensionType(`.text`, `text/plain`)
	mime.AddExtensionType(`.tgf`, `chemical/x-mdl-tgf`)
	mime.AddExtensionType(`.thmx`, `application/vnd.ms-officetheme`)
	mime.AddExtensionType(`.tiff`, `image/tiff`)
	mime.AddExtensionType(`.tif`, `image/tiff`)
	mime.AddExtensionType(`.tk`, `text/x-tcl`)
	mime.AddExtensionType(`.tm`, `text/texmacs`)
	mime.AddExtensionType(`.tr`, `application/x-troff`)
	mime.AddExtensionType(`.tsp`, `application/dsptype`)
	mime.AddExtensionType(`.ts`, `video/MP2T`)
	mime.AddExtensionType(`.ttf`, `application/font-sfnt`)
	mime.AddExtensionType(`.ttl`, `text/turtle`)
	mime.AddExtensionType(`.txt`, `text/plain`)
	mime.AddExtensionType(`.udeb`, `application/vnd.debian.binary-package`)
	mime.AddExtensionType(`.uls`, `text/iuls`)
	mime.AddExtensionType(`.ustar`, `application/x-ustar`)
	mime.AddExtensionType(`.vcard`, `text/vcard`)
	mime.AddExtensionType(`.vcd`, `application/x-cdlink`)
	mime.AddExtensionType(`.vcf`, `text/vcard`)
	mime.AddExtensionType(`.vcs`, `text/x-vcalendar`)
	mime.AddExtensionType(`.vmd`, `chemical/x-vmd`)
	mime.AddExtensionType(`.vrml`, `model/vrml`)
	mime.AddExtensionType(`.vrml`, `x-world/x-vrml`)
	mime.AddExtensionType(`.vrm`, `x-world/x-vrml`)
	mime.AddExtensionType(`.vsd`, `application/vnd.visio`)
	mime.AddExtensionType(`.vss`, `application/vnd.visio`)
	mime.AddExtensionType(`.vst`, `application/vnd.visio`)
	mime.AddExtensionType(`.vsw`, `application/vnd.visio`)
	mime.AddExtensionType(`.wad`, `application/x-doom`)
	mime.AddExtensionType(`.wav`, `audio/x-wav`)
	mime.AddExtensionType(`.wax`, `audio/x-ms-wax`)
	mime.AddExtensionType(`.wbmp`, `image/vnd.wap.wbmp`)
	mime.AddExtensionType(`.wbxml`, `application/vnd.wap.wbxml`)
	mime.AddExtensionType(`.webm`, `video/webm`)
	mime.AddExtensionType(`.wk`, `application/x-123`)
	mime.AddExtensionType(`.wma`, `audio/x-ms-wma`)
	mime.AddExtensionType(`.wmd`, `application/x-ms-wmd`)
	mime.AddExtensionType(`.wmlc`, `application/vnd.wap.wmlc`)
	mime.AddExtensionType(`.wmlsc`, `application/vnd.wap.wmlscriptc`)
	mime.AddExtensionType(`.wmls`, `text/vnd.wap.wmlscript`)
	mime.AddExtensionType(`.wml`, `text/vnd.wap.wml`)
	mime.AddExtensionType(`.wm`, `video/x-ms-wm`)
	mime.AddExtensionType(`.wmv`, `video/x-ms-wmv`)
	mime.AddExtensionType(`.wmx`, `video/x-ms-wmx`)
	mime.AddExtensionType(`.wmz`, `application/x-ms-wmz`)
	mime.AddExtensionType(`.woff`, `application/font-woff`)
	mime.AddExtensionType(`.wp5`, `application/vnd.wordperfect5.1`)
	mime.AddExtensionType(`.wpd`, `application/vnd.wordperfect`)
	mime.AddExtensionType(`.wrl`, `model/vrml`)
	mime.AddExtensionType(`.wrl`, `x-world/x-vrml`)
	mime.AddExtensionType(`.wsc`, `text/scriptlet`)
	mime.AddExtensionType(`.wvx`, `video/x-ms-wvx`)
	mime.AddExtensionType(`.wz`, `application/x-wingz`)
	mime.AddExtensionType(`.x3db`, `model/x3d+binary`)
	mime.AddExtensionType(`.x3d`, `model/x3d+xml`)
	mime.AddExtensionType(`.x3dv`, `model/x3d+vrml`)
	mime.AddExtensionType(`.xbm`, `image/x-xbitmap`)
	mime.AddExtensionType(`.xcf`, `application/x-xcf`)
	mime.AddExtensionType(`.xht`, `application/xhtml+xml`)
	mime.AddExtensionType(`.xhtml`, `application/xhtml+xml`)
	mime.AddExtensionType(`.xlam`, `application/vnd.ms-excel.addin.macroEnabled.12`)
	mime.AddExtensionType(`.xlb`, `application/vnd.ms-excel`)
	mime.AddExtensionType(`.xls`, `application/vnd.ms-excel`)
	mime.AddExtensionType(`.xlsb`, `application/vnd.ms-excel.sheet.binary.macroEnabled.12`)
	mime.AddExtensionType(`.xlsm`, `application/vnd.ms-excel.sheet.macroEnabled.12`)
	mime.AddExtensionType(`.xlt`, `application/vnd.ms-excel`)
	mime.AddExtensionType(`.xltm`, `application/vnd.ms-excel.template.macroEnabled.12`)
	mime.AddExtensionType(`.xml`, `application/xml`)
	mime.AddExtensionType(`.xpi`, `application/x-xpinstall`)
	mime.AddExtensionType(`.xpm`, `image/x-xpixmap`)
	mime.AddExtensionType(`.xsd`, `application/xml`)
	mime.AddExtensionType(`.xsl`, `application/xslt+xml`)
	mime.AddExtensionType(`.xslt`, `application/xslt+xml`)
	mime.AddExtensionType(`.xspf`, `application/xspf+xml`)
	mime.AddExtensionType(`.xtel`, `chemical/x-xtel`)
	mime.AddExtensionType(`.xul`, `application/vnd.mozilla.xul+xml`)
	mime.AddExtensionType(`.xwd`, `image/x-xwindowdump`)
	mime.AddExtensionType(`.xyz`, `chemical/x-xyz`)
	mime.AddExtensionType(`.xz`, `application/x-xz`)
	mime.AddExtensionType(`.yaml`, `application/x-yaml`)
	mime.AddExtensionType(`.yml`, `application/x-yaml`)
	mime.AddExtensionType(`.zip`, `application/zip`)
	mime.AddExtensionType(`.zmt`, `chemical/x-mopac-input`)

	for _, filename := range mimeFiles {
		if lines, err := ReadAllLines(filename); err == nil {
			for _, line := range lines {
				line = strings.TrimSpace(line)

				if line == `` || line[0] == '#' {
					continue
				}

				mimetype, extset := stringutil.SplitPair(line, ` `)
				mimetype = strings.TrimSpace(mimetype)
				exts := strings.Split(strings.TrimSpace(extset), ` `)

				for _, ext := range exts {
					ext = strings.TrimSpace(ext)

					if len(ext) > 0 {
						if ext[0] != '.' {
							ext = `.` + ext
						}

						mime.AddExtensionType(ext, mimetype)
					}
				}
			}
		}
	}
}

// GetMimeType goes out of its way to really, really try to figure out the MIME type of a given
// filename or io.Reader.  If the first argument satisfies the io.Seeker interface, the seeker
// will Seek() back to the beginning.  If it satisfies the io.Closer interface, it will be closed.
func GetMimeType(filenameOrReader interface{}, fallback ...string) string {
	once.Do(InitMime)

	if filename, ok := filenameOrReader.(string); ok {
		if mt := mime.TypeByExtension(filepath.Ext(filename)); mt != `` {
			return mt
		} else if IsNonemptyFile(filename) {
			if file, err := os.Open(filename); err == nil {
				defer file.Close()
				filenameOrReader = file
			}
		}
	}

	// if we've gotten here, try to read the actual contents of the file to detect the type
	if reader, ok := filenameOrReader.(io.Reader); ok {
		if detected, err := filetype.MatchReader(reader); err == nil && detected.MIME.Value != `` {
			return detected.MIME.Value
		}

		if seeker, ok := filenameOrReader.(io.Seeker); ok {
			seeker.Seek(0, io.SeekStart)
		}
	}

	if len(fallback) > 0 && fallback[0] != `` {
		return fallback[0]
	} else {
		return ``
	}
}
