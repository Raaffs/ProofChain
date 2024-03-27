// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Verification{

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
        bool    isVerifier;
    }
    struct Document{
        address     requester;
        address     verifiedBy; 
        string      name;
        string      description; 
        string      docAddressOnIPFS;
        uint        docIndex;
        DocStatus   status;     
    }

 

    mapping(address=>Verifier) public verifiers;
    mapping(address=>User) private users;
    mapping(string=>Document) private documentList;
    mapping(address=>bool) userList;
    mapping(address=>uint[]) userDocIndex;

    address[]   requesters;
    address[]   verifiedBy; 
    string[]    names;
    string[]    descriptions;
    string[]    docAddressOnIPFS; 
    DocStatus[] status;

    uint docIndexCounter=0;
    
    function registerAsUser(string calldata _name, string calldata _email) public{
        users[msg.sender]=User({
            name:       _name,
            email:      _email
        });
        userList[msg.sender]=true;
    }
    function registerAsVerifier(string memory _name, string memory _email, string memory _aadharNum, string memory _institute) public{
        verifiers[msg.sender]=Verifier({
            name:           _name,
            email:          _email,
            AadharNumber:   _aadharNum,
            institute:      _institute,
            isVerifier:     true
        });
    }

    address private owner;
    address public accountAddress; 
    


    function addDocument(string calldata _name, string calldata _description, string calldata _docAddressOnIPFS) public  {
        bool isUser=userList[msg.sender];
        
        require(isUser==true,"register first to verify");

        Document memory docs=Document({
            requester:          msg.sender,
            verifiedBy:         address(0),
            name:               _name,
            description:        _description,
            docAddressOnIPFS:   _docAddressOnIPFS,
            docIndex:           docIndexCounter,
            status:             DocStatus.pending
        });


        documentList[_docAddressOnIPFS]=docs;
        //Only God and I know what I am doing

        //It'd have been great if solidity allowed to return mapping but it doesn't and now I am losing my brain over this
        //This shit below keeps track of each index of document uploaded by unique each user and stores them in an array 
        //We using this so that on application side we can properly retrive each document and their correspoing IPFS addresses. 
        userDocIndex[msg.sender].push(docIndexCounter);
        requesters.push(msg.sender);
        verifiedBy.push(address(0));
        names.push(_name);
        descriptions.push(_description);
        docAddressOnIPFS.push(_docAddressOnIPFS);
        status.push(DocStatus.pending);

        docIndexCounter++;
    }

    function getDocumentList() public view returns (address[] memory requester ,address[] memory verifer ,string[] memory name,string[] memory desc,string[] memory ipfsAddress,DocStatus[] memory stats ,uint[] memory userDocId){
        if(verifiers[msg.sender].isVerifier){
            return (requesters,verifiedBy,names,descriptions,docAddressOnIPFS,status,new uint[](0));
        }
        return(requesters,verifiedBy,names,descriptions,new string[](0),status,userDocIndex[msg.sender]);
    }

    function verifyDocuments(string memory _docAddressOnIPFS, DocStatus _status) public payable {
        bool isVerifier=verifiers[msg.sender].isVerifier;
        require(isVerifier==true,"You're not verifier");
        documentList[_docAddressOnIPFS].status=_status;
        documentList[_docAddressOnIPFS].verifiedBy=msg.sender;
        uint index = documentList[_docAddressOnIPFS].docIndex;

        status[index]=_status;
        verifiedBy[index]=msg.sender;
    }
 
}
