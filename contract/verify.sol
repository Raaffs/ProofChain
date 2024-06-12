    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.0;
    contract Verification{

        constructor(){
            owner=msg.sender;
        }

        enum DocStatus{
            accepted,
            rejected,
            pending
        }
        struct User{
            string  name;
            string  email;
        }

        struct Verifier{
            string  name;
            string  email;
            string  AadharNumber;
            string  institute;
            //Will need this in ./documents.sol to check if account verifying the documents
            //is actually a verifier. 
            bool    isApprovedVerifier;
        }
        struct Document{
            address     requester;
            address     verifiedBy; 
            string      name;
            string      description; 
            string      ipfshash;
            uint        docIndex;
            DocStatus   status;     
        }


        mapping(string=>Verifier) public verifiers;
        mapping(address=>User) private users;
        mapping(string=>Document) private documentList;
        mapping(address=>bool) userList;
        mapping(address=>uint[]) userDocIndex;
        //jack-ass-hack
        address[]   requesters;
        address[]   verifiedBy; 
        string[]    names;
        string[]    descriptions;
        string[]    ipfshash; 
        DocStatus[] status;

        uint docIndexCounter=0;
        
        function registerAsUser(string calldata _name, string calldata _email) public{
            users[msg.sender]=User({
                name:       _name,
                email:      _email
            });
            userList[msg.sender]=true;
        }
        function registerAsVerifier(string memory _name, string memory _email, string memory _aadharNum, string memory _institute,string memory _verifierAddr) public{
            verifiers[_verifierAddr]=Verifier({
                name:           _name,
                email:          _email,
                AadharNumber:   _aadharNum,
                institute:      _institute,
                isApprovedVerifier:     false
            });
        }

        function approveVerifier(string memory verifier)public{
            require(msg.sender==owner,"Only admin can perfom this action");
            verifiers[verifier].isApprovedVerifier=true;
        }

        address private owner;
        address public accountAddress; 
        


        function addDocument(string memory _name, string memory _ipfshash ,string memory _description) public{
            bool isUser=userList[msg.sender];
            
            require(isUser==true,"register first to verify");

            Document memory docs=Document({
                requester:          msg.sender,
                verifiedBy:         address(0),
                name:               _name,
                description:        _description,
                ipfshash:           _ipfshash,
                docIndex:           docIndexCounter,
                status:             DocStatus.pending
            });


            documentList[_ipfshash]=docs;
            //Only God and I know what I am doing

            //It'd have been great if solidity allowed to return mapping but it doesn't and now I am losing my brain over this
            //This shit below keeps track of each index of document uploaded by unique each user and stores them in an array 
            //We using this so that on application side we can properly retrive each document and their correspoing IPFS addresses. 

            //few months later:
            //I have no clue what userDocIndex does. 
            userDocIndex[msg.sender].push(docIndexCounter);
            requesters.push(msg.sender);
            verifiedBy.push(address(0));
            names.push(_name);
            descriptions.push(_description);
            ipfshash.push(_ipfshash);
            status.push(DocStatus.pending);

            docIndexCounter++;
        }

        function getDocumentList(string memory verifier) public view returns (address[] memory requester ,address[] memory verifer ,string[] memory name,string[] memory ipfs,string[] memory desc,DocStatus[] memory stats ,uint[] memory userDocId){
            if(verifiers[verifier].isApprovedVerifier){
                return (requesters,verifiedBy,names,ipfshash,descriptions,status,new uint[](0));
            }
            return(requesters,verifiedBy,names,new string[](0),descriptions,status,userDocIndex[msg.sender]);
        }

        function getDocumentsForVerifier(string memory verifier)public view returns (address[] memory requester ,address[] memory verifer ,string[] memory name,string[] memory ipfs,string[] memory desc,DocStatus[] memory stats){
            require(verifiers[verifier].isApprovedVerifier,"Only verifier can access this field");
            return (requesters,verifiedBy,names,ipfshash,descriptions,status);

        }



        function verifyDocuments(string memory ipfs, DocStatus _status,string memory verifier) public payable {
            require(verifiers[verifier].isApprovedVerifier==true,"You're not an approved verifier");
            documentList[ipfs].status=_status;
            documentList[ipfs].verifiedBy=msg.sender;
            uint index = documentList[ipfs].docIndex;

            status[index]=_status;
            verifiedBy[index]=msg.sender;
        }

        function checkVerifierStatus(string memory verifier)public view returns (bool){
            return verifiers[verifier].isApprovedVerifier;
        }
    
    }
