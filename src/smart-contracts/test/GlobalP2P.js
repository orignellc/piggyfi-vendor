const { expect } = require("chai");

before(async () => {
  const [_, owner, addressOne, addressTwo] = await ethers.getSigners();
  // const WalletLogicV1 = await ethers.getContractFactory("WalletLogicV1");
  // const walletLogic = await WalletLogicV1.deploy();
  // const GlobalP2P = await ethers.getContractFactory("GlobalP2P");
  // const globalp2p = await GlobalP2P.deploy();
  // const MockUSD = await ethers.getContractFactory("MockUSD");
  // const cUSD = await MockUSD.deploy();

  // await globalp2p.initialize(walletLogic.address, cUSD.address);

  this.owner = owner;
  // this.walletLogic = walletLogic;
  // this.globalp2p = globalp2p;
  // this.addressOne = addressOne;
  // this.addressTwo = addressTwo;
  // this.cUSD = cUSD;
});

describe("Global P2P Wallet Creation", () => {
  it.only("Send Funds", async () => {
    trx = await this.owner.sendTransaction({
      to: "0x2195d5D9A6144dC9A698FdC222aE9b7dC6226A9a",
      value: ethers.utils.parseUnits("20"),
    });

    console.log(trx);
    bal = await this.owner.getBalance();
    console.log(bal.toString());
    // console.log(ethers.utils.parseUnits(bal.toString()));
  });

  it("should confirm wallet logic is accurate", async () => {
    const logicAddress = await this.globalp2p.walletLogic();

    expect(this.walletLogic.address).to.equal(logicAddress);
  });

  it("should pause all transaction", async () => {
    const trx = await this.globalp2p.pause();

    console.log(trx);

    const paused = await this.globalp2p.paused();

    expect(paused).to.be.true;
  });

  it("should unpause all transaction", async () => {
    await this.globalp2p.unpause();

    const paused = await this.globalp2p.paused();

    expect(paused).to.be.false;
  });

  it("should deploy new wallet", async () => {
    this.uuid1 = "83608376-0f79-11ed-861d-0242ac120002";

    await expect(this.globalp2p.deployWallet(this.uuid1)).to.emit(
      this.globalp2p,
      "newWallet"
    );
  });

  it("should accept ether sent to it gp2p", async () => {
    const amount = ethers.utils.parseUnits("1");

    await this.addressOne.sendTransaction({
      to: this.globalp2p.address,
      value: amount,
    });

    const balance = await ethers.provider.getBalance(this.globalp2p.address);

    expect(balance).to.equal(amount);
  });

  it("should send ether balance back to owner from gp2p", async () => {
    const amount = ethers.utils.parseUnits("1");

    expect(this.globalp2p.withdrawBalance()).to.changeEtherBalances(
      [this.owner, this.globalp2p],
      [amount, -amount]
    );
  });

  it("should revert deploy same uuid", async () => {
    expect(this.globalp2p.deployWallet(this.uuid1)).to.be.revertedWith(
      "GP:User exist"
    );
  });

  it("should revert deploy new wallet", async () => {
    this.uuid2 = "83608632-0f79-11ed-861d-0242ac120002";
    expect(
      this.globalp2p.connect(this.addressOne).deployWallet(this.uuid2)
    ).to.be.revertedWith("Ownable: caller is not the owner");
  });

  it("should revert send ether balance to owner", async () => {
    const amount = ethers.utils.parseUnits("1");

    await expect(
      this.globalp2p.connect(this.addressOne).withdrawBalance(amount)
    ).to.be.revertedWith("Ownable: caller is not the owner");
  });
});

