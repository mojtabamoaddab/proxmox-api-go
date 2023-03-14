package proxmox

import "strconv"

type QemuScsiDisk struct {
	AsyncIO    QemuDiskAsyncIO   `json:"asyncio,omitempty"`
	Backup     bool              `json:"backup,omitempty"`
	Bandwidth  QemuDiskBandwidth `json:"bandwidth,omitempty"`
	Cache      QemuDiskCache     `json:"cache,omitempty"`
	Discard    bool              `json:"discard,omitempty"`
	EmulateSSD bool              `json:"emulatessd,omitempty"`
	Format     *QemuDiskFormat   `json:"format,omitempty"`
	Id         *uint             `json:"id,omitempty"`
	IOThread   bool              `json:"iothread,omitempty"`
	ReadOnly   bool              `json:"readonly,omitempty"`
	Replicate  bool              `json:"replicate,omitempty"`
	Serial     QemuDiskSerial    `json:"serial,omitempty"`
	Size       uint              `json:"size,omitempty"`
	Storage    string            `json:"storage,omitempty"`
}

func (disk *QemuScsiDisk) convertDataStructure() *qemuDisk {
	return &qemuDisk{
		AsyncIO:    disk.AsyncIO,
		Backup:     disk.Backup,
		Bandwidth:  disk.Bandwidth,
		Cache:      disk.Cache,
		Discard:    disk.Discard,
		EmulateSSD: disk.EmulateSSD,
		Format:     disk.Format,
		Id:         disk.Id,
		IOThread:   disk.IOThread,
		ReadOnly:   disk.ReadOnly,
		Replicate:  disk.Replicate,
		Serial:     disk.Serial,
		Size:       disk.Size,
		Storage:    disk.Storage,
		Type:       scsi,
	}
}

type QemuScsiDisks struct {
	Disk_0  *QemuScsiStorage `json:"0,omitempty"`
	Disk_1  *QemuScsiStorage `json:"1,omitempty"`
	Disk_2  *QemuScsiStorage `json:"2,omitempty"`
	Disk_3  *QemuScsiStorage `json:"3,omitempty"`
	Disk_4  *QemuScsiStorage `json:"4,omitempty"`
	Disk_5  *QemuScsiStorage `json:"5,omitempty"`
	Disk_6  *QemuScsiStorage `json:"6,omitempty"`
	Disk_7  *QemuScsiStorage `json:"7,omitempty"`
	Disk_8  *QemuScsiStorage `json:"8,omitempty"`
	Disk_9  *QemuScsiStorage `json:"9,omitempty"`
	Disk_10 *QemuScsiStorage `json:"10,omitempty"`
	Disk_11 *QemuScsiStorage `json:"11,omitempty"`
	Disk_12 *QemuScsiStorage `json:"12,omitempty"`
	Disk_13 *QemuScsiStorage `json:"13,omitempty"`
	Disk_14 *QemuScsiStorage `json:"14,omitempty"`
	Disk_15 *QemuScsiStorage `json:"15,omitempty"`
	Disk_16 *QemuScsiStorage `json:"16,omitempty"`
	Disk_17 *QemuScsiStorage `json:"17,omitempty"`
	Disk_18 *QemuScsiStorage `json:"18,omitempty"`
	Disk_19 *QemuScsiStorage `json:"19,omitempty"`
	Disk_20 *QemuScsiStorage `json:"20,omitempty"`
	Disk_21 *QemuScsiStorage `json:"21,omitempty"`
	Disk_22 *QemuScsiStorage `json:"22,omitempty"`
	Disk_23 *QemuScsiStorage `json:"23,omitempty"`
	Disk_24 *QemuScsiStorage `json:"24,omitempty"`
	Disk_25 *QemuScsiStorage `json:"25,omitempty"`
	Disk_26 *QemuScsiStorage `json:"26,omitempty"`
	Disk_27 *QemuScsiStorage `json:"27,omitempty"`
	Disk_28 *QemuScsiStorage `json:"28,omitempty"`
	Disk_29 *QemuScsiStorage `json:"29,omitempty"`
	Disk_30 *QemuScsiStorage `json:"30,omitempty"`
}

// TODO write test
func (disks QemuScsiDisks) mapToApiValues(currentDisks *QemuScsiDisks, vmID uint, params map[string]interface{}, changes *qemuUpdateChanges) {
	tmpCurrentDisks := QemuScsiDisks{}
	if currentDisks != nil {
		tmpCurrentDisks = *currentDisks
	}
	diskMap := disks.mapToIntMap()
	currentDiskMap := tmpCurrentDisks.mapToIntMap()
	for i := range diskMap {
		diskMap[i].convertDataStructure().markDiskChanges(currentDiskMap[i].convertDataStructure(), vmID, "scsi"+strconv.Itoa(int(i)), params, changes)
	}
}

