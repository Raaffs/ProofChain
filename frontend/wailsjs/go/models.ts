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