describe("WalletLogic Contract", () => {
  it("should send 500 cUSD to wallet with starting balance", async () => {
    this.agent1Address = await this.globalp2p.wallets(this.uuid1);
    const amount = ethers.utils.parseUnits("500");

    await this.cUSD.mint(this.agent1Address, amount);

    const balance = await this.cUSD.balanceOf(this.agent1Address);

    expect(balance).to.equal(amount);
  });

  it("should get version of wallet", async () => {
    const agent = await getAgent([
      "function version() public view returns (string memory)",
    ]);

    version = await agent.version();

    expect(version).to.be.eq("v1.1.0");
  });

  it("should confirm funding wallet fire events", async () => {
    const amount = ethers.utils.parseEther("500");

    expect(this.cUSD.mint(this.agent1Address, amount)).to.emit(
      this.agent,
      "newBalance"
    );
  });

  it("should confirm contracts total balance is 500 cUSD", async () => {
    const totalBalance = await this.cUSD.balanceOf(this.agent1Address);

    expect(totalBalance).to.equal(ethers.utils.parseUnits("1000"));
  });

  it("should execute transaction", async () => {
    const agent = await getAgent([
      "function execTransaction(uint256 amount, uint256 fee, address payable customer, string memory quoteId) public",
      "function pendingFees() public view returns (uint256)",
      "function spendableBalance() public view returns (uint256)",
    ]);

    const purchase = ethers.utils.parseUnits("100");
    const fee = ethers.utils.parseUnits("1.5");

    // Agent wallet contains 1000 cUSD at this point
    await agent
      .connect(ethers.provider.getSigner())
      .execTransaction(
        purchase,
        fee,
        this.addressTwo.address,
        "83608754-0f79-11ed-861d-0242ac120002"
      );

    const newBalance1 = await agent.spendableBalance();
    const newBalance2 = await this.cUSD.balanceOf(this.addressTwo.address);
    const pendingFees = await agent.pendingFees();

    expect(pendingFees).to.equal(fee);
    expect(newBalance2).to.equal(ethers.utils.parseUnits("100"));
    expect(newBalance1).to.equal(ethers.utils.parseUnits("898.5"));
  });

  it("should freeze 100 cUSD from contracts balance", async () => {
    const agent = await getAgent(["function freeze(uint256 amount) public"]);

    const freeze = ethers.utils.parseUnits("100");
    const outstandingBal = ethers.utils.parseUnits("798.5");

    // Agent wallet contains 1000 cUSD at this point
    expect(await agent.connect(ethers.provider.getSigner()).freeze(freeze))
      .to.emit(agent, "newFreeze")
      .withArgs(freeze, outstandingBal, this.agent1Address);
  });

  it("should confirm amount allowed to withdraw is equal withdrawalble balance", async () => {
    const agent = await getAgent([
      "function withdrawLiquidity(uint256 amount, uint256 withdrawalFee, address payable to) public",
      "function spendableBalance() public view returns (uint256)",
    ]);

    expect(
      await agent.connect(ethers.provider.getSigner()).withdrawLiquidity(
        ethers.utils.parseUnits("790.51500"), //790.51500
        ethers.utils.parseUnits("7.98500"), //7.98500 Assuming 1% is Fee (1% of 798.5 - withdrawable)
        this.addressTwo.address
      )
    )
      .to.changeTokenBalances(
        this.cUSD,
        [this.agent1Address, this.addressTwo.address],
        [-790.515, 790.515]
      )
      .emit("newWithdrawal");
  });

  it("should have 100 cUSD withdrawable when we unfreeze disputed amount", async () => {
    const agent = await getAgent(["function unfreeze(uint256 amount) public"]);

    const freeze = ethers.utils.parseUnits("100");

    expect(await agent.connect(ethers.provider.getSigner()).unfreeze(freeze))
      .to.emit(agent, "newUnFreeze")
      .withArgs(freeze, freeze, this.agent1Address);
  });
});

const getAgent = (abi) => {
  const agent = new ethers.Contract(this.agent1Address, abi, ethers.provider);

  return agent;
};

// 83608880-0f79-11ed-861d-0242ac120002

// 83608998-0f79-11ed-861d-0242ac120002

// 83608ac4-0f79-11ed-861d-0242ac120002
