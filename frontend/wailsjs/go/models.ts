export namespace blockchain {
	
	export class VerificationDocument {
	    ID: number;
	    Requester: string;
	    Verifier: string;
	    Institute: string;
	    Name: string;
	    Desc: string;
	    IpfsAddress: string;
	    ShaHash: string;
	    Stats: number;
	
	    static createFrom(source: any = {}) {
	        return new VerificationDocument(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Requester = source["Requester"];
	        this.Verifier = source["Verifier"];
	        this.Institute = source["Institute"];
	        this.Name = source["Name"];
	        this.Desc = source["Desc"];
	        this.IpfsAddress = source["IpfsAddress"];
	        this.ShaHash = source["ShaHash"];
	        this.Stats = source["Stats"];
	    }
	}

}

export namespace models {
	
	export class CertificateData {
	    certificateName: string;
	    publicAddress: string;
	    name: string;
	    address: string;
	    age: string;
	    birthDate: string;
	    uniqueId: string;
	
	    static createFrom(source: any = {}) {
	        return new CertificateData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.certificateName = source["certificateName"];
	        this.publicAddress = source["publicAddress"];
	        this.name = source["name"];
	        this.address = source["address"];
	        this.age = source["age"];
	        this.birthDate = source["birthDate"];
	        this.uniqueId = source["uniqueId"];
	    }
	}

}