func (disks QemuScsiDisks) mapToIntMap() map[uint8]*QemuScsiStorage {
	return map[uint8]*QemuScsiStorage{
		0:  disks.Disk_0,
		1:  disks.Disk_1,
		2:  disks.Disk_2,
		3:  disks.Disk_3,
		4:  disks.Disk_4,
		5:  disks.Disk_5,
		6:  disks.Disk_6,
		7:  disks.Disk_7,
		8:  disks.Disk_8,
		9:  disks.Disk_9,
		10: disks.Disk_10,
		11: disks.Disk_11,
		12: disks.Disk_12,
		13: disks.Disk_13,
		14: disks.Disk_14,
		15: disks.Disk_15,
		16: disks.Disk_16,
		17: disks.Disk_17,
		18: disks.Disk_18,
		19: disks.Disk_19,
		20: disks.Disk_20,
		21: disks.Disk_21,
		22: disks.Disk_22,
		23: disks.Disk_23,
		24: disks.Disk_24,
		25: disks.Disk_25,
		26: disks.Disk_26,
		27: disks.Disk_27,
		28: disks.Disk_28,
		29: disks.Disk_29,
		30: disks.Disk_30,
	}
}

// TODO write test
func (QemuScsiDisks) mapToStruct(params map[string]interface{}) *QemuScsiDisks {
	disks := QemuScsiDisks{}
	var structPopulated bool
	if _, isSet := params["scsi0"]; isSet {
		disks.Disk_0 = QemuScsiStorage{}.mapToStruct(params["scsi0"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi1"]; isSet {
		disks.Disk_1 = QemuScsiStorage{}.mapToStruct(params["scsi1"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi2"]; isSet {
		disks.Disk_2 = QemuScsiStorage{}.mapToStruct(params["scsi2"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi3"]; isSet {
		disks.Disk_3 = QemuScsiStorage{}.mapToStruct(params["scsi3"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi4"]; isSet {
		disks.Disk_4 = QemuScsiStorage{}.mapToStruct(params["scsi4"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi5"]; isSet {
		disks.Disk_5 = QemuScsiStorage{}.mapToStruct(params["scsi5"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi6"]; isSet {
		disks.Disk_6 = QemuScsiStorage{}.mapToStruct(params["scsi6"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi7"]; isSet {
		disks.Disk_7 = QemuScsiStorage{}.mapToStruct(params["scsi7"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi8"]; isSet {
		disks.Disk_8 = QemuScsiStorage{}.mapToStruct(params["scsi8"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi9"]; isSet {
		disks.Disk_9 = QemuScsiStorage{}.mapToStruct(params["scsi9"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi10"]; isSet {
		disks.Disk_10 = QemuScsiStorage{}.mapToStruct(params["scsi10"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi11"]; isSet {
		disks.Disk_11 = QemuScsiStorage{}.mapToStruct(params["scsi11"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi12"]; isSet {
		disks.Disk_12 = QemuScsiStorage{}.mapToStruct(params["scsi12"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi13"]; isSet {
		disks.Disk_13 = QemuScsiStorage{}.mapToStruct(params["scsi13"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi14"]; isSet {
		disks.Disk_14 = QemuScsiStorage{}.mapToStruct(params["scsi14"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi15"]; isSet {
		disks.Disk_15 = QemuScsiStorage{}.mapToStruct(params["scsi15"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi16"]; isSet {
		disks.Disk_16 = QemuScsiStorage{}.mapToStruct(params["scsi16"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi17"]; isSet {
		disks.Disk_17 = QemuScsiStorage{}.mapToStruct(params["scsi17"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi18"]; isSet {
		disks.Disk_18 = QemuScsiStorage{}.mapToStruct(params["scsi18"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi19"]; isSet {
		disks.Disk_19 = QemuScsiStorage{}.mapToStruct(params["scsi19"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi20"]; isSet {
		disks.Disk_20 = QemuScsiStorage{}.mapToStruct(params["scsi20"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi21"]; isSet {
		disks.Disk_21 = QemuScsiStorage{}.mapToStruct(params["scsi21"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi22"]; isSet {
		disks.Disk_22 = QemuScsiStorage{}.mapToStruct(params["scsi22"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi23"]; isSet {
		disks.Disk_23 = QemuScsiStorage{}.mapToStruct(params["scsi23"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi24"]; isSet {
		disks.Disk_24 = QemuScsiStorage{}.mapToStruct(params["scsi24"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi25"]; isSet {
		disks.Disk_25 = QemuScsiStorage{}.mapToStruct(params["scsi25"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi26"]; isSet {
		disks.Disk_26 = QemuScsiStorage{}.mapToStruct(params["scsi26"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi27"]; isSet {
		disks.Disk_27 = QemuScsiStorage{}.mapToStruct(params["scsi27"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi28"]; isSet {
		disks.Disk_28 = QemuScsiStorage{}.mapToStruct(params["scsi28"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi29"]; isSet {
		disks.Disk_29 = QemuScsiStorage{}.mapToStruct(params["scsi29"].(string))
		structPopulated = true
	}
	if _, isSet := params["scsi30"]; isSet {
		disks.Disk_30 = QemuScsiStorage{}.mapToStruct(params["scsi30"].(string))
		structPopulated = true
	}
	if structPopulated {
		return &disks
	}
	return nil
}

type QemuScsiPassthrough struct {
	AsyncIO    QemuDiskAsyncIO
	Backup     bool
	Bandwidth  QemuDiskBandwidth
	Cache      QemuDiskCache
	Discard    bool
	EmulateSSD bool
	File       string
	IOThread   bool
	ReadOnly   bool
	Replicate  bool
	Serial     QemuDiskSerial `json:"serial,omitempty"`
	Size       uint           //size is only returned and setting it has no effect
}

func (passthrough *QemuScsiPassthrough) convertDataStructure() *qemuDisk {
	return &qemuDisk{
		AsyncIO:    passthrough.AsyncIO,
		Backup:     passthrough.Backup,
		Bandwidth:  passthrough.Bandwidth,
		Cache:      passthrough.Cache,
		Discard:    passthrough.Discard,
		EmulateSSD: passthrough.EmulateSSD,
		File:       passthrough.File,
		IOThread:   passthrough.IOThread,
		ReadOnly:   passthrough.ReadOnly,
		Replicate:  passthrough.Replicate,
		Serial:     passthrough.Serial,
		Type:       scsi,
	}
}

type QemuScsiStorage struct {
	CdRom       *QemuCdRom
	CloudInit   *QemuCloudInitDisk
	Disk        *QemuScsiDisk
	Passthrough *QemuScsiPassthrough
}

// TODO write test
// converts to qemuStorage
func (storage *QemuScsiStorage) convertDataStructure() *qemuStorage {
	if storage == nil {
		return nil
	}
	generalizedStorage := qemuStorage{
		CdRom:     storage.CdRom,
		CloudInit: storage.CloudInit,
	}
	if storage.Disk != nil {
		generalizedStorage.Disk = storage.Disk.convertDataStructure()
	}
	if storage.Passthrough != nil {
		generalizedStorage.Passthrough = storage.Passthrough.convertDataStructure()
	}
	return &generalizedStorage
}

// TODO write test
func (QemuScsiStorage) mapToStruct(param string) *QemuScsiStorage {
	settings := splitStringOfSettings(param)
	tmpCdRom := qemuCdRom{}.mapToStruct(settings)
	if tmpCdRom != nil {
		if tmpCdRom.CdRom {
			return &QemuScsiStorage{CdRom: QemuCdRom{}.mapToStruct(*tmpCdRom)}
		} else {
			return &QemuScsiStorage{CloudInit: QemuCloudInitDisk{}.mapToStruct(*tmpCdRom)}
		}
	}

	tmpDisk := qemuDisk{}.mapToStruct(settings)
	if tmpDisk == nil {
		return nil
	}
	if tmpDisk.File == "" {
		return &QemuScsiStorage{Disk: &QemuScsiDisk{
			AsyncIO:    tmpDisk.AsyncIO,
			Backup:     tmpDisk.Backup,
			Bandwidth:  tmpDisk.Bandwidth,
			Cache:      tmpDisk.Cache,
			Discard:    tmpDisk.Discard,
			EmulateSSD: tmpDisk.EmulateSSD,
			Format:     tmpDisk.Format,
			Id:         tmpDisk.Id,
			IOThread:   tmpDisk.IOThread,
			ReadOnly:   tmpDisk.ReadOnly,
			Replicate:  tmpDisk.Replicate,
			Serial:     tmpDisk.Serial,
			Size:       tmpDisk.Size,
			Storage:    tmpDisk.Storage,
		}}
	}
	return &QemuScsiStorage{Passthrough: &QemuScsiPassthrough{
		AsyncIO:    tmpDisk.AsyncIO,
		Backup:     tmpDisk.Backup,
		Bandwidth:  tmpDisk.Bandwidth,
		Cache:      tmpDisk.Cache,
		Discard:    tmpDisk.Discard,
		EmulateSSD: tmpDisk.EmulateSSD,
		File:       tmpDisk.File,
		IOThread:   tmpDisk.IOThread,
		ReadOnly:   tmpDisk.ReadOnly,
		Replicate:  tmpDisk.Replicate,
		Serial:     tmpDisk.Serial,
		Size:       tmpDisk.Size,
	}}
}
